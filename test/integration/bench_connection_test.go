// +build bench

package nsmd_integration_tests

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"

	"github.com/networkservicemesh/networkservicemesh/test/kubetest"
)

func TestOneTimeConnectionMemif(t *testing.T) {
	t.Skip()
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}
	testOneTimeConnection(t, 1, kubetest.DeployVppAgentNSC, kubetest.DeployVppAgentICMP, kubetest.IsVppAgentNsePinged)
}
func TestOneTimeConnection(t *testing.T) {
	t.Skip()
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}
	testOneTimeConnection(t, 2, kubetest.DeployNSC, kubetest.DeployICMP, kubetest.IsNsePinged)
}

func TestMovingConnection(t *testing.T) {
	t.Skip()
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}
	testMovingConnection(t, 2, kubetest.DeployNSC, kubetest.DeployICMP, kubetest.IsNsePinged)
}

func TestMovingConnectionMemif(t *testing.T) {
	t.Skip()
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}
	testMovingConnection(t, 1, kubetest.DeployVppAgentNSC, kubetest.DeployVppAgentICMP, kubetest.IsVppAgentNsePinged)
}

func TestOneToOneConnection(t *testing.T) {
	t.Skip()
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}
	testOneToOneConnection(t, 2, kubetest.DeployNSC, kubetest.DeployICMP, kubetest.IsNsePinged)
}

func TestOneToOneConnectionMemif(t *testing.T) {
	t.Skip()
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}
	testOneToOneConnection(t, 1, kubetest.DeployVppAgentNSC, kubetest.DeployVppAgentICMP, kubetest.IsVppAgentNsePinged)
}

func TestOneTimeNseAndNscConnection(t *testing.T) {
	RegisterTestingT(t)
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}
	assert := NewWithT(t)
	testOneTimeNseAndNscConnection(t, 2, kubetest.DefaultTestingPodFixture(assert), assert)
}

func TestOneTimeNseAndNscConnectionMemif(t *testing.T) {
	RegisterTestingT(t)
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}
	assert := NewWithT(t)
	testOneTimeNseAndNscConnection(t, 1, kubetest.VppAgentTestingPodFixture(assert), assert)
}

func testOneTimeNseAndNscConnection(t *testing.T, nodeCount int, fixture kubetest.TestingPodFixture, assert *WithT) {
	k8s, err := kubetest.NewK8s(assert, true)
	defer k8s.Cleanup()
	Expect(err).To(BeNil())
	defer kubetest.ShowLogs(k8s, t)
	nodes := createNodes(assert, k8s, nodeCount)
	doneChannel := make(chan bool, nscCount)

	for i := 0; i < nscMaxCount; i++ {
		go func(index int) {
			defer func() {
				path := "/home/circleci/project/.tests/cloud_test/packet-1/logs/" + t.Name() + "/nsc-nse/" + strconv.Itoa(index)
				getPath := func() string {
					return path
				}
				kubetest.MakeSnapshot(k8s, t, getPath)
			}()
			nse := fixture.DeployNse(k8s, nodes[0].Node, icmpDefaultName+strconv.Itoa(index), defaultTimeout)
			nsc := fixture.DeployNsc(k8s, nodes[nodeCount-1].Node, nscDefaultName+strconv.Itoa(index), defaultTimeout)
			defer k8s.DeletePods(nse, nsc)
			fixture.CheckNsc(k8s, nsc)
			logrus.Infof("start write to channel, %v", index)
			doneChannel <- true
			logrus.Infof("done write to channel, %v", index)
		}(i)
	}

	for i := 0; i < nscMaxCount; i++ {
		logrus.Infof("done %v / %v", i, nscMaxCount)
		<-doneChannel
	}
}

