package registry

import (
	"sort"

	version "github.com/mcuadros/go-version"
)

type Tags struct {
	List []string
}

func (t Tags) Len() int {
	return len(t.List)
}

func (t *Tags) Less(i, j int) bool {
	return version.CompareSimple(version.Normalize(t.List[i]), version.Normalize(t.List[j])) == -1
}

func (t *Tags) Swap(i, j int) {
	t.List[i], t.List[j] = t.List[j], t.List[i]
}

func (t Tags) Last() string {
	if len(t.List) == 0 {
		return ""
	}
	return t.List[len(t.List)-1]
}

func (t Tags) Latest() string {
	sort.Sort(&t)
	return t.Last()
}
