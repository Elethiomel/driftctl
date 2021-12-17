package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type DNSZoneEnumerator struct {
	repository repository.DesignateRepository
	factory    resource.ResourceFactory
}

func NewDNSZoneV2Enumerator(repo repository.DesignateRepository, factory resource.ResourceFactory) *DNSZoneEnumerator {
	return &DNSZoneEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *DNSZoneEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackDNSZoneV2ResourceType
}

func (e *DNSZoneEnumerator) Enumerate() ([]*resource.Resource, error) {
	zones, err := e.repository.ListAllZones()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(zones))

	for _, zone := range zones {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				zone,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
