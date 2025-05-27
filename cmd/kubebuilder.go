package cmd

import (
	"github.com/charmbracelet/log"
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

func Run() {
	// Bundle plugin which built the golang projects scaffold with base.go/v4 and kustomize/v2 plugins
	gov4Bundle, _ := plugin.NewBundleWithOptions(plugin.WithName(golang.DefaultNameQualifier),
		plugin.WithVersion(plugin.Version{Number: 4}),
		plugin.WithPlugins(kustomizecommonv2.Plugin{}, golangv4.Plugin{}),
	)

	fs := machinery.Filesystem{
		FS: afero.NewOsFs(),
	}
	externalPlugins, err := cli.DiscoverExternalPlugins(fs.FS)
	if err != nil {
		log.Error(err)
	}

	c, err := cli.New(
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
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}
