package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type NetworkingSecgroupEnumerator struct {
	repository repository.NeutronRepository
	factory    resource.ResourceFactory
}

func NewNetworkingSecgroupV2Enumerator(repo repository.NeutronRepository, factory resource.ResourceFactory) *NetworkingSecgroupEnumerator {
	return &NetworkingSecgroupEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *NetworkingSecgroupEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackNetworkingSecgroupV2ResourceType
}

func (e *NetworkingSecgroupEnumerator) Enumerate() ([]*resource.Resource, error) {
	secgroups, err := e.repository.ListAllSecgroups()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(secgroups))

	for _, secgroup := range secgroups {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				secgroup,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
