package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type ComputeInstanceEnumerator struct {
	repository repository.NovaRepository
	factory    resource.ResourceFactory
}

func NewComputeInstanceV2Enumerator(repo repository.NovaRepository, factory resource.ResourceFactory) *ComputeInstanceEnumerator {
	return &ComputeInstanceEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *ComputeInstanceEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackComputeInstanceV2ResourceType
}

func (e *ComputeInstanceEnumerator) Enumerate() ([]*resource.Resource, error) {
	instances, err := e.repository.ListAllInstances()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(instances))

	for _, instance := range instances {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				instance,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
