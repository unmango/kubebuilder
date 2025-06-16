package v1alpha

import (
	"fmt"

	"sigs.k8s.io/kubebuilder/v4/pkg/config"
	"sigs.k8s.io/kubebuilder/v4/pkg/machinery"
	"sigs.k8s.io/kubebuilder/v4/pkg/model/resource"
	"sigs.k8s.io/kubebuilder/v4/pkg/plugin"
)

var _ plugin.CreateAPISubcommand = &createAPISubcommand{}

type createAPISubcommand struct {
	config config.Config

	resource *resource.Resource
}

// InjectResource implements plugin.CreateAPISubcommand.
func (c *createAPISubcommand) InjectResource(res *resource.Resource) error {
	c.resource = res

	if err := c.resource.Validate(); err != nil {
		return fmt.Errorf("error validating resource: %w", err)
	}

	return nil
}

// Scaffold implements plugin.CreateAPISubcommand.
func (c *createAPISubcommand) Scaffold(machinery.Filesystem) error {
	return nil
}
