package v1alpha

import (
	"sigs.k8s.io/kubebuilder/v4/pkg/config"
	cfgv3 "sigs.k8s.io/kubebuilder/v4/pkg/config/v3"
	"sigs.k8s.io/kubebuilder/v4/pkg/model/stage"
	"sigs.k8s.io/kubebuilder/v4/pkg/plugin"
	"sigs.k8s.io/kubebuilder/v4/pkg/plugins"
)

const pluginName = "kind.unmango." + plugins.DefaultNameQualifier

var (
	pluginVersion            = plugin.Version{Number: 1, Stage: stage.Alpha}
	supportedProjectVersions = []config.Version{cfgv3.Version}
)

var (
	_ plugin.Init = Plugin{}
)

type Plugin struct {
	initSubcommand
}

// Name implements plugin.Init.
func (p Plugin) Name() string {
	return pluginName
}

// SupportedProjectVersions implements plugin.Init.
func (p Plugin) SupportedProjectVersions() []config.Version {
	return supportedProjectVersions
}

// Version implements plugin.Init.
func (p Plugin) Version() plugin.Version {
	return pluginVersion
}

// GetInitSubcommand implements plugin.Init.
func (p Plugin) GetInitSubcommand() plugin.InitSubcommand {
	return &p.initSubcommand
}
