package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type ComputeInterfaceAttachEnumerator struct {
	repository repository.NovaRepository
	factory    resource.ResourceFactory
}

func NewComputeInterfaceAttachV2Enumerator(repo repository.NovaRepository, factory resource.ResourceFactory) *ComputeInterfaceAttachEnumerator {
	return &ComputeInterfaceAttachEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *ComputeInterfaceAttachEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackComputeInterfaceAttachV2ResourceType
}

func (e *ComputeInterfaceAttachEnumerator) Enumerate() ([]*resource.Resource, error) {
	interfaceAttachments, err := e.repository.ListAllInterfaceAttachments()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(interfaceAttachments))

	for _, interfaceAttachment := range interfaceAttachments {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				interfaceAttachment,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
