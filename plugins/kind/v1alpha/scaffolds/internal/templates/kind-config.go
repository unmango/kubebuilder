package templates

import "sigs.k8s.io/kubebuilder/v4/pkg/machinery"

var _ machinery.Template = &KindConfig{}

type KindConfig struct {
	machinery.TemplateMixin
	machinery.ProjectNameMixin
}

// SetTemplateDefaults implements machinery.Template.
func (k *KindConfig) SetTemplateDefaults() error {
	k.TemplateBody = kindConfigTemplate
	return nil
}

const kindConfigTemplate = `kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
name: {{ .ProjectName }}
networking:
  dnsSearch: []
`
