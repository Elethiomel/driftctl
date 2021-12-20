package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackNetworkingSubnetV2ResourceType = "openstack_networking_subnet_v2"

func initOpenStackNetworkingSubnetV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackNetworkingSubnetV2ResourceType, func(res *resource.Resource) {
	})
}
