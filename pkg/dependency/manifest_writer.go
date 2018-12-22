package dependency

import (
	"encoding/json"
	"fmt"
	"io"
	"text/tabwriter"

	"gopkg.in/yaml.v3"
)

type ManifestWriter struct {
	Writer        io.Writer
	FilterOptions FilterOptions
}

type FilterOptions struct {
	Outdated bool
	Types    []string
}

func (mf ManifestWriter) WriteTable(m Manifest) {
	fmt.Fprintf(mf.Writer, "String: %q\n", m.APIVersion)
	tw := tabwriter.NewWriter(mf.Writer, 0, 0, 5, ' ', 0)
	fmt.Fprintln(mf.Writer)
	fmt.Fprintln(tw, "Name\tCurrent String\tLatest String\tType\tMask\tNotes")
	fmt.Fprintln(tw, "------\t------\t------\t------\t------\t------")
	for _, dep := range filteredDependencies(m.Dependencies, mf.FilterOptions) {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\t%s\t\n", dep.Name, dep.Version, dep.LatestVersion, dep.GetType(), dep.Mask, dep.Notes)
	}
	tw.Flush()
}

func (mf ManifestWriter) WriteYAML(m Manifest) error {
	m.Dependencies = filteredDependencies(m.Dependencies, mf.FilterOptions)
	b, err := yaml.Marshal(m)
	if err != nil {
		return fmt.Errorf("could not marshal to YAML: %v", err)
	}
	_, err = fmt.Fprintf(mf.Writer, "%s", string(b))
	if err != nil {
		return fmt.Errorf("could not write YAML: %v", err)
	}
	return nil
}

func (mf ManifestWriter) WriteJSON(m Manifest) error {
	m.Dependencies = filteredDependencies(m.Dependencies, mf.FilterOptions)
	b, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return fmt.Errorf("could not marshal to JSON: %v", err)
	}
	_, err = fmt.Fprintf(mf.Writer, "%s\n", string(b))
	if err != nil {
		return fmt.Errorf("could not write JSON: %v", err)
	}
	return nil
}

func filteredDependencies(deps []Spec, filter FilterOptions) []Spec {
	filteredDependencies := make([]Spec, 0)
	for _, dep := range deps {
		// skip	if requesting specific type(s)
		if len(filter.Types) > 0 && !contains(filter.Types, dep.Type) {
			continue
		}
		// skip if requesting only outdate versions
		if filter.Outdated && (dep.Version == dep.LatestVersion) {
			continue
		}
		filteredDependencies = append(filteredDependencies, dep)
	}
	return filteredDependencies
}
