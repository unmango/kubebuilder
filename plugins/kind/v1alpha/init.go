package v1alpha

import "sigs.k8s.io/kubebuilder/v4/pkg/machinery"

type initSubcommand struct{}

// Scaffold implements plugin.InitSubcommand.
func (i *initSubcommand) Scaffold(fs machinery.Filesystem) error {
	panic("unimplemented")
}
