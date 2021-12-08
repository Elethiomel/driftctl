package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type ComputeSecgroupEnumerator struct {
	repository repository.NovaRepository
	factory    resource.ResourceFactory
}

func NewComputeSecgroupV2Enumerator(repo repository.NovaRepository, factory resource.ResourceFactory) *ComputeSecgroupEnumerator {
	return &ComputeSecgroupEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *ComputeSecgroupEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackComputeSecgroupV2ResourceType
}

func (e *ComputeSecgroupEnumerator) Enumerate() ([]*resource.Resource, error) {
	secgroups, err := e.repository.ListAllSecgroups()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(secgroups))

	for _, secgroup := range secgroups {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				secgroup,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
