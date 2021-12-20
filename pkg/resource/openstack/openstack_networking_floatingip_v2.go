package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackNetworkingFloatingipV2ResourceType = "openstack_networking_floatingip_v2"

func initOpenStackNetworkingFloatingipV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackNetworkingFloatingipV2ResourceType, func(res *resource.Resource) {
	})
}
