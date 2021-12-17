package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type ComputeVolumeAttachEnumerator struct {
	repository repository.NovaRepository
	factory    resource.ResourceFactory
}

func NewComputeVolumeAttachV2Enumerator(repo repository.NovaRepository, factory resource.ResourceFactory) *ComputeVolumeAttachEnumerator {
	return &ComputeVolumeAttachEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *ComputeVolumeAttachEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackComputeVolumeAttachV2ResourceType
}

func (e *ComputeVolumeAttachEnumerator) Enumerate() ([]*resource.Resource, error) {
	volumeAttachments, err := e.repository.ListAllVolumeAttachments()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(volumeAttachments))

	for _, volumeAttachment := range volumeAttachments {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				volumeAttachment,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
