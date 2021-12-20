package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackComputeSecgroupV2ResourceType = "openstack_compute_secgroup_v2"

func initOpenStackComputeSecgroupV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackComputeSecgroupV2ResourceType, func(res *resource.Resource) {
	})
}
