package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/alerter"
	"github.com/cloudskiff/driftctl/pkg/output"
	"github.com/cloudskiff/driftctl/pkg/remote/cache"
	"github.com/cloudskiff/driftctl/pkg/remote/common"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
	"github.com/cloudskiff/driftctl/pkg/terraform"
	gcos "github.com/gophercloud/gophercloud/openstack"
	"github.com/sirupsen/logrus"
)

const RemoteOpenStackTerraform = "aws+tf"

/**
 * Initialize remote (configure credentials, launch tf providers and start gRPC clients)
 * Required to use Scanner
 */
func Init(
	// Version required by the user
	version string,
	// Util to send alert
	alerter *alerter.Alerter,
	// Library that contains all providers
	providerLibrary *terraform.ProviderLibrary,
	// Library that contains enumerators and details fetchers for each supported resources
	remoteLibrary *common.RemoteLibrary,
	// Progress displayer
	progress output.Progress,
	// Repository for all resource schemas
	resourceSchemaRepository *resource.SchemaRepository,
	// Factory used to create driftctl resource
	factory resource.ResourceFactory,
	// driftctl configuration directory (where Terraform provider is downloaded)
	configDir string) error {

	// You need to define the default version of the Terraform provider when the user does not specify one
	if version == "" {
		version = "1.45.0"
	}

	// Creation of the Terraform provider
	provider, err := NewOpenStackTerraformProvider(version, progress, configDir)
	if err != nil {
		return err
	}
	// And then initialization
	err = provider.Init()
	if err != nil {
		return err
	}

	opts, err := gcos.AuthOptionsFromEnv()
	if err != nil {
		logrus.Fatalf("Could not load Openstack auth options from environment : %s", err)
	}

	gopherProvider, err := gcos.AuthenticatedClient(opts)
	if err != nil {
		logrus.Fatalf("Could not authenticate with Openstack : %s", err)
	}

	// You'll need to create a new cache that will be used to cache fetched lists of resources
	repositoryCache := cache.New(100)
	NovaRepository := repository.NewNovaRepository(gopherProvider, repositoryCache)
	logrus.Infof("NovaRepository %+v", NovaRepository)
	// Deserializer is used to convert cty value returned by Terraform provider to driftctl Resource
	//	deserializer := resource.NewDeserializer(factory)

	// Adding the provider to the library
	providerLibrary.AddProvider(terraform.OPENSTACK, provider)

	remoteLibrary.AddEnumerator(NewComputeKeypairV2Enumerator(NovaRepository, factory))
	remoteLibrary.AddEnumerator(NewComputeFlavorV2Enumerator(NovaRepository, factory))
	remoteLibrary.AddEnumerator(NewComputeInstanceV2Enumerator(NovaRepository, factory))
	remoteLibrary.AddEnumerator(NewComputeSecgroupV2Enumerator(NovaRepository, factory))

	err = resourceSchemaRepository.Init(terraform.OPENSTACK, provider.Version(), provider.Schema())
	if err != nil {
		return err
	}

	openstack.InitResourcesMetadata(resourceSchemaRepository)

	return nil
}
