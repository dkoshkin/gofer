package registry

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	dockerhubAuthURL = "https://auth.docker.io/token?service=registry.docker.io&scope=repository:%s:pull"

	tagsURLTemplate = "/v2/%s/tags/list?%s"

	gcrTokenEnv = "GOOGLE_ACCESS_TOKEN"
)

type Client interface {
	AuthHeader(image string) (string, error)
	TagsURL(image, pagination string) string
	HTTPClient() *http.Client
}

type basicHTTPClient struct {
	client  *http.Client
	baseURL string
}

func (c *basicHTTPClient) HTTPClient() *http.Client {
	return c.client
}

func (c *basicHTTPClient) TagsURL(image, pagination string) string {
	return fmt.Sprintf(c.baseURL+tagsURLTemplate, image, pagination)
}

type dockerhubClient struct {
	basicHTTPClient
}

type githubAuthResponse struct {
	Token string
}

func (c *dockerhubClient) AuthHeader(image string) (string, error) {
	// get authorization token
	authURL := fmt.Sprintf(dockerhubAuthURL, image)
	resp, err := c.client.Get(authURL)
	if err != nil {
		return "", fmt.Errorf("could not get authorization token: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var unmarshaledAuthResp githubAuthResponse
	if err := json.Unmarshal(body, &unmarshaledAuthResp); err != nil {
		return "", fmt.Errorf("could not unmarshal authorization token: %v", err)
	}
	return fmt.Sprintf("Bearer %s", unmarshaledAuthResp.Token), nil
}

type gcrClient struct {
	basicHTTPClient
	tokenProvider func() (string, error)
}

func (c gcrClient) AuthHeader(_ string) (string, error) {
	token, err := c.tokenProvider()
	if err != nil {
		return "", err
	}
	encoded := b64.RawStdEncoding.EncodeToString([]byte("_token:" + token))
	return fmt.Sprintf("Basic %s", encoded), nil
}

func envOrgcloudTokenProvider() (string, error) {
	token := os.Getenv(gcrTokenEnv)
	if token == "" {
		tokenBytes, err := exec.Command("gcloud", "auth", "print-access-token").Output()
		if err != nil {
			return "", fmt.Errorf("error running %q: %v and %q is not set", "gcloud auth print-access-token", err, gcrTokenEnv)
		}
		token = string(tokenBytes)
	}
	return strings.TrimSpace(token), nil
}

type quayioClient struct {
	basicHTTPClient
}

func (c quayioClient) AuthHeader(_ string) (string, error) {
	return "", nil
}
