package repository

import (
	"fmt"

	"github.com/cloudskiff/driftctl/pkg/remote/cache"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/sirupsen/logrus"
)

type NeutronRepository interface {
	ListAllPorts() ([]string, error)
}

type neutronRepository struct {
	client *gophercloud.ServiceClient
	cache  cache.Cache
}

func NewNeutronRepository(providerClient *gophercloud.ProviderClient, c cache.Cache) *neutronRepository {
	client, err := openstack.NewNetworkV2(providerClient, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Fatalf("Could not create Openstack Client for Neutron : %s", err)
	}

	logrus.Infof("Neutron %+v\n", client)
	return &neutronRepository{
		client,
		c,
	}
}

func (r *neutronRepository) ListAllPorts() ([]string, error) {
	if v := r.cache.Get("neutronListAllPorts"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := ports.List(r.client, ports.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allPorts, err := ports.ExtractPorts(allPages)
	if err != nil {
		panic(err)
	}

	for _, port := range allPorts {
		k = append(k, port.ID)
		logrus.Infof("port %s\n", port.ID)
		fmt.Printf("%+v\n", port)
	}

	//pager := ports.List(r.client, port.ListOpts{})

	// Define an anonymous function to be executed on each page's iteration

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no ports found")
	}

	r.cache.Put("neutronListAllPorts", k)
	return k, nil
}
