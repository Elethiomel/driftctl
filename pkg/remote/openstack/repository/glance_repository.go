package repository

import (
	"fmt"

	"github.com/cloudskiff/driftctl/pkg/remote/cache"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	"github.com/sirupsen/logrus"
)

type GlanceRepository interface {
	ListAllImages() ([]string, error)
}

type glanceRepository struct {
	client *gophercloud.ServiceClient
	cache  cache.Cache
}

func NewGlanceRepository(providerClient *gophercloud.ProviderClient, c cache.Cache) *glanceRepository {
	client, err := openstack.NewComputeV2(providerClient, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Fatalf("Could not create Openstack Client for Glance : %s", err)
	}

	return &glanceRepository{
		client,
		c,
	}
}

func (r *glanceRepository) ListAllImages() ([]string, error) {
	if v := r.cache.Get("glanceListAllImages"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := images.ListDetail(r.client, images.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allImages, err := images.ExtractImages(allPages)
	if err != nil {
		panic(err)
	}

	for _, image := range allImages {
		k = append(k, image.ID)
	}

	//pager := images.List(r.client, image.ListOpts{})

	// Define an anonymous function to be executed on each page's iteration

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no images found")
	}

	r.cache.Put("glanceListAllImages", k)
	return k, nil
}
