package upstream

import (
	"github.com/spf13/afero"
	"sigs.k8s.io/kubebuilder/v4/pkg/cli"
	cfgv3 "sigs.k8s.io/kubebuilder/v4/pkg/config/v3"
	"sigs.k8s.io/kubebuilder/v4/pkg/machinery"
	"sigs.k8s.io/kubebuilder/v4/pkg/plugin"
	kustomizecommonv2 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/common/kustomize/v2"
	"sigs.k8s.io/kubebuilder/v4/pkg/plugins/golang"
	deployimagev1alpha1 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/golang/deploy-image/v1alpha1"
	golangv4 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/golang/v4"
	grafanav1alpha1 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/optional/grafana/v1alpha"
	helmv1alpha1 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/optional/helm/v1alpha"
)

// https://github.com/kubernetes-sigs/kubebuilder/blob/master/cmd/cmd.go

func CliOptions() (cli.Option, error) {
	gov4Bundle, err := plugin.NewBundleWithOptions(plugin.WithName(golang.DefaultNameQualifier),
		plugin.WithVersion(plugin.Version{Number: 4}),
		plugin.WithPlugins(kustomizecommonv2.Plugin{}, golangv4.Plugin{}),
	)
	if err != nil {
		return nil, err
	}

	fs := machinery.Filesystem{
		FS: afero.NewOsFs(),
	}
	externalPlugins, err := cli.DiscoverExternalPlugins(fs.FS)
	if err != nil {
		return nil, err
	}

	opts := []cli.Option{
		cli.WithCommandName("kubebuilder"),
		// cli.WithVersion(versionString()),
		// cli.WithCliVersion(getKubebuilderVersion()),
		cli.WithPlugins(
			golangv4.Plugin{},
			gov4Bundle,
			&kustomizecommonv2.Plugin{},
			&deployimagev1alpha1.Plugin{},
			&grafanav1alpha1.Plugin{},
			&helmv1alpha1.Plugin{},
		),
		cli.WithPlugins(externalPlugins...),
		cli.WithDefaultPlugins(cfgv3.Version, gov4Bundle),
		cli.WithDefaultProjectVersion(cfgv3.Version),
		cli.WithCompletion(),
	}

	return fold(opts), nil
}

func fold(opts []cli.Option) cli.Option {
	return func(c *cli.CLI) error {
		for _, apply := range opts {
			if err := apply(c); err != nil {
				return err
			}
		}

		return nil
	}
}
