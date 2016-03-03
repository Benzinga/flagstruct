package flagstruct

import "fmt"

type unhandledTypeError struct {
	typ interface{}
}

func (e unhandledTypeError) Error() string {
	return fmt.Sprintf("unhandled flag type %t", e.typ)
}
