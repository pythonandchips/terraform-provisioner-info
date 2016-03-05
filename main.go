package main

import (
	"fmt"

	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func ResourceProvisionerBuilder() terraform.ResourceProvisioner {
	return &ResourceProvisioner{}
}

func main() {
	serveOpts := &plugin.ServeOpts{
		ProvisionerFunc: ResourceProvisionerBuilder,
	}

	plugin.Serve(serveOpts)
}

type ResourceProvisioner struct{}

func (p *ResourceProvisioner) Apply(o terraform.UIOutput, s *terraform.InstanceState, c *terraform.ResourceConfig) error {

	o.Output("INSTANCE STATE\n")
	o.Output(fmt.Sprintf("Instance ID: %s\n", s.ID))
	o.Output("Attributes:\n")
	for k, v := range s.Attributes {
		o.Output(fmt.Sprintf("%s : %s \n", k, v))
	}

	o.Output("Meta:\n")
	for k, v := range s.Meta {
		o.Output(fmt.Sprintf("%s : %s \n", k, v))
	}
	o.Output("----------------------------------------\n")

	o.Output("RESOURCE CONFIG\n")
	for k, v := range c.Config {
		o.Output(fmt.Sprintf("%s : %s \n", k, v))
	}
	o.Output("----------------------------------------\n")

	o.Output("RESOURCE RAW\n")
	for k, v := range c.Raw {
		o.Output(fmt.Sprintf("%s : %s \n", k, v))
	}
	o.Output("----------------------------------------\n")

	o.Output("RESOURCE COMUTED\n")
	for k, v := range c.ComputedKeys {
		o.Output(fmt.Sprintf("%s : %s \n", k, v))
	}
	o.Output("----------------------------------------\n")

	return nil
}

func (p *ResourceProvisioner) Validate(c *terraform.ResourceConfig) ([]string, []error) {
	return []string{}, nil
}
