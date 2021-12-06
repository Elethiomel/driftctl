package backend

import (
	"io"
	"strings"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/objectstorage/v1/objects"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const BackendKeySwift = "swift"

type SwiftBackend struct {
	key       string
	container string
	reader    io.ReadCloser
	client    *gophercloud.ServiceClient
}

func NewSwiftReader(path string) (*SwiftBackend, error) {
	containerPath := strings.Split(path, "/")
	if len(containerPath) < 2 {
		return nil, errors.Errorf("Unable to parse Swift path: %s. Must be BUCKET_NAME/PATH/TO/OBJECT", path)
	}
	container := containerPath[0]
	key := strings.Join(containerPath[1:], "/")

	logrus.Debugf("container %s, key %s", container, key)
	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		logrus.Fatalf("Could not load Openstack auth options from environment : %s", err)
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		logrus.Fatalf("Could not authenticate with Openstack : %s", err)
	}

	client, err := openstack.NewObjectStorageV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Fatalf("Could not creater Openstack Client for Object Storage : %s", err)
	}

	backend := SwiftBackend{
		key:       key,
		container: container,
		client:    client,
	}

	return &backend, nil
}

func (s *SwiftBackend) Read(p []byte) (n int, err error) {
	if s.reader == nil {
		// Configure options
		opts := objects.DownloadOpts{}

		// Download everything into a DownloadResult struct
		res := objects.Download(s.client, s.container, s.key, opts)
		s.reader = res.Body
	}
	return s.reader.Read(p)
}

func (s *SwiftBackend) Close() error {
	if s.reader != nil {
		return s.reader.Close()
	}
	return errors.New("Unable to close reader as nothing was opened")
}
