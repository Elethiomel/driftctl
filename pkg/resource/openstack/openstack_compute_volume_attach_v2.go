package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/sirupsen/logrus"
)

const OpenStackComputeVolumeAttachV2ResourceType = "openstack_compute_volume_attach_v2"

func initOpenStackComputeVolumeAttachV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackComputeVolumeAttachV2ResourceType, func(res *resource.Resource) {
		logrus.Infof("KEYPAIR %+v\n", res)
	})
}
