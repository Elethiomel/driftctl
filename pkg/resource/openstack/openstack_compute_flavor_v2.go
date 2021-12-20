package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackComputeFlavorV2ResourceType = "openstack_compute_flavor_v2"

func initOpenStackComputeFlavorV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackComputeFlavorV2ResourceType, func(res *resource.Resource) {
	})
}
