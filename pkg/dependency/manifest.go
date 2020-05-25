package dependency

import (
	"gopkg.in/yaml.v3"

	"fmt"
	"github.com/dkoshkin/gofer/pkg/fetcher"
	"github.com/dkoshkin/gofer/pkg/fetcher/docker"
	"github.com/dkoshkin/gofer/pkg/fetcher/github"
	"sort"
	"strings"
)

// Manifest contains a list of dependencies
type Manifest struct {
	APIVersion   string `json:"apiVersion"`
	Dependencies []Spec `json:"dependencies"`
}

func FromBytes(in []byte) (*Manifest, error) {
	manifest := &Manifest{}
	err := yaml.Unmarshal(in, manifest)
	if err != nil {
		return nil, err
	}

	return manifest, nil
}

// Append will add a dependency to the struct if it does not already exist
// Returns false if dependency was not added
func (m *Manifest) Append(dep Spec) bool {
	var found bool
	for _, d := range m.Dependencies {
		if d.Name == dep.Name {
			found = true
			break
		}
	}
	if found {
		return false
	}
	m.Dependencies = append(m.Dependencies, dep)
	// if not found, dep was appended
	return true
}

func (m *Manifest) Latest() (*Manifest, error) {
	updatedManifest := &Manifest{APIVersion: m.APIVersion}
	dc := docker.New()
	gc := github.New()
	for _, dep := range m.Dependencies {
		depType := dep.GetType()
		switch depType {
		case DockerType:
			latest, err := dc.LatestVersion(dep.Name, dep.Mask)
			if err != nil {
				if err == fetcher.ErrEmptyVerionsList {
					dep.Notes = fmt.Sprintf("could not find latest tag")
				} else {
					dep.Notes = fmt.Sprintf("error retrieving latest tag: %v", err)
				}
			}
			dep.LatestVersion = latest.String()
			dep.Notes = ""
		case GithubType:
			latest, err := gc.LatestVersion(dep.Name, dep.Mask)
			if err != nil {
				if err == fetcher.ErrEmptyVerionsList {
					dep.Notes = fmt.Sprintf("could not find latest tag")
				} else {
					dep.Notes = fmt.Sprintf("error retrieving latest tag: %v", err)
				}
			}
			dep.LatestVersion = latest.String()
			dep.Notes = ""
		case ManualType:
		case UnknownType:
			dep.Notes = fmt.Sprintf("could not determine type")
		default:
			dep.Notes = fmt.Sprintf("unhandled type %q", depType)
		}
		dep.Type = depType
		updatedManifest.Dependencies = append(updatedManifest.Dependencies, dep)
	}

	return updatedManifest, nil
}

func (m *Manifest) ToMap() (string, map[string]Spec, error) {
	dependenciesMap := map[string]Spec{}
	for n := range m.Dependencies {
		dep := m.Dependencies[n]
		hash, err := dep.Hash()
		if err != nil {
			return "", nil, err
		}
		dependenciesMap[hash] = dep
	}

	return m.APIVersion, dependenciesMap, nil
}

func (m *Manifest) FromMap(version string, dependenciesMap map[string]Spec) *Manifest {
	for k := range dependenciesMap {
		m.Dependencies = append(m.Dependencies, dependenciesMap[k])
	}
	sort.Sort(ByHash(m.Dependencies))

	return m
}

type ByHash []Spec

func (h ByHash) Len() int      { return len(h) }
func (h ByHash) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ByHash) Less(i, j int) bool {
	first, _ := h[i].Hash()
	second, _ := h[j].Hash()
	return strings.Compare(first, second) <= 0
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
