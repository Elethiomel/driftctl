package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type ComputeKeypairEnumerator struct {
	repository repository.NovaRepository
	factory    resource.ResourceFactory
}

func NewComputeKeypairV2Enumerator(repo repository.NovaRepository, factory resource.ResourceFactory) *ComputeKeypairEnumerator {
	return &ComputeKeypairEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *ComputeKeypairEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackComputeKeypairV2ResourceType
}

func (e *ComputeKeypairEnumerator) Enumerate() ([]*resource.Resource, error) {
	keyPairs, err := e.repository.ListAllKeypairs()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(keyPairs))

	for _, keyPair := range keyPairs {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				keyPair,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
