package openstack

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const OpenStackNetworkingSecgroupRuleV2ResourceType = "openstack_networking_secgroup_rule_v2"

func initOpenStackNetworkingSecgroupRuleV2MetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(OpenStackNetworkingSecgroupRuleV2ResourceType, func(res *resource.Resource) {
	})
}
