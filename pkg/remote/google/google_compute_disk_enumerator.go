package google

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/google/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/google"
)

type GoogleComputeDiskEnumerator struct {
	repository repository.AssetRepository
	factory    resource.ResourceFactory
}

func NewGoogleComputeDiskEnumerator(repo repository.AssetRepository, factory resource.ResourceFactory) *GoogleComputeDiskEnumerator {
	return &GoogleComputeDiskEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *GoogleComputeDiskEnumerator) SupportedType() resource.ResourceType {
	return google.GoogleComputeDiskResourceType
}

func (e *GoogleComputeDiskEnumerator) Enumerate() ([]*resource.Resource, error) {
	resources, err := e.repository.SearchAllDisks()

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
				map[string]interface{}{
					"name": res.GetDisplayName(),
				},
			),
		)
	}

	return results, err
}
