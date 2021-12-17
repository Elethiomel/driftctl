package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/sirupsen/logrus"
)

const OpenStackNetworkingNetworkV2ResourceType = "openstack_networking_network_v2"

func initOpenStackNetworkingNetworkV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackNetworkingNetworkV2ResourceType, func(res *resource.Resource) {
		logrus.Infof("FLAVOR %+v\n", res)
	})
}
