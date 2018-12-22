package docker

import (
	"fmt"
	"github.com/dkoshkin/gofer/pkg/dependency"
	"github.com/dkoshkin/gofer/pkg/registry"
	"github.com/dkoshkin/gofer/pkg/versioned"
)

type Client struct {
}

// New returns a dependency fetcher for Docker images
func New() dependency.Fetcher {
	return Client{}
}

func (c Client) AllVersions(image, mask string) (*versioned.Versions, error) {
	dc := registry.New()
	tags, err := dc.Tags(image)
	if err != nil {
		return nil, err
	}
	if len(tags) == 0 {
		return nil, dependency.ErrEmptyVerionsList
	}

	versions := versioned.FromStringSlice(tags)
	filtered := versioned.Filter(versions, mask)

	return filtered, nil
}

func (c Client) LatestVersion(url, mask string) (*versioned.Versioned, error) {
	versions, err := c.AllVersions(url, mask)
	if err != nil {
		return nil, fmt.Errorf("could not list all tags: %v", err)
	}

	return versions.Latest(), nil
}
