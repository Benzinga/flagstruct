package flagstruct

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type customVal int

func (i *customVal) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = customVal(v)
	return err
}
func (i *customVal) Get() interface{} { return int(*i) }
func (i *customVal) String() string   { return fmt.Sprintf("%v", *i) }

func TestStruct(t *testing.T) {
	conf := struct {
		TestBool       bool   `flag:"test_bool" usage:"bool value"`
		TestInt        int    `flag:"test_int" usage:"int value"`
		TestInt64      int64  `flag:"test_int64" usage:"int64 value"`
		TestUint       uint   `flag:"test_uint" usage:"uint value"`
		TestUint64     uint64 `flag:"test_uint64" usage:"uint64 value"`
		_              struct{}
		TestString     string        `flag:"test_string" usage:"string value"`
		TestFloat64    float64       `flag:"test_float64" usage:"float64 value"`
		TestDuration   time.Duration `flag:"test_duration" usage:"time.Duration value~"`
		TestCustom     customVal     `flag:"test_custom" usage:"~custom~ value"`
		TestShort      bool          `flag:"x" usage:"short flag"`
		NonFlag        int
		testunexported int
	}{}

	flagset := NewFlagSet("program", flag.ContinueOnError)
	err := flagset.Struct(&conf)
	if err != nil {
		t.Error(err)
	}

	// Check to see each value is as expected.
	m := make(map[string]*flag.Flag)
	flagset.VisitAll(func(f *flag.Flag) {
		m[f.Name] = f

		switch {
		case f.Value.String() == "0":
			return
		case f.Value.String() == "false":
			return
		case f.Name == "test_duration" && f.Value.String() == "0s":
			return
		case f.Name == "test_string" && f.Value.String() == "":
			return
		}

		t.Error("bad value", f.Value.String(), "for", f.Name)
	})

	if len(m) != 10 {
		t.Error("wrong number of flags", len(m))
	}

	flagset.Parse([]string{"-test_bool", "-test_duration=1.0s", "-test_custom=23", "-test_string=a.out"})

	if !conf.TestBool {
		t.Error("expected boolean flag to be true")
	}

	if conf.TestDuration.Seconds() != 1 {
		t.Errorf("expected duration 1s, actual value %q", conf.TestDuration.String())
	}

	if conf.TestCustom != customVal(23) {
		t.Errorf("expected customVal 23, actual value %q", conf.TestCustom.String())
	}

	// Test output
	buf := bytes.Buffer{}
	flagset.SetOutput(&buf)
	flagset.PrintStruct(&conf)

	expectedp := "" +
		"  -test_bool\n    \tbool value (default true)\n" +
		"  -test_int int\n    \tint value\n" +
		"  -test_int64 int\n    \tint64 value\n" +
		"  -test_uint uint\n    \tuint value\n" +
		"  -test_uint64 uint\n    \tuint64 value\n" +
		"\n" +
		"  -test_string string\n    \tstring value (default \"a.out\")\n" +
		"  -test_float64 float\n    \tfloat64 value\n" +
		"  -test_duration duration\n    \ttime.Duration value~ (default 1s)\n" +
		"  -test_custom custom\n    \tcustom value (default 23)\n" +
		"  -x\tshort flag\n"
	if buf.String() != expectedp {
		t.Error("print output differs from expected.")
		t.Logf("expected:\n%q\n", expectedp)
		t.Logf("actual:\n%q\n", buf.String())
	}
}

func TestGlobal(t *testing.T) {
	m := make(map[string]*flag.Flag)
	visitor := func(f *flag.Flag) {
		if !strings.HasPrefix(f.Name, "test_") {
			return
		}

		m[f.Name] = f

		if f.Value.String() != "false" {
			t.Error("bad value", f.Value.String(), "for", f.Name)
		}
	}

	conf := struct {
		Bool bool `flag:"test_bool" usage:"bool value"`
	}{}

	// Test Struct + Parse
	CommandLine = NewFlagSet("program", flag.ContinueOnError)
	Struct(&conf)
	CommandLine.VisitAll(visitor)
	Parse()
	if len(m) != 1 {
		t.Error("wrong number of flags", len(m))
	}

	// Test ParseStruct
	CommandLine = NewFlagSet("program", flag.ContinueOnError)
	ParseStruct(&conf)

	// Test output
	buf := bytes.Buffer{}
	CommandLine.SetOutput(&buf)
	PrintStruct(&conf)

	expectedp := "  -test_bool\n    \tbool value\n"
	if buf.String() != expectedp {
		t.Error("print output differs from expected.")
		t.Logf("expected:\n%q\n", expectedp)
		t.Logf("actual:\n%q\n", buf.String())
	}
}

