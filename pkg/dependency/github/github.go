package github

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	gh "github.com/google/go-github/v31/github"
	"golang.org/x/oauth2"

	"github.com/dkoshkin/gofer/pkg/dependency"
	"github.com/dkoshkin/gofer/pkg/versioned"
)

const (
	githubTokeneEnv = "GITHUB_ACCESS_TOKEN"
)

type Client struct {
	github *gh.Client
	token  string
}

// New returns a dependency fetcher for github
func New() dependency.Fetcher {
	var tc *http.Client
	// use client token if provided
	token := os.Getenv(githubTokeneEnv)
	if token != "" {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc = oauth2.NewClient(ctx, ts)

	}
	client := gh.NewClient(tc)
	return Client{github: client, token: token}
}

func (c Client) AllVersions(url, mask string) (*versioned.Versions, error) {
	project, err := projectFromURL(url)
	if err != nil {
		return nil, err
	}
	ownerRepoPair := strings.Split(project, "/")
	if len(ownerRepoPair) != 2 {
		return nil, fmt.Errorf("%q not a valid Github owner:repo format", project)
	}
	releases, _, err := c.github.Repositories.ListReleases(context.Background(), ownerRepoPair[0], ownerRepoPair[1], nil)
	if err != nil {
		return nil, fmt.Errorf("could not get versions :%v", err)
	}
	if len(releases) == 0 {
		return nil, dependency.ErrEmptyVerionsList
	}
	tags := toStringSlice(releases)

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

func projectFromURL(url string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("invalid Github URL %q", url)
	}
	trimmed := strings.TrimLeft(url, "https://")
	return strings.TrimLeft(trimmed, "github.com/"), nil
}

func toStringSlice(releases []*gh.RepositoryRelease) []string {
	out := make([]string, 0)
	for n := range releases {
		out = append(out, *releases[n].TagName)
	}

	return out
}
