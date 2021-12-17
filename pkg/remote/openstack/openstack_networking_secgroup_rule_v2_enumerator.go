package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type NetworkingSecgroupRuleEnumerator struct {
	repository repository.NeutronRepository
	factory    resource.ResourceFactory
}

func NewNetworkingSecgroupRuleV2Enumerator(repo repository.NeutronRepository, factory resource.ResourceFactory) *NetworkingSecgroupRuleEnumerator {
	return &NetworkingSecgroupRuleEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *NetworkingSecgroupRuleEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackNetworkingSecgroupRuleV2ResourceType
}

func (e *NetworkingSecgroupRuleEnumerator) Enumerate() ([]*resource.Resource, error) {
	secgroupRules, err := e.repository.ListAllSecgroupRules()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(secgroupRules))

	for _, secgroupRule := range secgroupRules {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				secgroupRule,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