func testOneTimeConnection(t *testing.T, nodeCount int, nscDeploy, icmpDeploy kubetest.PodSupplier, nsePing kubetest.NsePinger) {
	g := NewWithT(t)

	k8s, err := kubetest.NewK8s(g, true)
	defer k8s.Cleanup()

	g.Expect(err).To(BeNil())
	defer kubetest.ShowLogs(k8s, t)
	nodes := createNodes(g, k8s, nodeCount)
	icmpDeploy(k8s, nodes[nodeCount-1].Node, icmpDefaultName, defaultTimeout)

	doneChannel := make(chan nscPingResult, nscCount)
	defer close(doneChannel)

	for count := nscCount; count > 0; count-- {
		go createNscAndPingIcmp(g, k8s, count, nodes[0].Node, doneChannel, nscDeploy, nsePing)
	}

	for count := nscCount; count > 0; count-- {
		nscPingResult := <-doneChannel
		g.Expect(nscPingResult.success).To(Equal(true))
	}
}

func testMovingConnection(t *testing.T, nodeCount int, nscDeploy, icmpDeploy kubetest.PodSupplier, pingNse kubetest.NsePinger) {
	g := NewWithT(t)

	k8s, err := kubetest.NewK8s(g, true)
	defer k8s.Cleanup()

	g.Expect(err).To(BeNil())
	defer kubetest.ShowLogs(k8s, t)
	nodes := createNodes(g, k8s, nodeCount)

	icmpDeploy(k8s, nodes[nodeCount-1].Node, icmpDefaultName, defaultTimeout)
	doneChannel := make(chan nscPingResult, nscCount)
	defer close(doneChannel)

	for testCount := 0; testCount < nscMaxCount; testCount += nscCount {
		for count := nscCount; count > 0; count-- {
			go createNscAndPingIcmp(g, k8s, count, nodes[0].Node, doneChannel, nscDeploy, pingNse)
		}

		for count := nscCount; count > 0; count-- {
			nscPingResult := <-doneChannel
			g.Expect(nscPingResult.success).To(Equal(true))
			k8s.DeletePods(nscPingResult.nsc)
		}
	}
}

func testOneToOneConnection(t *testing.T, nodeCount int, nscDeploy, icmpDeploy kubetest.PodSupplier, pingNse kubetest.NsePinger) {
	g := NewWithT(t)

	k8s, err := kubetest.NewK8s(g, true)
	defer k8s.Cleanup()

	g.Expect(err).To(BeNil())
	defer kubetest.ShowLogs(k8s, t)
	nodes := createNodes(g, k8s, nodeCount)
	doneChannel := make(chan nscPingResult, 1)
	defer close(doneChannel)

	for testCount := 0; testCount < nscMaxCount; testCount += nscCount {
		icmp := icmpDeploy(k8s, nodes[nodeCount-1].Node, icmpDefaultName, defaultTimeout)
		createNscAndPingIcmp(g, k8s, 1, nodes[0].Node, doneChannel, nscDeploy, pingNse)
		result := <-doneChannel
		g.Expect(result.success).To(Equal(true))
		k8s.DeletePods(icmp, result.nsc)
	}
}

type nscPingResult struct {
	success bool
	nsc     *v1.Pod
}

func createNodes(g *WithT, k8s *kubetest.K8s, count int) []*kubetest.NodeConf {
	g.Expect(count > 0 && count < 3).Should(Equal(true))
	nodes, err := kubetest.SetupNodesConfig(k8s, count, defaultTimeout, kubetest.NoHealNSMgrPodConfig(k8s), k8s.GetK8sNamespace())
	g.Expect(err).To(BeNil())

	g.Expect(len(nodes), count)
	return nodes
}

func createNscAndPingIcmp(g *WithT, k8s *kubetest.K8s, id int, node *v1.Node, done chan nscPingResult, nscDeploy kubetest.PodSupplier, pingNse kubetest.NsePinger) {
	nsc := nscDeploy(k8s, node, nscDefaultName+strconv.Itoa(id), defaultTimeout)
	g.Expect(nsc.Name).To(Equal(nscDefaultName + strconv.Itoa(id)))
	done <- nscPingResult{
		success: pingNse(k8s, nsc),
		nsc:     nsc,
	}
}
