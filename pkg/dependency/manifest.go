package dependency

// Manifest contains a list of dependencies
type Manifest struct {
	APIVersion   string `json:"apiVersion"`
	Dependencies []Spec `json:"dependencies"`
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
