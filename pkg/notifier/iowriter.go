package notifier

import (
	"fmt"
	"github.com/dkoshkin/gofer/pkg/dependency"
	"io"
)

type IOWriter struct {
	writer io.Writer
}

func NewIOWriter(writer io.Writer) Notifier {
	return &IOWriter{writer: writer}
}

func (n *IOWriter) Send(newDependencies []dependency.Spec, updatedDependencies []dependency.Spec) error {
	if len(newDependencies) > 0 {
		fmt.Fprintln(n.writer, "New Dependencies:")
		for _, dep := range newDependencies {
			fmt.Fprintln(n.writer, dep)
		}
	}

	if len(updatedDependencies) > 0 {
		fmt.Fprintln(n.writer, "Updated Dependencies:")
		for _, dep := range updatedDependencies {
			fmt.Fprintln(n.writer, dep)
		}
	}

	return nil
}
