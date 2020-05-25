package dependency

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	UnknownType = "unknown"
	ManualType  = "manual"
	DockerType  = "docker"
	GithubType  = "github"

	githubTypePrefix      = "https://github.com/"
	githubTypePrefixShort = "github.com/"
)

// Spec describes a resource
// Type: github, docker, manual
// Source will be specific to a 'Type'
type Spec struct {
	Name          string `yaml:"name" json:"name"`
	Type          string `yaml:"type" json:"type"`
	Version       string `yaml:"version" json:"version"`
	LatestVersion string `yaml:"latestVersion,omitempty" json:"latestVersion"`
	Mask          string `yaml:"mask,omitempty" json:"mask"`
	Notes         string `yaml:"notes,omitempty" json:"notes"`
}

func (s Spec) Hash() (string, error) {
	h := sha1.New()
	_, err := fmt.Fprintf(h, "%s%s%s", s.Name, s.Type, s.Mask)
	if err != nil {
		return "", fmt.Errorf("could not get hash: %v", err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func (s Spec) GetType() string {
	if s.Type != "" {
		return s.Type
	}
	return DetermineType(s.Name)
}

// DetermineType will try to determine the type for the spec
// 'unknown' will be returned if it cannot be determined
func DetermineType(source string) string {
	// just return on custom type
	if source == ManualType {
		return ManualType
	}
	specType := UnknownType
	if strings.HasPrefix(source, githubTypePrefix) || strings.HasPrefix(source, githubTypePrefixShort) {
		specType = GithubType
	} else {
		specType = DockerType
	}
	return specType
}
