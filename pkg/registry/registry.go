package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	parser "github.com/novln/docker-parser"
)

const (
	dockerhubHostname = "docker.io"
	dockerhubAPIURL   = "registry-1.docker.io"
	gcrHostname       = "gcr.io"
	quayioHostname    = "quay.io"

	paginationHeader = "link"
)

type Registry struct {
	Client *http.Client
	// set baseURLs here for simpler testing
	dockerhuBaseURL string
	gcrBaseURL      string
	quayioBaseURL   string
}

type tagsResponse struct {
	Tags []string
}

func New() *Registry {
	return &Registry{
		Client:          &http.Client{},
		dockerhuBaseURL: fmt.Sprintf("https://%s", dockerhubAPIURL),
		gcrBaseURL:      fmt.Sprintf("https://%s", gcrHostname),
		quayioBaseURL:   fmt.Sprintf("https://%s", quayioHostname),
	}
}

// Tags return all tags for an image
func (r Registry) Tags(image string) ([]string, error) {
	parsed, err := parser.Parse(image)
	if err != nil {
		return nil, fmt.Errorf("could not parse image %q: %v", image, err)
	}
	var client Client
	httpClient := basicHTTPClient{client: r.Client}
	switch regsitry := parsed.Registry(); regsitry {
	case dockerhubHostname:
		httpClient.baseURL = r.dockerhuBaseURL
		client = &dockerhubClient{httpClient}
	case gcrHostname:
		httpClient.baseURL = r.gcrBaseURL
		client = &gcrClient{httpClient, envOrgcloudTokenProvider}
	case quayioHostname:
		httpClient.baseURL = r.quayioBaseURL
		client = &quayioClient{httpClient}
	default:
		return nil, fmt.Errorf("unsupported registry %q", regsitry)
	}

	return getTags(parsed.ShortName(), client)
}

func getTags(image string, client Client) ([]string, error) {
	tags := []string{}
	var paginationParam string
	// use a loop incase results are paginated
	for {
		req, err := http.NewRequest("GET", client.TagsURL(image, paginationParam), nil)
		if err != nil {
			return nil, fmt.Errorf("could not get request for %q: %v", image, err)
		}
		header, err := client.AuthHeader(image)
		if err != nil {
			return nil, fmt.Errorf("could not get auth header for %q: %v", image, err)
		}
		req.Header.Add("Authorization", header)

		resp, err := client.HTTPClient().Do(req)
		if err != nil {
			return nil, fmt.Errorf("could not get tags for image %q: %v", image, err)
		}
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("got a bad return code %d for image %q", resp.StatusCode, image)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("could not read tags response: %v", err)
		}
		defer resp.Body.Close()

		var unmarshaledTagsResp tagsResponse
		if err := json.Unmarshal(body, &unmarshaledTagsResp); err != nil {
			return nil, fmt.Errorf("could not unmarshal tags response: %v", err)
		}

		tags = append(tags, unmarshaledTagsResp.Tags...)
		// pagination header will be empty when on the last page
		paginationHeader := resp.Header.Get(paginationHeader)
		if paginationHeader == "" {
			break
		}
		paginationParam = regexp.MustCompile("[?>;]").Split(paginationHeader, -1)[1]
	}
	return tags, nil
}
