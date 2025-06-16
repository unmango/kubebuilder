package scaffolds

import (
	"fmt"

	"github.com/unmango/kubebuilder/plugins/kind/v1alpha/scaffolds/internal/templates"
	"sigs.k8s.io/kubebuilder/v4/pkg/config"
	"sigs.k8s.io/kubebuilder/v4/pkg/machinery"
	"sigs.k8s.io/kubebuilder/v4/pkg/plugins"
)

var _ plugins.Scaffolder = &initScaffolder{}

type initScaffolder struct {
	config config.Config
	fs     machinery.Filesystem
}

func NewInitScaffolder(cfg config.Config) plugins.Scaffolder {
	return &initScaffolder{config: cfg}
}

// InjectFS implements plugins.Scaffolder.
func (i *initScaffolder) InjectFS(fs machinery.Filesystem) {
	i.fs = fs
}

// Scaffold implements plugins.Scaffolder.
func (i *initScaffolder) Scaffold() error {
	scaffold := machinery.NewScaffold(i.fs,
		machinery.WithConfig(i.config),
	)

	templates := []machinery.Builder{
		&templates.KindConfig{},
	}

	if err := scaffold.Execute(templates...); err != nil {
		return fmt.Errorf("failed to scaffold kind config: %w", err)
	}

	return nil
}
