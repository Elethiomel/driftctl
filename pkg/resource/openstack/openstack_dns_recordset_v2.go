package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackDNSRecordsetV2ResourceType = "openstack_dns_recordset_v2"

func initOpenStackDNSRecordsetV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackDNSRecordsetV2ResourceType, func(res *resource.Resource) {
	})
}
