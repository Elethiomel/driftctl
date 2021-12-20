package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackBlockstorageVolumeV2ResourceType = "openstack_blockstorage_volume_v2"

func initOpenStackBlockstorageVolumeV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackBlockstorageVolumeV2ResourceType, func(res *resource.Resource) {
	})
}
