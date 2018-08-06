package manager

import "github.com/dkoshkin/gofer/pkg/dependency"

// ReadWriter persists dependencies
type ReadWriter interface {
	Init(apiVersion string) (*dependency.Manifest, error)
	Read() (*dependency.Manifest, error)
	Write(manifest dependency.Manifest) error
}
