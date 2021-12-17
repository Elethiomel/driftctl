package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type NetworkingNetworkEnumerator struct {
	repository repository.NeutronRepository
	factory    resource.ResourceFactory
}

func NewNetworkingNetworkV2Enumerator(repo repository.NeutronRepository, factory resource.ResourceFactory) *NetworkingNetworkEnumerator {
	return &NetworkingNetworkEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *NetworkingNetworkEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackNetworkingNetworkV2ResourceType
}

func (e *NetworkingNetworkEnumerator) Enumerate() ([]*resource.Resource, error) {
	networks, err := e.repository.ListAllNetworks()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(networks))

	for _, network := range networks {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				network,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
