package github

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dkoshkin/gofer/pkg/dependency"
	gh "github.com/google/go-github/v20/github"
	"golang.org/x/oauth2"
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

func (c Client) LatestVersion(url, mask string) (string, error) {
	project, err := projectFromURL(url)
	if err != nil {
		return "", err
	}
	ownerRepoPair := strings.Split(project, "/")
	if len(ownerRepoPair) != 2 {
		return "", fmt.Errorf("%q not a valid Github owner:repo format", project)
	}
	releases, _, err := c.github.Repositories.ListReleases(context.Background(), ownerRepoPair[0], ownerRepoPair[1], nil)
	if err != nil {
		return "", fmt.Errorf("could not get versions :%v", err)
	}
	if len(releases) == 0 {
		return "", dependency.ErrEmptyVerionsList
	}
	tags := Tags{List: releases}
	filtertered := filterTags(tags, mask)

	return filtertered.Latest(), nil
}

func projectFromURL(url string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("invalid Github URL %q", url)
	}
	trimmed := strings.TrimLeft(url, "https://")
	return strings.TrimLeft(trimmed, "github.com/"), nil
}
