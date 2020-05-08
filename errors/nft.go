package errors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidL4ProtoType = errors.New("invalid protocol types parameter")
	ErrTableWithNameNotFound = func(name string) error {return fmt.Errorf("nf table with name: %v, not found", name)}
)
