package e2e_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/kubebuilder/v4/pkg/plugin/util"
	"sigs.k8s.io/kubebuilder/v4/test/e2e/utils"
)

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}

var _ = BeforeSuite(func() {
	kbc, err := utils.NewTestContext(util.KubebuilderBinName, "GO111MODULE=on")
	Expect(err).NotTo(HaveOccurred())
	Expect(kbc.Prepare()).To(Succeed())

	By("installing the cert-manager bundle")
	Expect(kbc.InstallCertManager()).To(Succeed())

	By("installing the prometheus operator")
	Expect(kbc.InstallPrometheusOperManager()).To(Succeed())
})

var _ = AfterSuite(func() {
	kbc, err := utils.NewTestContext(util.KubebuilderBinName, "GO111MODULE=on")
	Expect(err).NotTo(HaveOccurred())
	Expect(kbc.Prepare()).To(Succeed())

	By("uninstalling the Prometheus manager bundle")
	kbc.UninstallPrometheusOperManager()

	By("uninstalling the cert-manager bundle")
	kbc.UninstallCertManager()
})
