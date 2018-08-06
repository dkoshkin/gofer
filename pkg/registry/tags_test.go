package registry

import "testing"

func TestLatest(t *testing.T) {
	tests := []struct {
		tags     *Tags
		expected string
	}{
		{tags: &Tags{List: []string{"v1.10.6", "v1.10.5", "v1.10.4"}}, expected: "v1.10.6"},
		{tags: &Tags{List: []string{"v1.10.6", "v1.10.5", "v1.10.4", "alpha", "stable"}}, expected: "v1.10.6"},
		{tags: &Tags{List: []string{"v1.10.6", "v1.10.5", "v1.10.4", "alpha", "stable", "v1.9.10"}}, expected: "v1.10.6"},
		{tags: &Tags{List: []string{"2.6", "2.7", "3.1", "3.2", "3.3", "3.4", "3.5", "3.6", "3.7", "3.8", "edge", "latest"}}, expected: "3.8"},
		{tags: &Tags{List: []string{"v1.10.0", "v1.10.0-alpha.0", "v1.10.0-alpha.1", "v1.10.0-alpha.2", "v1.10.0-alpha.3", "v1.10.0-beta.0", "v1.10.0-beta.1", "v1.10.0-beta.2", "v1.10.0-beta.3", "v1.10.0-beta.4", "v1.10.0-rc.1"}}, expected: "v1.10.0"},
		{tags: &Tags{List: []string{"v1.10.0", "v1.11.0-alpha.0", "v1.11.0-alpha.1", "v1.11.0-alpha.2", "v1.11.0-alpha.3", "v1.11.0-beta.0", "v1.11.0-beta.1", "v1.11.0-beta.2", "v1.11.0-beta.3", "v1.11.0-beta.4", "v1.11.0-rc.1"}}, expected: "v1.11.0-rc.1"},
	}

	for _, test := range tests {
		latest := test.tags.Latest()
		if latest != test.expected {
			t.Errorf("expected latest tag to be %q, instead got %q", test.expected, latest)
		}
	}
}
