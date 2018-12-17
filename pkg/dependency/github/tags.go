package github

import (
	"fmt"
	"regexp"
	"sort"

	gh "github.com/google/go-github/v20/github"
	version "github.com/mcuadros/go-version"
)

type Tags struct {
	List []*gh.RepositoryRelease
}

func (t Tags) Len() int {
	return len(t.List)
}

func (t *Tags) Less(i, j int) bool {
	return version.CompareSimple(version.Normalize(*t.List[i].TagName), version.Normalize(*t.List[j].TagName)) == -1
}

func (t *Tags) Swap(i, j int) {
	t.List[i], t.List[j] = t.List[j], t.List[i]
}

func (t Tags) Last() string {
	if len(t.List) == 0 {
		return ""
	}
	tag := t.List[len(t.List)-1].TagName
	if tag == nil {
		return ""
	}
	return *tag
}

func (t Tags) Latest() string {
	sort.Sort(&t)
	return t.Last()
}

func filterTags(tags Tags, mask string) Tags {
	// if not mask just return all tags
	if mask == "" {
		return tags
	}
	// with mask filter out tags
	rgxMask := regexp.MustCompile(fmt.Sprintf("^%s$", mask))
	filteredTags := Tags{}
	for _, tag := range tags.List {
		if rgxMask.MatchString(*tag.TagName) {
			filteredTags.List = append(filteredTags.List, tag)
		}
	}
	return filteredTags
}
