package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type NetworkingPortEnumerator struct {
	repository repository.NeutronRepository
	factory    resource.ResourceFactory
}

func NewNetworkingPortV2Enumerator(repo repository.NeutronRepository, factory resource.ResourceFactory) *NetworkingPortEnumerator {
	return &NetworkingPortEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *NetworkingPortEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackNetworkingPortV2ResourceType
}

func (e *NetworkingPortEnumerator) Enumerate() ([]*resource.Resource, error) {
	flavors, err := e.repository.ListAllPorts()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(flavors))

	for _, flavor := range flavors {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				flavor,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
