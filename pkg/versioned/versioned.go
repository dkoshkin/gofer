package versioned

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/mcuadros/go-version"
)

type Versioned string

func FromString(in string) *Versioned {
	versioned := Versioned(in)
	return &versioned
}

func (v Versioned) String() string {
	return string(v)
}

type Versions struct {
	List []Versioned
}

func FromStringSlice(in []string) *Versions {
	versions := &Versions{
		List: make([]Versioned, 0),
	}

	for _, element := range in {
		versions.List = append(versions.List, Versioned(element))
	}
	return versions
}

func (t Versions) Len() int {
	return len(t.List)
}

func (t *Versions) Less(i, j int) bool {
	return version.CompareSimple(version.Normalize(string(t.List[i])), version.Normalize(string(t.List[j]))) == -1
}

func (t *Versions) Swap(i, j int) {
	t.List[i], t.List[j] = t.List[j], t.List[i]
}

func (t Versions) Last() *Versioned {
	if len(t.List) == 0 {
		return nil
	}
	return &t.List[len(t.List)-1]
}

func (t Versions) Latest() *Versioned {
	sort.Sort(&t)
	return t.Last()
}

func Filter(in *Versions, mask string) *Versions {
	sort.Sort(in)
	// if no mask just return all tags
	if mask == "" {
		return in
	}
	// with mask filter out
	rgxMask := regexp.MustCompile(fmt.Sprintf("^%s$", mask))
	filteredVersions := &Versions{}
	for _, tag := range in.List {
		if rgxMask.MatchString(string(tag)) {
			filteredVersions.List = append(filteredVersions.List, tag)
		}
	}
	return filteredVersions
}
