package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/output"
	"github.com/cloudskiff/driftctl/pkg/remote/terraform"
	tf "github.com/cloudskiff/driftctl/pkg/terraform"
)

// Define your actual provider representation, it is required to compose with terraform.TerraformProvider, a name and a version
// Please note that the name should match the real Terraform provider name.
type OpenStackTerraformProvider struct {
	*terraform.TerraformProvider
	name    string
	version string
}

type openstackConfig struct {
	Token        string
	Owner        string `cty:"owner"`
	Organization string
}

func NewOpenStackTerraformProvider(version string, progress output.Progress, configDir string) (*OpenStackTerraformProvider, error) {
	// Just pass your version and name
	p := &OpenStackTerraformProvider{
		version: version,
		name:    "openstack",
	}
	// Use Terraform ProviderInstaller to retrieve the provider if needed
	installer, err := tf.NewProviderInstaller(tf.ProviderConfig{
		Key:       p.name,
		Version:   version,
		ConfigDir: configDir,
	})
	if err != nil {
		return nil, err
	}

	// ProviderConfig is dependent on the Terraform provider needs.
	tfProvider, err := terraform.NewTerraformProvider(installer, terraform.TerraformProviderConfig{
		Name:         p.name,
		DefaultAlias: p.GetConfig().getDefaultOwner(),
		GetProviderConfig: func(owner string) interface{} {
			return openstackConfig{
				Owner: p.GetConfig().getDefaultOwner(),
			}
		},
	}, progress)
	if err != nil {
		return nil, err
	}
	p.TerraformProvider = tfProvider
	return p, err
}

func (a *OpenStackTerraformProvider) Name() string {
	return a.name
}

func (p *OpenStackTerraformProvider) Version() string {
	return p.version
}

func (c openstackConfig) getDefaultOwner() string {
	if c.Organization != "" {
		return c.Organization
	}
	return c.Owner
}

func (p OpenStackTerraformProvider) GetConfig() openstackConfig {
	return openstackConfig{
		Token:        "token",
		Owner:        "owner",
		Organization: "org",
	}
}
