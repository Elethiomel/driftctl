package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type NetworkingSubnetEnumerator struct {
	repository repository.NeutronRepository
	factory    resource.ResourceFactory
}

func NewNetworkingSubnetV2Enumerator(repo repository.NeutronRepository, factory resource.ResourceFactory) *NetworkingSubnetEnumerator {
	return &NetworkingSubnetEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *NetworkingSubnetEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackNetworkingSubnetV2ResourceType
}

func (e *NetworkingSubnetEnumerator) Enumerate() ([]*resource.Resource, error) {
	subnets, err := e.repository.ListAllSubnets()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(subnets))

	for _, subnet := range subnets {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				subnet,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
