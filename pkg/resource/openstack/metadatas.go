package openstack

import "github.com/cloudskiff/driftctl/pkg/resource"

func InitResourcesMetadata(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	initOpenStackComputeKeypairV2MetaData(resourceSchemaRepository)
}
