package e2e_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/kubebuilder/v4/pkg/plugin/util"
	"sigs.k8s.io/kubebuilder/v4/test/e2e/utils"
)

var _ = Describe("kubebuilder", func() {
	Context("plugin kind/v1alpha", func() {
		var kbc *utils.TestContext

		BeforeEach(func() {
			var err error
			kbc, err = utils.NewTestContext(util.KubebuilderBinName, "GO111MODULE=on")
			Expect(err).NotTo(HaveOccurred())
			Expect(kbc.Prepare()).To(Succeed())
		})

		AfterEach(func() {
			kbc.Destroy()
		})

		It("should generate a runnable project", func() {
			By("Initializing the project")
			err := kbc.Init(
				"--plugins", "kind.unmango.kubebuilder.io/v1-alpha",
				"--project-version", "3",
				"--domain", kbc.Domain,
			)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
