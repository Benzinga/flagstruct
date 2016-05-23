package flagstruct

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"reflect"
)

// A FlagSet represents a set of defined flags.
type FlagSet struct {
	*flag.FlagSet
	name          string
	errorHandling flag.ErrorHandling
	output        io.Writer
	env           map[string]Value
}

// NewFlagSet returns a new, empty flag set with the specified name and error
// handling property.
func NewFlagSet(name string, errorHandling flag.ErrorHandling) *FlagSet {
	return &FlagSet{flag.NewFlagSet(name, errorHandling), name, errorHandling, nil, map[string]Value{}}
}

// MakeStructUsage creates a usage function from a struct.
func (s *FlagSet) MakeStructUsage(conf interface{}) func() {
	// Cache struct usage (otherwise default values change)
	buf, oldout := bytes.Buffer{}, s.output
	s.output = &buf
	s.PrintStruct(conf)
	s.output = oldout

	return func() {
		if s.name == "" {
			fmt.Fprintf(s.out(), "Usage:\n")
		} else {
			fmt.Fprintf(s.out(), "Usage of %s:\n", s.name)
		}
		fmt.Fprint(s.out(), buf.String())
	}
}

// MakeUsage creates a usage function that prints the flags.
func (s *FlagSet) MakeUsage() func() {
	return func() {
		if s.name == "" {
			fmt.Fprintf(s.out(), "Usage:\n")
		} else {
			fmt.Fprintf(s.out(), "Usage of %s:\n", s.name)
		}
		s.PrintDefaults()
	}
}

func (s *FlagSet) out() io.Writer {
	if s.output == nil {
		return os.Stderr
	}
	return s.output
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, os.Stderr is used.
func (s *FlagSet) SetOutput(output io.Writer) {
	s.output = output
	s.FlagSet.SetOutput(output)
}

// Struct loads parameters based off of a struct object.
func (s *FlagSet) Struct(conf interface{}) error {
	var err error

	v := reflect.ValueOf(conf).Elem()
	t := reflect.TypeOf(conf).Elem()

	for i, l := 0, t.NumField(); i < l; i++ {
		var key, name, usage string
		var addr interface{}
		var val Value

		ft, fv := t.Field(i), v.Field(i)

		// Skip unexported fields.
		if ft.PkgPath != "" {
			continue
		}

		// Handle 'env' flag.
		key = ft.Tag.Get("env")
		if key != "" && key != "-" {
			addr = fv.Addr().Interface()
			s.env[key], err = ValueFromPointer(addr)
			if err != nil {
				goto HandleErr
			}
		}

		name, usage = ft.Tag.Get("flag"), ft.Tag.Get("usage")
		if name == "" || name == "-" {
			continue
		}

		addr = fv.Addr().Interface()

		// Get Value from pointer.
		val, err = ValueFromPointer(addr)

	HandleErr:
		if err != nil {
			if s.errorHandling == flag.ContinueOnError {
				return err
			}

			// Panic even on exit-on-error case; do not swallow error.
			panic(err)
		}

		s.Var(val.(Value), name, usage)
	}

	return nil
}

// ParseStruct parses configuration flags based on the struct passed to `conf`.
func (s *FlagSet) ParseStruct(conf interface{}, arguments []string) error {
	if err := s.Struct(conf); err != nil {
		return err
	}

	return s.Parse(arguments)
}

// ParseEnv parses environment variables.
func (s *FlagSet) ParseEnv() error {
	var err error

	for key, val := range s.env {
		v, ok := os.LookupEnv(key)
		if !ok {
			continue
		}
		err = val.Set(v)
		if err != nil {
			break
		}
	}

	if err != nil {
		switch s.errorHandling {
		case flag.ExitOnError:
			exit(2)
		case flag.PanicOnError:
			panic(err)
		}
	}

	return err
}
