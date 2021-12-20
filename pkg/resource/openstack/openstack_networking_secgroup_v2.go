package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackNetworkingSecgroupV2ResourceType = "openstack_networking_secgroup_v2"

func initOpenStackNetworkingSecgroupV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackNetworkingSecgroupV2ResourceType, func(res *resource.Resource) {
	})
}
