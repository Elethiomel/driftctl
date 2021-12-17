package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type NetworkingFloatingipEnumerator struct {
	repository repository.NeutronRepository
	factory    resource.ResourceFactory
}

func NewNetworkingFloatingipV2Enumerator(repo repository.NeutronRepository, factory resource.ResourceFactory) *NetworkingFloatingipEnumerator {
	return &NetworkingFloatingipEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *NetworkingFloatingipEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackNetworkingFloatingipV2ResourceType
}

func (e *NetworkingFloatingipEnumerator) Enumerate() ([]*resource.Resource, error) {
	floatingips, err := e.repository.ListAllFloatingips()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(floatingips))

	for _, floatingip := range floatingips {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				floatingip,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
