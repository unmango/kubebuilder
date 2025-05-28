package v1alpha

import (
	"fmt"

	"github.com/unmango/kubebuilder/plugins/kind/v1alpha/scaffolds"
	"sigs.k8s.io/kubebuilder/v4/pkg/config"
	"sigs.k8s.io/kubebuilder/v4/pkg/machinery"
	"sigs.k8s.io/kubebuilder/v4/pkg/plugin"
)

var _ plugin.InitSubcommand = &initSubcommand{}

type initSubcommand struct {
	config config.Config
}

func (i *initSubcommand) InjectConfig(c config.Config) error {
	i.config = c
	return nil
}

// Scaffold implements plugin.InitSubcommand.
func (i *initSubcommand) Scaffold(fs machinery.Filesystem) error {
	scaffolder := scaffolds.NewInitScaffolder(i.config)
	scaffolder.InjectFS(fs)
	if err := scaffolder.Scaffold(); err != nil {
		return fmt.Errorf("failed to scaffold init subcommand: %w", err)
	}

	return nil
}
