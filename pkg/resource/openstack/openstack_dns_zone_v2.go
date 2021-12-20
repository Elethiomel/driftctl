package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackDNSZoneV2ResourceType = "openstack_dns_zone_v2"

func initOpenStackDNSZoneV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackDNSZoneV2ResourceType, func(res *resource.Resource) {
	})
}
