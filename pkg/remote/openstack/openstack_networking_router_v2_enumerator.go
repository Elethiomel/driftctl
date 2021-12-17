package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type NetworkingRouterEnumerator struct {
	repository repository.NeutronRepository
	factory    resource.ResourceFactory
}

func NewNetworkingRouterV2Enumerator(repo repository.NeutronRepository, factory resource.ResourceFactory) *NetworkingRouterEnumerator {
	return &NetworkingRouterEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *NetworkingRouterEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackNetworkingRouterV2ResourceType
}

func (e *NetworkingRouterEnumerator) Enumerate() ([]*resource.Resource, error) {
	routers, err := e.repository.ListAllRouters()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(routers))

	for _, router := range routers {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				router,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
