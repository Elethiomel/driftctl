package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackComputeInstanceV2ResourceType = "openstack_compute_instance_v2"

func initOpenStackComputeInstanceV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackComputeInstanceV2ResourceType, func(res *resource.Resource) {
	})
}
