package flagstruct_test

import (
	"github.com/Benzinga/flagstruct"
)

// This snippet shows how to use flag and environment variable parsing with
// flagstruct.
func Example_simple() {
	conf := struct {
		Compress bool   `flag:"z" usage:"whether or not to use compression" env:"COMPRESS"`
		OutputFn string `flag:"out" usage:"output ~filename~"`
	}{
		Compress: true,
	}

	// Parse flags based on structure.
	flagstruct.Configure(&conf)
}

// This snippet shows how to use flag and environment parsing without using
// the Configure helper that couples configuration with parsing.
func Example_advanced() {
	conf := struct {
		Compress bool   `flag:"z" usage:"whether or not to use compression" env:"COMPRESS"`
		OutputFn string `flag:"out" usage:"output ~filename~"`
	}{
		Compress: true,
	}

	// Set up flags.
	flagstruct.Struct(&conf)

	// Parse environment.
	flagstruct.ParseEnv()

	// Parse flags.
	flagstruct.Parse()
}
