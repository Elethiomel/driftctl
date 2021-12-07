package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/sirupsen/logrus"
)

const OpenStackComputeKeypairV2ResourceType = "openstack_compute_keypair_v2"

func initOpenStackComputeKeypairV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackComputeKeypairV2ResourceType, func(res *resource.Resource) {
		logrus.Infof("KEYPAIR %+v\n", res)
	})
}
