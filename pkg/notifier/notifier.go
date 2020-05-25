package notifier

import "github.com/dkoshkin/gofer/pkg/dependency"

type Notifier interface {
	Send(newDependencies []dependency.Spec, updatedDependencies []dependency.Spec) error
}
