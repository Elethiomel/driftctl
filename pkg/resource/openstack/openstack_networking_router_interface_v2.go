package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackNetworkingRouterInterfaceV2ResourceType = "openstack_networking_router_interface_v2"

func initOpenStackNetworkingRouterInterfaceV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackNetworkingRouterInterfaceV2ResourceType, func(res *resource.Resource) {
	})
}
