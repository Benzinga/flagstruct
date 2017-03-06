// Package flagstruct is a simple package that allows you to express flag and
// environment variable based configuration using structs and struct tagging.
package flagstruct

import (
	"flag"
	"os"
)

var exit = os.Exit

// CommandLine is the default set of command-line flags, parsed from os.Args.
var CommandLine = &FlagSet{flag.CommandLine, os.Args[0], flag.ExitOnError, nil, map[string]Value{}}

// MakeUsage creates a usage function you can set to flag.Usage.
// For example, you can do this:
//     flag.Usage = flagstruct.MakeUsage()
func MakeUsage() func() {
	return CommandLine.MakeUsage()
}

// MakeStructUsage creates a usage function from a struct that can be set to
// flag.Usage.
func MakeStructUsage(conf interface{}) func() {
	return CommandLine.MakeStructUsage(conf)
}

// Struct loads parameters based off of a struct object.
func Struct(conf interface{}) error {
	return CommandLine.Struct(conf)
}

// Parse parses the command line parameters from argv.
func Parse() error {
	return CommandLine.Parse(os.Args[1:])
}

// ParseStruct parses configuration flags based on the struct passed to `conf`.
func ParseStruct(conf interface{}) error {
	return CommandLine.ParseStruct(conf, os.Args[1:])
}

// ParseEnv parses the environment flags in the structure.
func ParseEnv() error {
	return CommandLine.ParseEnv()
}

// PrintStruct prints configuration flags based on the struct passed to `conf`.
func PrintStruct(conf interface{}) {
	CommandLine.PrintStruct(conf)
}
