// This is copied from Golang's flag library, because it is private :(

package flagstruct

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// boolValue represents a boolean value.
type boolValue bool

// Set implements the Value interface.
func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = boolValue(v)
	return err
}

// Get implements the Value interface.
func (b *boolValue) Get() interface{} { return bool(*b) }

// String implements the Value interface.
func (b *boolValue) String() string { return fmt.Sprintf("%v", *b) }

// IsBoolFlag signals boolean flag behavior to Go's flag library.
func (b *boolValue) IsBoolFlag() bool { return true }

// intValue represents an integer value.
type intValue int

// Set implements the Value interface.
func (i *intValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = intValue(v)
	return err
}

// Get implements the Value interface.
func (i *intValue) Get() interface{} { return int(*i) }

// String implements the Value interface.
func (i *intValue) String() string { return fmt.Sprintf("%v", *i) }

// int64Value represents a 64-bit integer value.
type int64Value int64

// Set implements the Value interface.
func (i *int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = int64Value(v)
	return err
}

// Get implements the Value interface.
func (i *int64Value) Get() interface{} { return int64(*i) }

// String implements the Value interface.
func (i *int64Value) String() string { return fmt.Sprintf("%v", *i) }

// uintValue represents an unsigned integer value.
type uintValue uint

// Set implements the Value interface.
func (i *uintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uintValue(v)
	return err
}

// Get implements the Value interface.
func (i *uintValue) Get() interface{} { return uint(*i) }

// String implements the Value interface.
func (i *uintValue) String() string { return fmt.Sprintf("%v", *i) }

// uint64Value represents an unsigned 64-bit integer value.
type uint64Value uint64

// Set implements the Value interface.
func (i *uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uint64Value(v)
	return err
}

// Get implements the Value interface.
func (i *uint64Value) Get() interface{} { return uint64(*i) }

// String implements the Value interface.
func (i *uint64Value) String() string { return fmt.Sprintf("%v", *i) }

// stringValue represents a string value.
type stringValue string

// Set implements the Value interface.
func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}

// Get implements the Value interface.
func (s *stringValue) Get() interface{} { return string(*s) }

// String implements the Value interface.
func (s *stringValue) String() string { return fmt.Sprintf("%s", *s) }

// float64Value represents a 64-bit floating point value.
type float64Value float64

// Set implements the Value interface.
func (f *float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*f = float64Value(v)
	return err
}

// Get implements the Value interface.
func (f *float64Value) Get() interface{} { return float64(*f) }

// String implements the Value interface.
func (f *float64Value) String() string { return fmt.Sprintf("%v", *f) }

// durationValue represents a duration of time.
type durationValue time.Duration

// Set implements the Value interface.
func (d *durationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	*d = durationValue(v)
	return err
}

// Get implements the Value interface.
func (d *durationValue) Get() interface{} { return time.Duration(*d) }

// String implements the Value interface.
func (d *durationValue) String() string { return (*time.Duration)(d).String() }

// Value is an interface used for flag values.
type Value interface {
	String() string
	Set(string) error
	Get() interface{}
}

// valueFromPointer uses reflection to determine what value type to use.
func valueFromPointer(ptr interface{}) (Value, error) {
	switch f := ptr.(type) {
	case *bool:
		return (*boolValue)(f), nil
	case *float64:
		return (*float64Value)(f), nil
	case *int:
		return (*intValue)(f), nil
	case *int64:
		return (*int64Value)(f), nil
	case *string:
		return (*stringValue)(f), nil
	case *uint:
		return (*uintValue)(f), nil
	case *uint64:
		return (*uint64Value)(f), nil
	case *time.Duration:
		return (*durationValue)(f), nil
	case Value:
		return f, nil
	default:
		if ptr == nil {
			return nil, unhandledTypeError{nil}
		}
		t := reflect.ValueOf(ptr).Elem().Interface()
		return nil, unhandledTypeError{t}
	}
}
