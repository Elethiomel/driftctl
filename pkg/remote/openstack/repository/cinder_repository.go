package repository

import (
	"fmt"

	"github.com/cloudskiff/driftctl/pkg/remote/cache"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
	"github.com/sirupsen/logrus"
)

type CinderRepository interface {
	ListAllVolumes() ([]string, error)
}

type cinderRepository struct {
	client *gophercloud.ServiceClient
	cache  cache.Cache
}

func NewCinderRepository(providerClient *gophercloud.ProviderClient, c cache.Cache) *cinderRepository {
	client, err := openstack.NewBlockStorageV2(providerClient, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Fatalf("Could not create Openstack Client for Cinder : %s", err)
	}

	return &cinderRepository{
		client,
		c,
	}
}

func (r *cinderRepository) ListAllVolumes() ([]string, error) {
	if v := r.cache.Get("cinderListAllVolumes"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := volumes.List(r.client, volumes.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allVolumes, err := volumes.ExtractVolumes(allPages)
	if err != nil {
		panic(err)
	}

	for _, volume := range allVolumes {
		k = append(k, volume.ID)
	}

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no volumes found")
	}

	r.cache.Put("cinderListAllVolumes", k)
	return k, nil
}
