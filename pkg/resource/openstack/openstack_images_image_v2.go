package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/sirupsen/logrus"
)

const OpenStackImagesImageV2ResourceType = "openstack_images_image_v2"

func initOpenStackImagesImageV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackImagesImageV2ResourceType, func(res *resource.Resource) {
		logrus.Infof("FLAVOR %+v\n", res)
	})
}
