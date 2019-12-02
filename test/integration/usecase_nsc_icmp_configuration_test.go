// +build single_cluster_suite

package nsmd_integration_tests

import (
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"

	"github.com/networkservicemesh/networkservicemesh/test/kubetest"
)

func TestNSCAndICMPLocal(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	testNSCAndICMP(t, 1, false)
}

func TestNSCAndICMPRemote(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	testNSCAndICMP(t, 2, false)
}

func TestNSCAndICMPWebhookLocal(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	testNSCAndICMP(t, 1, true)
}

func TestNSCAndICMPWebhookRemote(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	testNSCAndICMP(t, 2, true)
}

/**
If passed 1 both will be on same node, if not on different.
*/
func testNSCAndICMP(t *testing.T, nodesCount int, useWebhook bool) {
	g := NewWithT(t)

	k8s, err := kubetest.NewK8s(g, kubetest.DefaultClear)
	defer k8s.Cleanup()
	defer k8s.ProcessArtifacts(t)
	g.Expect(err).To(BeNil())

	if useWebhook {
		awc, awDeployment, awService := kubetest.DeployAdmissionWebhook(k8s, "nsm-admission-webhook", "networkservicemesh/admission-webhook", k8s.GetK8sNamespace(), defaultTimeout)
		defer kubetest.DeleteAdmissionWebhook(k8s, "nsm-admission-webhook-certs", awc, awDeployment, awService, k8s.GetK8sNamespace())
	}

	nodesSetup, err := kubetest.SetupNodes(k8s, nodesCount, defaultTimeout)
	g.Expect(err).To(BeNil())

	// Run ICMP on latest node
	_ = kubetest.DeployICMP(k8s, nodesSetup[nodesCount-1].Node, "icmp-responder-nse-1", defaultTimeout)

	var nscPodNode *v1.Pod
	if useWebhook {
		nscPodNode = kubetest.DeployNSCWebhook(k8s, nodesSetup[0].Node, "nsc-1", defaultTimeout)
	} else {
		nscPodNode = kubetest.DeployNSC(k8s, nodesSetup[0].Node, "nsc-1", defaultTimeout)
	}

	kubetest.CheckNSC(k8s, nscPodNode)
}
