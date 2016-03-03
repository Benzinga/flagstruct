package flagstruct

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"reflect"
	"time"
)

// A FlagSet represents a set of defined flags.
type FlagSet struct {
	*flag.FlagSet
	name          string
	errorHandling flag.ErrorHandling
	output        io.Writer
}

// NewFlagSet returns a new, empty flag set with the specified name and error
// handling property.
func NewFlagSet(name string, errorHandling flag.ErrorHandling) *FlagSet {
	return &FlagSet{flag.NewFlagSet(name, errorHandling), name, errorHandling, nil}
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
	v := reflect.ValueOf(conf).Elem()
	t := reflect.TypeOf(conf).Elem()

	for i, l := 0, t.NumField(); i < l; i++ {
		ft, fv := t.Field(i), v.Field(i)

		// Skip unexported fields.
		if ft.PkgPath != "" {
			continue
		}

		name, usage := ft.Tag.Get("flag"), ft.Tag.Get("usage")
		if name == "" {
			continue
		}

		addr := fv.Addr().Interface()

		switch f := addr.(type) {
		case *bool:
			s.BoolVar(f, name, *f, usage)
		case *float64:
			s.Float64Var(f, name, *f, usage)
		case *int:
			s.IntVar(f, name, *f, usage)
		case *int64:
			s.Int64Var(f, name, *f, usage)
		case *string:
			s.StringVar(f, name, *f, usage)
		case *uint:
			s.UintVar(f, name, *f, usage)
		case *uint64:
			s.Uint64Var(f, name, *f, usage)
		case *time.Duration:
			s.DurationVar(f, name, *f, usage)
		case flag.Value:
			s.Var(f, name, usage)
		default:
			err := unhandledTypeError{fv.Interface()}

			if s.errorHandling == flag.ContinueOnError {
				return err
			}

			// Panic even on exit-on-error case; do not swallow error.
			panic(err)
		}
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
