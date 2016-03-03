package flagstruct

import (
	"flag"
	"fmt"
	"reflect"
	"time"
)

// This is a copy of flag.UnquoteUsage that doesn't rely on flag.Flag.
// Unfortunately, this is the only way around, since we have no access to the
// flag types.
func unquoteUsage(u string, value interface{}) (name string, usage string) {
	// Look for a back-quoted name, but avoid the strings package.
	usage = u
	for i := 0; i < len(usage); i++ {
		// Also allow ~, because we can't use ` in struct tag easily.
		if usage[i] == '~' || usage[i] == '`' {
			for j := i + 1; j < len(usage); j++ {
				if usage[j] == '~' || usage[j] == '`' {
					name = usage[i+1 : j]
					usage = usage[:i] + name + usage[j+1:]
					return name, usage
				}
			}
			break // Only one back quote; use type name.
		}
	}

	// No explicit name, so use type if we can find one.
	name = "value"
	switch value.(type) {
	case bool:
		name = ""
	case time.Duration:
		name = "duration"
	case float64:
		name = "float"
	case int, int64:
		name = "int"
	case string:
		name = "string"
	case uint, uint64:
		name = "uint"
	}
	return
}

// Copied from flag.
func isZeroValue(value string) bool {
	switch value {
	case "false":
		return true
	case "":
		return true
	case "0":
		return true
	}
	return false
}

// PrintDefaults prints to standard error the default values of all
// defined command-line flags in the set. This is copied from flag, adjusted
// to allow ~ quotes in default values.
func (s *FlagSet) PrintDefaults() {
	s.VisitAll(func(f *flag.Flag) {
		buf := fmt.Sprintf("  -%s", f.Name)
		val := f.Value.(flag.Getter).Get()
		name, usage := unquoteUsage(f.Usage, val)
		if len(name) > 0 {
			buf += " " + name
		}
		if len(buf) <= 4 {
			buf += "\t"
		} else {
			buf += "\n    \t"
		}
		buf += usage
		if !isZeroValue(f.DefValue) {
			if _, ok := val.(string); ok {
				buf += fmt.Sprintf(" (default %q)", f.DefValue)
			} else {
				buf += fmt.Sprintf(" (default %v)", f.DefValue)
			}
		}
		fmt.Fprint(s.out(), buf, "\n")
	})
}

// PrintStruct prints flags based on the struct passed to `conf`.
func (s *FlagSet) PrintStruct(conf interface{}) {
	v := reflect.ValueOf(conf).Elem()
	t := reflect.TypeOf(conf).Elem()

	for i, l := 0, t.NumField(); i < l; i++ {
		ft, fv := t.Field(i), v.Field(i)

		// _ can be used to seperate sections.
		if ft.Name == "_" {
			fmt.Fprint(s.out(), "\n")
			continue
		}

		// Skip unexported fields.
		if ft.PkgPath != "" {
			continue
		}

		name, usage := ft.Tag.Get("flag"), ft.Tag.Get("usage")
		if name == "" {
			continue
		}

		typn, usage := unquoteUsage(usage, fv.Interface())
		val := fv.Interface()

		buf := fmt.Sprintf("  -%s", name)
		if len(typn) > 0 {
			buf += " " + typn
		}

		// Usage on same-line for short flags
		if len(buf) <= 4 {
			buf += "\t"
		} else {
			buf += "\n    \t"
		}

		buf += usage

		// Add default value if non-zero
		if val != reflect.Zero(fv.Type()).Interface() {
			if _, ok := val.(string); ok {
				buf += fmt.Sprintf(" (default %q)", val)
			} else {
				buf += fmt.Sprintf(" (default %v)", val)
			}
		}
		fmt.Fprint(s.out(), buf, "\n")
	}
}
