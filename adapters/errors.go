package adapters

import "fmt"

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() error {
	return fmt.Errorf("object with name '%s' not found", e.Name)
}
