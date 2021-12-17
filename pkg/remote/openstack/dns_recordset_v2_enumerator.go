package openstack

import (
	remoteerror "github.com/cloudskiff/driftctl/pkg/remote/error"
	"github.com/cloudskiff/driftctl/pkg/remote/openstack/repository"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/cloudskiff/driftctl/pkg/resource/openstack"
)

type DNSRecordsetEnumerator struct {
	repository repository.DesignateRepository
	factory    resource.ResourceFactory
}

func NewDNSRecordsetV2Enumerator(repo repository.DesignateRepository, factory resource.ResourceFactory) *DNSRecordsetEnumerator {
	return &DNSRecordsetEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *DNSRecordsetEnumerator) SupportedType() resource.ResourceType {
	return openstack.OpenStackDNSRecordsetV2ResourceType
}

func (e *DNSRecordsetEnumerator) Enumerate() ([]*resource.Resource, error) {
	recordsets, err := e.repository.ListAllRecordsets()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(recordsets))

	for _, recordset := range recordsets {
		results = append(
			results,
			e.factory.CreateAbstractResource(
				string(e.SupportedType()),
				recordset,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