func TestBadTypes(t *testing.T) {
	conf := struct {
		TestInt16 int16 `flag:"test_int16" usage:"int16 value"`
	}{}

	CommandLine = NewFlagSet("program", flag.ContinueOnError)
	err := Struct(&conf)
	if err == nil || err.Error() != "unhandled flag type %!t(int16=0)" {
		t.Error("unexpected error", err)
	}

	err = ParseStruct(&conf)
	if err == nil || err.Error() != "unhandled flag type %!t(int16=0)" {
		t.Error("unexpected error", err)
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Error("expected panic did not occur")
		}
		if r.(error).Error() != "unhandled flag type %!t(int16=0)" {
			t.Error("wrong error", r.(error).Error())
		}
	}()
	CommandLine = NewFlagSet("program", flag.PanicOnError)
	Struct(&conf)
}

func TestOut(t *testing.T) {
	s := NewFlagSet("program", flag.ContinueOnError)
	if s.out() != os.Stderr {
		t.Error("expected out to default to stderr")
	}
	s.SetOutput(os.Stdout)
	if s.out() != os.Stdout {
		t.Error("expected out to set to stdout")
	}
}

func TestStructUsage(t *testing.T) {
	conf := struct {
		X          bool   `flag:"x"`
		Bool       bool   `flag:"test_bool" usage:"bool value"`
		Str        string `flag:"test_str"`
		Str2       string `flag:"test_str2"`
		_          struct{}
		TestCustom customVal `flag:"test_custom" usage:"~custom~ value"`
	}{X: true, Str: "x"}

	buf := bytes.Buffer{}
	CommandLine = NewFlagSet("program", flag.PanicOnError)
	CommandLine.SetOutput(&buf)
	CommandLine.Struct(&conf)

	buf.Reset()

	MakeUsage()()
	expectedp := "Usage of program:\n  -test_bool\n    \tbool value\n  -test_custom custom\n    \tcustom value\n  -test_str string\n    \t (default \"x\")\n  -test_str2 string\n    \t\n  -x\t (default true)\n"
	if buf.String() != expectedp {
		t.Errorf("usage output differs from expected.\nexpected:\n%q\nactual:\n%q\n", expectedp, buf.String())
	}

	buf.Reset()

	MakeStructUsage(&conf)()
	expectedp = "Usage of program:\n  -x\t (default true)\n  -test_bool\n    \tbool value\n  -test_str string\n    \t (default \"x\")\n  -test_str2 string\n    \t\n\n  -test_custom custom\n    \tcustom value\n"
	if buf.String() != expectedp {
		t.Errorf("usage output differs from expected.\nexpected:\n%q\nactual:\n%q\n", expectedp, buf.String())
	}

	CommandLine.name = ""
	buf.Reset()

	MakeUsage()()
	expectedp = "Usage:\n  -test_bool\n    \tbool value\n  -test_custom custom\n    \tcustom value\n  -test_str string\n    \t (default \"x\")\n  -test_str2 string\n    \t\n  -x\t (default true)\n"
	if buf.String() != expectedp {
		t.Errorf("usage output differs from expected.\nexpected:\n%q\nactual:\n%q\n", expectedp, buf.String())
	}

	buf.Reset()

	MakeStructUsage(&conf)()
	expectedp = "Usage:\n  -x\t (default true)\n  -test_bool\n    \tbool value\n  -test_str string\n    \t (default \"x\")\n  -test_str2 string\n    \t\n\n  -test_custom custom\n    \tcustom value\n"
	if buf.String() != expectedp {
		t.Errorf("usage output differs from expected.\nexpected:\n%q\nactual:\n%q\n", expectedp, buf.String())
	}
}
