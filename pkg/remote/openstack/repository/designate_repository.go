package repository

import (
	"fmt"

	"github.com/cloudskiff/driftctl/pkg/remote/cache"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/zones"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/sirupsen/logrus"
)

type DesignateRepository interface {
	ListAllRecordsets() ([]string, error)
	ListAllZones() ([]string, error)
}

type designateRepository struct {
	client *gophercloud.ServiceClient
	cache  cache.Cache
}

func NewDesignateRepository(providerClient *gophercloud.ProviderClient, c cache.Cache) *designateRepository {
	client, err := openstack.NewDNSV2(providerClient, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Fatalf("Could not create Openstack Client for Designate : %s", err)
	}

	return &designateRepository{
		client,
		c,
	}
}

func (r *designateRepository) ListAllRecordsets() ([]string, error) {
	if v := r.cache.Get("designateListAllRecordsets"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	zones, err := r.ListAllZones()
	for _, zone := range zones {

		pager := recordsets.ListByZone(r.client, zone, recordsets.ListOpts{})
		// Define an anonymous function to be executed on each page's iteration
		err := pager.EachPage(func(page pagination.Page) (bool, error) {

			recordsetsList, err := recordsets.ExtractRecordSets(page)
			if err != nil {
				logrus.Fatalf("Error 0 paging through objects : %s", err)
			}

			for _, n := range recordsetsList {
				k = append(k, zone+"/"+n.ID)
			}

			if err != nil {
				logrus.Fatalf("Error 1 paging through objects : %s", err)
			}

			return true, nil
		})

		if err != nil {
			logrus.Infof("Error 2 paging through objects : %s", err)
		}
	}
	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}
	if len(k) == 0 {
		return k, fmt.Errorf("no recordsets found")
	}

	r.cache.Put("designateListAllRecordsets", k)
	return k, nil
}

func (r *designateRepository) ListAllZones() ([]string, error) {
	if v := r.cache.Get("designateListAllZones"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := zones.List(r.client, zones.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allZones, err := zones.ExtractZones(allPages)
	if err != nil {
		panic(err)
	}

	for _, zone := range allZones {
		k = append(k, zone.ID)
	}

	//pager := zones.List(r.client, zone.ListOpts{})

	// Define an anonymous function to be executed on each page's iteration

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no zones found")
	}

	r.cache.Put("designateListAllZones", k)
	return k, nil
}
