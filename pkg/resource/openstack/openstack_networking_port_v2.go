package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackNetworkingPortV2ResourceType = "openstack_networking_port_v2"

func initOpenStackNetworkingPortV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackNetworkingPortV2ResourceType, func(res *resource.Resource) {
	})
}
