package cli

import (
	kindv1alpha "github.com/unmango/kubebuilder/plugins/kind/v1alpha"
	"github.com/unmango/kubebuilder/upstream"
	"sigs.k8s.io/kubebuilder/v4/pkg/cli"
)

func New() (*cli.CLI, error) {
	coreOptions, err := upstream.CliOptions()
	if err != nil {
		return nil, err
	}

	return cli.New(
		coreOptions,
		cli.WithPlugins(&kindv1alpha.Plugin{}),
	)
}
