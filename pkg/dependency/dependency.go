package dependency

import "errors"

var ErrEmptyVerionsList = errors.New("no versions were retrieved")

// Fetcher retrieves information for a resource
type Fetcher interface {
	LatestVersion(source, mask string) (latestVersion string, err error)
}
