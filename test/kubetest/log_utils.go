package kubetest

import (
	"fmt"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// ShowLogs prints logs from containers in case of fail/panic or enabled logging in file
func ShowLogs(k8s *K8s, t *testing.T) {
	if r := recover(); r != nil {
		showLogs(k8s, t)
		panic(r)
	} else if t.Failed() || shouldShowLogs() {
		showLogs(k8s, t)
	}
}
func showLogs(k8s *K8s, t *testing.T) {
	pods := k8s.ListPods()
	for i := 0; i < len(pods); i++ {
		showPodLogs(k8s, t, &pods[i])
	}
}

func showPodLogs(k8s *K8s, t *testing.T, pod *v1.Pod) {
	for i := 0; i < len(pod.Spec.Containers); i++ {
		c := &pod.Spec.Containers[i]
		name := pod.Name + ":" + c.Name
		logs, err := k8s.GetLogs(pod, c.Name)
		writeLogFunc := logTransaction

		if shouldShowLogs() && t != nil {
			writeLogFunc = func(name string, content string) {
				logErr := logFile(name, filepath.Join(logsDir(), t.Name()), content)
				if logErr != nil {
					logrus.Errorf("Can't log in file, reason %v", logErr)
					logTransaction(name, content)
				} else {
					logrus.Infof("Saved log for %v. Check dir %v", name, logsDir())
				}
			}
		}

		if err == nil {
			writeLogFunc(name, logs)
		}
		logs, err = k8s.GetLogsWithOptions(pod, &v1.PodLogOptions{
			Container: c.Name,
			Previous:  true,
		})
		if err == nil {
			writeLogFunc(name+"-previous", logs)
		}

	}
}

func logsDir() string {
	logDir := DefaultLogDir
	if dir, ok := os.LookupEnv(WritePodLogsDir); ok {
		logDir = dir
	}
	return logDir
}

func shouldShowLogs() bool {
	if v, ok := os.LookupEnv(WritePodLogsInFile); ok {
		if v == "true" {
			return true
		}
	}
	return false
}

func logFile(name, dir, content string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	path := filepath.Join(dir, name)
	var _, err = os.Stat(path)
	if os.IsExist(err) {
		err = os.Remove(path)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(name + ".log")
	if err != nil {
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func logTransaction(name, content string) {
	f := logrus.StandardLogger().Formatter
	logrus.SetFormatter(&innerLogFormatter{})

	drawer := transactionDrawer{
		buff:       strings.Builder{},
		lineLength: MaxTransactionLineWidth,
		drawUnit:   TransactionLogUnit,
	}
	drawer.drawLine()
	drawer.drawLineWithName(StartLogsOf + name)
	drawer.drawLine()
	drawer.drawText(content)
	drawer.drawLine()
	drawer.drawLineWithName(EndLogsOf + name)
	drawer.drawLine()
	logrus.Println(drawer.buff.String())
	logrus.SetFormatter(f)
}

type transactionDrawer struct {
	buff       strings.Builder
	lineLength int
	drawUnit   rune
}

func (t *transactionDrawer) drawText(text string) {
	_, _ = t.buff.WriteString(text)
}

func (t *transactionDrawer) drawLine() {
	_, _ = t.buff.WriteString(strings.Repeat(string(t.drawUnit), t.lineLength))
	_, _ = t.buff.WriteRune('\n')
}
func (t *transactionDrawer) drawLineWithName(name string) {
	sideWidth := int(math.Max(float64(t.lineLength-len(name)), 0)) / 2
	for i := 0; i < sideWidth; i++ {
		_, _ = t.buff.WriteRune(t.drawUnit)
	}
	_, _ = t.buff.WriteString(name)
	for i := t.buff.Len(); i < MaxTransactionLineWidth; i++ {
		_, _ = t.buff.WriteRune(t.drawUnit)
	}
	_, _ = t.buff.WriteRune('\n')
}

type innerLogFormatter struct {
}

func (*innerLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("[%v] %v\n%v", entry.Level.String(), entry.Time, entry.Message)), nil
}
