package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type NetworkingRouterInterfaceEnumerator struct {
	repository repository.NeutronRepository
	factory    resource.ResourceFactory
}

func NewNetworkingRouterInterfaceV2Enumerator(repo repository.NeutronRepository, factory resource.ResourceFactory) *NetworkingRouterInterfaceEnumerator {
	return &NetworkingRouterInterfaceEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *NetworkingRouterInterfaceEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackNetworkingRouterInterfaceV2ResourceType
}

func (e *NetworkingRouterInterfaceEnumerator) Enumerate() ([]*resource.Resource, error) {
	routerInterfaces, err := e.repository.ListAllRouterInterfaces()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(routerInterfaces))

	for _, routerInterface := range routerInterfaces {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				routerInterface,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
