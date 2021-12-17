package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/sirupsen/logrus"
)

const OpenStackComputeInterfaceAttachV2ResourceType = "openstack_compute_interface_attach_v2"

func initOpenStackComputeInterfaceAttachV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackComputeInterfaceAttachV2ResourceType, func(res *resource.Resource) {
		logrus.Infof("KEYPAIR %+v\n", res)
	})
}
