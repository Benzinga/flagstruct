// This is copied from Golang's flag library, because it is private :(

package flagstruct

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// BoolValue represents a boolean value.
type BoolValue bool

// Set implements the Value interface.
func (b *BoolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = BoolValue(v)
	return err
}

// Get implements the Value interface.
func (b *BoolValue) Get() interface{} { return bool(*b) }

// String implements the Value interface.
func (b *BoolValue) String() string { return fmt.Sprintf("%v", *b) }

// IsBoolFlag signals boolean flag behavior to Go's flag library.
func (b *BoolValue) IsBoolFlag() bool { return true }

// IntValue represents an integer value.
type IntValue int

// Set implements the Value interface.
func (i *IntValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = IntValue(v)
	return err
}

// Get implements the Value interface.
func (i *IntValue) Get() interface{} { return int(*i) }

// String implements the Value interface.
func (i *IntValue) String() string { return fmt.Sprintf("%v", *i) }

// Int64Value represents a 64-bit integer value.
type Int64Value int64

// Set implements the Value interface.
func (i *Int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = Int64Value(v)
	return err
}

// Get implements the Value interface.
func (i *Int64Value) Get() interface{} { return int64(*i) }

// String implements the Value interface.
func (i *Int64Value) String() string { return fmt.Sprintf("%v", *i) }

// UintValue represents an unsigned integer value.
type UintValue uint

// Set implements the Value interface.
func (i *UintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = UintValue(v)
	return err
}

// Get implements the Value interface.
func (i *UintValue) Get() interface{} { return uint(*i) }

// String implements the Value interface.
func (i *UintValue) String() string { return fmt.Sprintf("%v", *i) }

// Uint64Value represents an unsigned 64-bit integer value.
type Uint64Value uint64

// Set implements the Value interface.
func (i *Uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = Uint64Value(v)
	return err
}

// Get implements the Value interface.
func (i *Uint64Value) Get() interface{} { return uint64(*i) }

// String implements the Value interface.
func (i *Uint64Value) String() string { return fmt.Sprintf("%v", *i) }

// StringValue represents a string value.
type StringValue string

// Set implements the Value interface.
func (s *StringValue) Set(val string) error {
	*s = StringValue(val)
	return nil
}

// Get implements the Value interface.
func (s *StringValue) Get() interface{} { return string(*s) }

// String implements the Value interface.
func (s *StringValue) String() string { return fmt.Sprintf("%s", *s) }

// Float64Value represents a 64-bit floating point value.
type Float64Value float64

// Set implements the Value interface.
func (f *Float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*f = Float64Value(v)
	return err
}

// Get implements the Value interface.
func (f *Float64Value) Get() interface{} { return float64(*f) }

// String implements the Value interface.
func (f *Float64Value) String() string { return fmt.Sprintf("%v", *f) }

// DurationValue represents a duration of time.
type DurationValue time.Duration

// Set implements the Value interface.
func (d *DurationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	*d = DurationValue(v)
	return err
}

// Get implements the Value interface.
func (d *DurationValue) Get() interface{} { return time.Duration(*d) }

// String implements the Value interface.
func (d *DurationValue) String() string { return (*time.Duration)(d).String() }

// Value is an interface used for flag values.
type Value interface {
	String() string
	Set(string) error
	Get() interface{}
}

// ValueFromPointer uses reflection to determine what value type to use.
func ValueFromPointer(ptr interface{}) (Value, error) {
	switch f := ptr.(type) {
	case *bool:
		return (*BoolValue)(f), nil
	case *float64:
		return (*Float64Value)(f), nil
	case *int:
		return (*IntValue)(f), nil
	case *int64:
		return (*Int64Value)(f), nil
	case *string:
		return (*StringValue)(f), nil
	case *uint:
		return (*UintValue)(f), nil
	case *uint64:
		return (*Uint64Value)(f), nil
	case *time.Duration:
		return (*DurationValue)(f), nil
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
