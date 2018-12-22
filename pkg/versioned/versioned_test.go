package versioned

import (
	"reflect"
	"testing"
)

func TestLatest(t *testing.T) {
	tests := []struct {
		versions *Versions
		expected *Versioned
	}{
		{versions: FromStringSlice([]string{"v1.10.6", "v1.10.5", "v1.10.4"}), expected: FromString("v1.10.6")},
		{versions: FromStringSlice([]string{"v1.10.6", "v1.10.5", "v1.10.4", "alpha", "stable"}), expected: FromString("v1.10.6")},
		{versions: FromStringSlice([]string{"v1.10.6", "v1.10.5", "v1.10.4", "alpha", "stable", "v1.9.10"}), expected: FromString("v1.10.6")},
		{versions: FromStringSlice([]string{"2.6", "2.7", "3.1", "3.2", "3.3", "3.4", "3.5", "3.6", "3.7", "3.8", "edge", "latest"}), expected: FromString("3.8")},
		{versions: FromStringSlice([]string{"v1.10.0", "v1.10.0-alpha.0", "v1.10.0-alpha.1", "v1.10.0-alpha.2", "v1.10.0-alpha.3", "v1.10.0-beta.0", "v1.10.0-beta.1", "v1.10.0-beta.2", "v1.10.0-beta.3", "v1.10.0-beta.4", "v1.10.0-rc.1"}), expected: FromString("v1.10.0")},
		{versions: FromStringSlice([]string{"v1.10.0", "v1.11.0-alpha.0", "v1.11.0-alpha.1", "v1.11.0-alpha.2", "v1.11.0-alpha.3", "v1.11.0-beta.0", "v1.11.0-beta.1", "v1.11.0-beta.2", "v1.11.0-beta.3", "v1.11.0-beta.4", "v1.11.0-rc.1"}), expected: FromString("v1.11.0-rc.1")},
	}

	for _, test := range tests {
		latest := test.versions.Latest()
		if *latest != *test.expected {
			t.Errorf("expected latest tag to be %q, instead got %q", *test.expected, *latest)
		}
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		versions *Versions
		mask     string
		expected *Versions
	}{
		{
			versions: FromStringSlice([]string{"v1.10.6", "v1.10.5", "v1.10.4", "v1.11.1"}),
			mask:     "v1.10.[0-9]+",
			expected: FromStringSlice([]string{"v1.10.4", "v1.10.5", "v1.10.6"}),
		},
		{
			versions: FromStringSlice([]string{"v1.10.6", "v1.10.5", "v1.10.4", "v1.11.1"}),
			expected: FromStringSlice([]string{"v1.10.4", "v1.10.5", "v1.10.6", "v1.11.1"}),
		},
		{
			versions: FromStringSlice([]string{"v1.10.6", "v1.10.5", "v1.10.4", "v1.11.1"}),
			mask:     "v1.11.[0-9]+",
			expected: FromStringSlice([]string{"v1.11.1"}),
		},
		{
			versions: FromStringSlice([]string{"v1.10.6", "v1.10.5", "v1.10.4", "v1.11.1"}),
			mask:     "v1.12.[0-9]+",
			expected: FromStringSlice([]string{}),
		},
	}

	for _, test := range tests {
		filtered := Filter(test.versions, test.mask)
		if len(test.expected.List) > 0 && !reflect.DeepEqual(filtered, test.expected) {
			t.Errorf("expected filtered tags to be %q, instead got %q", test.expected, filtered)
		}
	}
}
