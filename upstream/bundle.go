package upstream

import (
	"sigs.k8s.io/kubebuilder/v4/pkg/plugin"
	kustomizecommonv2 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/common/kustomize/v2"
	"sigs.k8s.io/kubebuilder/v4/pkg/plugins/golang"
	golangv4 "sigs.k8s.io/kubebuilder/v4/pkg/plugins/golang/v4"
)

func NewGoV4Bundle() (plugin.Bundle, error) {
	return plugin.NewBundleWithOptions(plugin.WithName(golang.DefaultNameQualifier),
		plugin.WithVersion(plugin.Version{Number: 4}),
		plugin.WithPlugins(kustomizecommonv2.Plugin{}, golangv4.Plugin{}),
	)
}
