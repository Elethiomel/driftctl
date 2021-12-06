package enumerator

import (
	"fmt"
	"strings"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/cloudskiff/driftctl/pkg/iac/config"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/objects"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type SwiftEnumerator struct {
	config config.SupplierConfig
	client *gophercloud.ServiceClient
}

func NewSwiftEnumerator(config config.SupplierConfig) *SwiftEnumerator {
	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		logrus.Fatalf("Could not load Openstack auth options from environment : %s", err)
	}
	logrus.Warnf("ProviderClient %+v", opts)

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		logrus.Fatalf("Could not authenticate with Openstack : %s", err)
	}
	logrus.Warnf("ProvierClient %+v", provider)

	client, err := openstack.NewObjectStorageV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Fatalf("Could not creater Openstack Client for Object Storage : %s", err)
	}

	return &SwiftEnumerator{
		config,
		client,
	}
}

func (s *SwiftEnumerator) Origin() string {
	return s.config.String()
}

func (s *SwiftEnumerator) Enumerate() ([]string, error) {
	bucketPath := strings.Split(s.config.Path, "/")
	if len(bucketPath) < 2 {
		return nil, errors.Errorf("Unable to parse Swift path: %s. Must be BUCKET_NAME/PREFIX", s.config.Path)
	}

	bucket := bucketPath[0]
	// prefix should contains everything that does not have a glob pattern
	// Pattern should be the glob matcher string
	prefix, pattern := GlobObjectStorage(strings.Join(bucketPath[1:], "/"))

	fullPattern := strings.Join([]string{prefix, pattern}, "/")
	fullPattern = strings.Trim(fullPattern, "/")

	files := make([]string, 0)

	// We have the option of filtering objects by their attributes
	opts := &objects.ListOpts{Full: true}

	// Retrieve a pager (i.e. a paginated collection)
	pager := objects.List(s.client, bucket, opts)

	// Define an anonymous function to be executed on each page's iteration
	err := pager.EachPage(func(page pagination.Page) (bool, error) {

		// Get a slice of objects.Object structs
		objectList, err := objects.ExtractInfo(page)
		if err != nil {
			logrus.Fatalf("Error paging through objects : %s", err)
		}
		for _, n := range objectList {
			key := n.Name
			if match, _ := doublestar.Match(fullPattern, key); match {
				files = append(files, strings.Join([]string{bucket, key}, "/"))
			}
		}

		return true, nil
	})

	if err != nil {
		logrus.Fatalf("Error paging through objects : %s", err)
	}

	if len(files) == 0 {
		return files, fmt.Errorf("no Terraform state was found in %s, exiting", s.config.Path)
	}

	return files, nil
}
