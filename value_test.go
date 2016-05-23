package flagstruct

import (
	"flag"
	"testing"
	"time"
)

func TestBoolValue(t *testing.T) {
	b := false
	v := (*BoolValue)(&b)

	if v.Get() != false {
		t.Error("expected Get to return false")
	}

	err := v.Set("true")
	if err != nil {
		t.Errorf("Set returned %v, expected %v", err, nil)
	}

	if v.Get() != true {
		t.Error("expected Get to return true (after Set)")
	}

	if v.IsBoolFlag() != true {
		t.Error("expected IsBoolFlag to return true")
	}
}

func TestIntValue(t *testing.T) {
	i := 0
	v := (*IntValue)(&i)

	result := v.Get()
	if result != 0 {
		t.Errorf("Get returned %v, expected %v", result, 0)
	}

	err := v.Set("-10")
	if err != nil {
		t.Errorf("Set returned %v, expected %v", err, nil)
	}

	result = v.Get()
	if result != -10 {
		t.Errorf("Get returned %v, expected %v (after Set)", result, -10)
	}

	str := v.String()
	if str != "-10" {
		t.Errorf("String returned %v, expected %v", str, "-10")
	}
}

func TestInt64Value(t *testing.T) {
	i := int64(0)
	v := (*Int64Value)(&i)

	result := v.Get()
	if result != int64(0) {
		t.Errorf("Get returned %v, expected %v", result, 0)
	}

	err := v.Set("-10")
	if err != nil {
		t.Errorf("Set returned %v, expected %v", err, nil)
	}

	result = v.Get()
	if result != int64(-10) {
		t.Errorf("Get returned %v, expected %v (after Set)", result, -10)
	}

	str := v.String()
	if str != "-10" {
		t.Errorf("String returned %v, expected %v", str, "-10")
	}
}

func TestUintValue(t *testing.T) {
	i := uint(0)
	v := (*UintValue)(&i)

	result := v.Get()
	if result != uint(0) {
		t.Errorf("Get returned %v, expected %v", result, 0)
	}

	err := v.Set("10")
	if err != nil {
		t.Errorf("Set returned %v, expected %v", err, nil)
	}

	result = v.Get()
	if result != uint(10) {
		t.Errorf("Get returned %v, expected %v (after Set)", result, 10)
	}

	str := v.String()
	if str != "10" {
		t.Errorf("String returned %v, expected %v", str, "10")
	}
}

func TestUint64Value(t *testing.T) {
	i := uint64(0)
	v := (*Uint64Value)(&i)

	result := v.Get()
	if result != uint64(0) {
		t.Errorf("Get returned %v, expected %v", result, 0)
	}

	err := v.Set("10")
	if err != nil {
		t.Errorf("Set returned %v, expected %v", err, nil)
	}

	result = v.Get()
	if result != uint64(10) {
		t.Errorf("Get returned %v, expected %v (after Set)", result, 10)
	}

	str := v.String()
	if str != "10" {
		t.Errorf("String returned %v, expected %v", str, "10")
	}
}

func TestStringValue(t *testing.T) {
	s := "empty"
	v := (*StringValue)(&s)

	result := v.Get()
	if result != "empty" {
		t.Errorf("Get returned %v, expected %v", result, "empty")
	}

	err := v.Set("full")
	if err != nil {
		t.Errorf("Set returned %v, expected %v", err, nil)
	}

	result = v.Get()
	if result != "full" {
		t.Errorf("Get returned %v, expected %v (after Set)", result, "full")
	}

	str := v.String()
	if str != "full" {
		t.Errorf("String returned %v, expected %v", str, "full")
	}
}

func TestFloat64Value(t *testing.T) {
	f := float64(0.1)
	v := (*Float64Value)(&f)

	result := v.Get()
	if result != float64(0.1) {
		t.Errorf("Get returned %v, expected %v", result, float64(0.1))
	}

	err := v.Set("0.2")
	if err != nil {
		t.Errorf("Set returned %v, expected %v", err, nil)
	}

	result = v.Get()
	if result != float64(0.2) {
		t.Errorf("Get returned %v, expected %v (after Set)", result, float64(0.2))
	}

	str := v.String()
	if str != "0.2" {
		t.Errorf("String returned %v, expected %v", str, "0.2")
	}
}

func TestDurationValue(t *testing.T) {
	d, _ := time.ParseDuration("1s")
	v := (*DurationValue)(&d)

	result := v.Get().(time.Duration)
	if result.String() != "1s" {
		t.Errorf("Get returned %v, expected %v", result.String(), "1s")
	}

	err := v.Set("2s")
	if err != nil {
		t.Errorf("Set returned %v, expected %v", err, nil)
	}

	result = v.Get().(time.Duration)
	if result.String() != "2s" {
		t.Errorf("Get returned %v, expected %v (after Set)", result, "2s")
	}

	str := v.String()
	if str != "2s" {
		t.Errorf("String returned %v, expected %v", str, "2s")
	}
}

func TestValueFromPointer(t *testing.T) {
	_, err := ValueFromPointer(nil)
	if err == nil {
		t.Error("expected err to not be nil")
	} else if err.Error() != "unhandled flag type %!t(<nil>)" {
		t.Error("unexpected error", err)
	}

	// Test support for upstream flag interface implementations
	b := false
	fv := flag.Value((*BoolValue)(&b))

	_, err = ValueFromPointer(fv)
	if err != nil {
		t.Error("unexpected error", err)
	}
}
