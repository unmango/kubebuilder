package cli

import (
	kindv1alpha "github.com/unmango/kubebuilder/plugins/kind/v1alpha"
	"github.com/unmango/kubebuilder/plugins/unmango"
	"github.com/unmango/kubebuilder/upstream"
	"sigs.k8s.io/kubebuilder/v4/pkg/cli"
	"sigs.k8s.io/kubebuilder/v4/pkg/model/stage"
	"sigs.k8s.io/kubebuilder/v4/pkg/plugin"
)

func New() (*cli.CLI, error) {
	coreOptions, err := upstream.CliOptions()
	if err != nil {
		return nil, err
	}

	gov4Bundle, err := upstream.GoV4Bundle()
	if err != nil {
		return nil, err
	}

	unmangov1alphaBundle, err := plugin.NewBundleWithOptions(
		plugin.WithName(unmango.DefaultNameQualifier),
		plugin.WithVersion(plugin.Version{Number: 1, Stage: stage.Alpha}),
		plugin.WithPlugins(gov4Bundle, &kindv1alpha.Plugin{}),
	)
	if err != nil {
		return nil, err
	}

	return cli.New(
		coreOptions,
		cli.WithPlugins(
			&kindv1alpha.Plugin{},
			unmangov1alphaBundle,
		),
	)
}
