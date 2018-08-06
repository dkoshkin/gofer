package github

import "github.com/dkoshkin/gofer/pkg/dependency"

type Client struct {
}

// New returns a dependency fetcher for github
func New() dependency.Fetcher {
	return Client{}
}

func (c Client) LatestVersion(image, mask string) (string, error) {
	return "", nil
}
