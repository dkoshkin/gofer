package dependency

import (
	"bytes"
	"reflect"
	"testing"
)

func TestFilteredDependencies(t *testing.T) {
	filteredTests := []struct {
		name     string
		deps     []Spec
		expected []Spec
		filter   FilterOptions
	}{
		{
			name:     "filter empty list",
			deps:     []Spec{},
			expected: []Spec{},
			filter:   FilterOptions{Types: []string{"docker"}},
		},
		{
			name:     "filter empty list",
			deps:     []Spec{},
			expected: []Spec{},
			filter:   FilterOptions{Outdated: true},
		},
		{
			name:     "filter empty list",
			deps:     []Spec{},
			expected: []Spec{},
			filter:   FilterOptions{Types: []string{"docker"}, Outdated: true},
		},
		{
			name: "filter none",
			deps: []Spec{
				{
					Name:          "foo",
					Type:          "docker",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
				{
					Name:          "bar",
					Type:          "github",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
				{
					Name:          "foobar",
					Type:          "deb",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
			},
			expected: []Spec{
				{
					Name:          "foo",
					Type:          "docker",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
				{
					Name:          "bar",
					Type:          "github",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
				{
					Name:          "foobar",
					Type:          "deb",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
			},
		},
		{
			name: "filter one type",
			deps: []Spec{
				{
					Name:          "foo",
					Type:          "docker",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
				{
					Name:          "bar",
					Type:          "github",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
				{
					Name:          "foobar",
					Type:          "deb",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
			},
			expected: []Spec{
				{
					Name:          "foo",
					Type:          "docker",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
			},
			filter: FilterOptions{Types: []string{"docker"}},
		},
		{
			name: "filter multiple types",
			deps: []Spec{
				{
					Name:          "foo",
					Type:          "docker",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
				{
					Name:          "bar",
					Type:          "github",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
			},
			expected: []Spec{
				{
					Name:          "foo",
					Type:          "docker",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
				{
					Name:          "bar",
					Type:          "github",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
			},
			filter: FilterOptions{Types: []string{"docker", "github"}},
		},
		{
			name: "filter multiple types",
			deps: []Spec{
				{
					Name:          "foo",
					Type:          "docker",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
				{
					Name:          "bar",
					Type:          "github",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
				{
					Name:          "foobar",
					Type:          "deb",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
			},
			expected: []Spec{
				{
					Name:          "foo",
					Type:          "docker",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
				{
					Name:          "bar",
					Type:          "github",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
			},
			filter: FilterOptions{Types: []string{"docker", "github"}},
		},
		{
			name: "filter outdated, should return none",
			deps: []Spec{
				{
					Name:          "foo",
					Type:          "docker",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
				{
					Name:          "bar",
					Type:          "github",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
				{
					Name:          "foobar",
					Type:          "deb",
					Version:       "v1.0.0",
					LatestVersion: "v1.0.0",
				},
			},
			expected: []Spec{},
			filter:   FilterOptions{Outdated: true},
		},
		{
			name: "filter outdated",
			deps: []Spec{
				{
					Name:          "foo",
					Type:          "docker",
					Version:       "v1.0.0",
					LatestVersion: "v1.1.0",
				},
				{
					Name:          "bar",
					Type:          "github",
					Version:       "v27.0",
					LatestVersion: "v27.1",
				},
				{
					Name:          "foobar",
					Type:          "deb",
					Version:       "1.1",
					LatestVersion: "1.2",
				},
				{
					Name:          "barfoo",
					Type:          "deb",
					Version:       "1.1",
					LatestVersion: "1.1",
				},
			},
			expected: []Spec{
				{
					Name:          "foo",
					Type:          "docker",
					Version:       "v1.0.0",
					LatestVersion: "v1.1.0",
				},
				{
					Name:          "bar",
					Type:          "github",
					Version:       "v27.0",
					LatestVersion: "v27.1",
				},
				{
					Name:          "foobar",
					Type:          "deb",
					Version:       "1.1",
					LatestVersion: "1.2",
				},
			},
			filter: FilterOptions{Outdated: true},
		},
		{
			name: "filter outdated and types",
			deps: []Spec{
				{
					Name:          "foo",
					Type:          "docker",
					Version:       "v1.0.0",
					LatestVersion: "v1.1.0",
				},
				{
					Name:          "bar",
					Type:          "github",
					Version:       "v27.0",
					LatestVersion: "v27.1",
				},
				{
					Name:          "barfoo",
					Type:          "deb",
					Version:       "1.1",
					LatestVersion: "1.1",
				},
			},
			expected: []Spec{},
			filter:   FilterOptions{Outdated: true, Types: []string{"deb"}},
		},
	}

	for _, test := range filteredTests {
		deps := filteredDependencies(test.deps, test.filter)
		if !reflect.DeepEqual(deps, test.expected) {
			t.Errorf("test %q: expected %v, instead got %v", test.name, test.expected, deps)
		}
	}
}

var manifest = Manifest{
	APIVersion: "v1.0",
	Dependencies: []Spec{
		{
			Name:          "alpine",
			Type:          "docker",
			Version:       "3.6",
			LatestVersion: "3.8",
			Mask:          "",
			Notes:         "",
		},
		{
			Name:          "google/cadvisor",
			Type:          "docker",
			Version:       "v0.28.1",
			LatestVersion: "v0.30.2",
			Mask:          "",
			Notes:         "",
		},
		{
			Name:          "google/cadvisor",
			Type:          "docker",
			Version:       "v0.29.2",
			LatestVersion: "v0.29.2",
			Mask:          "v0.29.[0-9]+",
			Notes:         "",
		},
		{
			Name:          "cadvisor",
			Type:          "docker",
			Version:       "v0.28.1",
			LatestVersion: "",
			Mask:          "v0.29.[0-9]+",
			Notes:         "could not find latest tag",
		},
		{
			Name:          "gcr.io/google-containers/kube-apiserver",
			Type:          "docker",
			Version:       "v1.9.6",
			LatestVersion: "v1.9.9",
			Mask:          "v1.9.[0-9]+",
			Notes:         "",
		},
		{
			Name:          "quay.io/coreos/etcd",
			Type:          "docker",
			Version:       "v3.1.13",
			LatestVersion: "v3.1.18",
			Mask:          "v3.1.[0-9]+",
			Notes:         "",
		},
		{
			Name:          "https://github.com/kubernetes/kubernetes",
			Type:          "github",
			Version:       "v1.9.6",
			LatestVersion: "",
			Mask:          "v1.9.[0-9]+",
			Notes:         `unhandled type "github"`,
		},
	},
}

var tableText = `String: "v1.0"

Name                                         Current String     Latest String     Type       Mask             Notes
------                                       ------             ------            ------     ------           ------
alpine                                       3.6                3.8               docker                                                    
google/cadvisor                              v0.28.1            v0.30.2           docker                                                    
google/cadvisor                              v0.29.2            v0.29.2           docker     v0.29.[0-9]+                                   
cadvisor                                     v0.28.1                              docker     v0.29.[0-9]+     could not find latest tag     
gcr.io/google-containers/kube-apiserver      v1.9.6             v1.9.9            docker     v1.9.[0-9]+                                    
quay.io/coreos/etcd                          v3.1.13            v3.1.18           docker     v3.1.[0-9]+                                    
https://github.com/kubernetes/kubernetes     v1.9.6                               github     v1.9.[0-9]+      unhandled type "github"       
`

func TestWriteTable(t *testing.T) {
	var b bytes.Buffer
	mw := ManifestWriter{Writer: &b}
	mw.WriteTable(manifest)
	if b.String() != tableText {
		t.Log(b.String())
		t.Log(tableText)
		t.Error("text written to table output does not equal expected")
	}
}

var yamlText = `apiversion: v1.0
dependencies:
  - name: alpine
    type: docker
    version: "3.6"
    latestVersion: "3.8"
  - name: google/cadvisor
    type: docker
    version: v0.28.1
    latestVersion: v0.30.2
  - name: google/cadvisor
    type: docker
    version: v0.29.2
    latestVersion: v0.29.2
    mask: v0.29.[0-9]+
  - name: cadvisor
    type: docker
    version: v0.28.1
    mask: v0.29.[0-9]+
    notes: could not find latest tag
  - name: gcr.io/google-containers/kube-apiserver
    type: docker
    version: v1.9.6
    latestVersion: v1.9.9
    mask: v1.9.[0-9]+
  - name: quay.io/coreos/etcd
    type: docker
    version: v3.1.13
    latestVersion: v3.1.18
    mask: v3.1.[0-9]+
  - name: https://github.com/kubernetes/kubernetes
    type: github
    version: v1.9.6
    mask: v1.9.[0-9]+
    notes: unhandled type "github"
`

func TestWriteYaml(t *testing.T) {
	var b bytes.Buffer
	mw := ManifestWriter{Writer: &b}
	mw.WriteYAML(manifest)
	if b.String() != yamlText {
		t.Log(b.String())
		t.Log(yamlText)
		t.Error("text written to json output does not equal expected")
	}
}

var jsonText = `{
    "apiVersion": "v1.0",
    "dependencies": [
        {
            "name": "alpine",
            "type": "docker",
            "version": "3.6",
            "latestVersion": "3.8",
            "mask": "",
            "notes": ""
        },
        {
            "name": "google/cadvisor",
            "type": "docker",
            "version": "v0.28.1",
            "latestVersion": "v0.30.2",
            "mask": "",
            "notes": ""
        },
        {
            "name": "google/cadvisor",
            "type": "docker",
            "version": "v0.29.2",
            "latestVersion": "v0.29.2",
            "mask": "v0.29.[0-9]+",
            "notes": ""
        },
        {
            "name": "cadvisor",
            "type": "docker",
            "version": "v0.28.1",
            "latestVersion": "",
            "mask": "v0.29.[0-9]+",
            "notes": "could not find latest tag"
        },
        {
            "name": "gcr.io/google-containers/kube-apiserver",
            "type": "docker",
            "version": "v1.9.6",
            "latestVersion": "v1.9.9",
            "mask": "v1.9.[0-9]+",
            "notes": ""
        },
        {
            "name": "quay.io/coreos/etcd",
            "type": "docker",
            "version": "v3.1.13",
            "latestVersion": "v3.1.18",
            "mask": "v3.1.[0-9]+",
            "notes": ""
        },
        {
            "name": "https://github.com/kubernetes/kubernetes",
            "type": "github",
            "version": "v1.9.6",
            "latestVersion": "",
            "mask": "v1.9.[0-9]+",
            "notes": "unhandled type \"github\""
        }
    ]
}
`

func TestWriteJSON(t *testing.T) {
	var b bytes.Buffer
	mw := ManifestWriter{Writer: &b}
	mw.WriteJSON(manifest)
	if b.String() != jsonText {
		t.Log(b.String())
		t.Log(jsonText)
		t.Error("text written to json output does not equal expected")
	}
}
