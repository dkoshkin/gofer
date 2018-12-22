package dependency

import (
	"errors"
	"github.com/dkoshkin/gofer/pkg/versioned"
)

var ErrEmptyVerionsList = errors.New("no versions were retrieved")

// Fetcher retrieves information for a resource
type Fetcher interface {
	AllVersions(source, mask string) (versions *versioned.Versions, err error)
	LatestVersion(source, mask string) (version *versioned.Versioned, err error)
}
