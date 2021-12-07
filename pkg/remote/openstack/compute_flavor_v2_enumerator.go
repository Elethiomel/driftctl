package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type ComputeFlavorEnumerator struct {
	repository repository.NovaRepository
	factory    resource.ResourceFactory
}

func NewComputeFlavorV2Enumerator(repo repository.NovaRepository, factory resource.ResourceFactory) *ComputeFlavorEnumerator {
	return &ComputeFlavorEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *ComputeFlavorEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackComputeFlavorV2ResourceType
}

func (e *ComputeFlavorEnumerator) Enumerate() ([]*resource.Resource, error) {
	flavors, err := e.repository.ListAllFlavors()
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
