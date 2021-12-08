package repository

import (
	"fmt"

	"github.com/cloudskiff/driftctl/pkg/remote/cache"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/keypairs"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/secgroups"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/sirupsen/logrus"
)

type NovaRepository interface {
	ListAllKeypairs() ([]string, error)
	ListAllFlavors() ([]string, error)
	ListAllSecgroups() ([]string, error)
	ListAllInstances() ([]string, error)
}

type novaRepository struct {
	client *gophercloud.ServiceClient
	cache  cache.Cache
}

func NewNovaRepository(providerClient *gophercloud.ProviderClient, c cache.Cache) *novaRepository {
	client, err := openstack.NewComputeV2(providerClient, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Fatalf("Could not create Openstack Client for Nova : %s", err)
	}

	logrus.Infof("KEYPAIR %+v\n", client)
	return &novaRepository{
		client,
		c,
	}
}

func (r *novaRepository) ListAllKeypairs() ([]string, error) {
	if v := r.cache.Get("novaListAllKeypairs"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)
	pager := keypairs.List(r.client)
	// Define an anonymous function to be executed on each page's iteration
	err := pager.EachPage(func(page pagination.Page) (bool, error) {

		keypairsList, err := keypairs.ExtractKeyPairs(page)
		if err != nil {
			logrus.Fatalf("Error 0 paging through objects : %s", err)
		}

		for _, n := range keypairsList {
			k = append(k, n.Name)
			logrus.Infof("keypair %s\n", n.Name)
		}

		if err != nil {
			logrus.Fatalf("Error 1 paging through objects : %s", err)
		}

		return true, nil
	})

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no keypairs found")
	}

	r.cache.Put("novaListAllKeypairs", k)
	return k, nil
}

func (r *novaRepository) ListAllFlavors() ([]string, error) {
	if v := r.cache.Get("novaListAllFlavors"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := flavors.ListDetail(r.client, flavors.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		panic(err)
	}

	for _, flavor := range allFlavors {
		k = append(k, flavor.ID)
		logrus.Infof("flavor %s\n", flavor.ID)
		fmt.Printf("%+v\n", flavor)
	}

	//pager := flavors.List(r.client, flavor.ListOpts{})

	// Define an anonymous function to be executed on each page's iteration

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no flavors found")
	}

	r.cache.Put("novaListAllFlavors", k)
	return k, nil
}

func (r *novaRepository) ListAllInstances() ([]string, error) {
	if v := r.cache.Get("novaListAllInstances"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := servers.List(r.client, servers.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allInstances, err := servers.ExtractServers(allPages)
	if err != nil {
		panic(err)
	}

	for _, instance := range allInstances {
		k = append(k, instance.ID)
		logrus.Infof("instance %s\n", instance.ID)
		fmt.Printf("%+v\n", instance)
	}

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no instances found")
	}

	r.cache.Put("novaListAllInstances", k)
	return k, nil
}

func (r *novaRepository) ListAllSecgroups() ([]string, error) {
	if v := r.cache.Get("novaListAllSecgroups"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := secgroups.List(r.client).AllPages()
	if err != nil {
		panic(err)
	}

	allSecgroups, err := secgroups.ExtractSecurityGroups(allPages)
	if err != nil {
		panic(err)
	}

	for _, secgroup := range allSecgroups {
		k = append(k, secgroup.ID)
		logrus.Infof("secgroup %s\n", secgroup.ID)
		fmt.Printf("%+v\n", secgroup)
	}

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no secgroups found")
	}

	r.cache.Put("novaListAllSecgroups", k)
	return k, nil
}
