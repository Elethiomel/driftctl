package google

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/google/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/google"
)

type GoogleComputeRouterEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeRouterEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeRouterEnumerator {
	return &GoogleComputeRouterEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *GoogleComputeRouterEnumerator) SupportedType() resource.ResourceType {
	return google.GoogleComputeRouterResourceType
}

func (e *GoogleComputeRouterEnumerator) Enumerate() ([]*resource.Resource, error) {
	resources, err := e.repository.SearchAllRouters()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(resources))

	for _, res := range resources {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				trimResourceName(res.GetName()),
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
