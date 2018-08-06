package docker

import (
	"errors"

	"github.com/dkoshkin/gofer/pkg/dependency"
	"github.com/dkoshkin/gofer/pkg/registry"
)

type Client struct {
}

var NoTags = errors.New("tags could not be retrieved")

// New returns a dependency fetcher for Docker images
func New() dependency.Fetcher {
	return Client{}
}

func (c Client) LatestVersion(image, mask string) (string, error) {
	dc := registry.New()
	tags, err := dc.Tags(image, mask)
	if err != nil {
		return "", err
	}
	if tags == nil || len(tags.List) == 0 {
		return "", NoTags
	}

	return tags.Latest(), nil
}
