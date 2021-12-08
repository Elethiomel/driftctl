package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type BlockstorageVolumeEnumerator struct {
	repository repository.CinderRepository
	factory    resource.ResourceFactory
}

func NewBlockstorageVolumeV2Enumerator(repo repository.CinderRepository, factory resource.ResourceFactory) *BlockstorageVolumeEnumerator {
	return &BlockstorageVolumeEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *BlockstorageVolumeEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackBlockstorageVolumeV2ResourceType
}

func (e *BlockstorageVolumeEnumerator) Enumerate() ([]*resource.Resource, error) {
	volumes, err := e.repository.ListAllVolumes()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(volumes))

	for _, volume := range volumes {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				volume,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
