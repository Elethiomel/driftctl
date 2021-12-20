package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackComputeKeypairV2ResourceType = "openstack_compute_keypair_v2"

func initOpenStackComputeKeypairV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackComputeKeypairV2ResourceType, func(res *resource.Resource) {
	})
}
