package dependency

// Fetcher retrieves information for a resource
type Fetcher interface {
	LatestVersion(source, mask string) (latestVersion string, err error)
}
