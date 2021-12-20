package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackNetworkingRouterV2ResourceType = "openstack_networking_router_v2"

func initOpenStackNetworkingRouterV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackNetworkingRouterV2ResourceType, func(res *resource.Resource) {
	})
}
