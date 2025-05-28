package upstream

import (
	"github.com/spf13/afero"
	"sigs.k8s.io/kubebuilder/v4/pkg/cli"
	cfgv3 "sigs.k8s.io/kubebuilder/v4/pkg/config/v3"
	"sigs.k8s.io/kubebuilder/v4/pkg/machinery"
	kustomizecommonv2 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/common/kustomize/v2"
	deployimagev1alpha1 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/golang/deploy-image/v1alpha1"
	golangv4 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/golang/v4"
	grafanav1alpha1 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/optional/grafana/v1alpha"
	helmv1alpha1 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/optional/helm/v1alpha"
)

// https://github.com/kubernetes-sigs/kubebuilder/blob/master/cmd/cmd.go

func CliOptions() ([]cli.Option, error) {
	gov4Bundle, err := NewGoV4Bundle()
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

	return []cli.Option{
		cli.WithCommandName("kubebuilder"),
		// cli.WithVersion(versionString()),
		// cli.WithCliVersion(getKubebuilderVersion()),
		cli.WithPlugins(
			golangv4.Plugin{},
			// gov4Bundle,
			&kustomizecommonv2.Plugin{},
			&deployimagev1alpha1.Plugin{},
			&grafanav1alpha1.Plugin{},
			&helmv1alpha1.Plugin{},
		),
		cli.WithPlugins(externalPlugins...),
		cli.WithDefaultPlugins(cfgv3.Version, gov4Bundle),
		cli.WithDefaultProjectVersion(cfgv3.Version),
		cli.WithCompletion(),
	}, nil
}
