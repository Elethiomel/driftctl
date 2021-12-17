package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type ImagesImageEnumerator struct {
	repository repository.GlanceRepository
	factory    resource.ResourceFactory
}

func NewImagesImageV2Enumerator(repo repository.GlanceRepository, factory resource.ResourceFactory) *ImagesImageEnumerator {
	return &ImagesImageEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *ImagesImageEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackImagesImageV2ResourceType
}

func (e *ImagesImageEnumerator) Enumerate() ([]*resource.Resource, error) {
	images, err := e.repository.ListAllImages()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(images))

	for _, image := range images {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				image,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
