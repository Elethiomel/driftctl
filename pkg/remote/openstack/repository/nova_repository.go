package repository

import (
	"github.com/cloudskiff/driftctl/pkg/remote/cache"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/sirupsen/logrus"
)

type NovaRepository interface {
}

type novaRepository struct {
	client *gophercloud.ServiceClient
	cache  cache.Cache
}

func NewNovaRepository(providerClient *gophercloud.ProviderClient, c cache.Cache) *novaRepository {
	client, err := openstack.NewObjectStorageV1(providerClient, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Fatalf("Could not create Openstack Client for Nova : %s", err)
	}

	return &novaRepository{
		client,
		c,
	}
}
