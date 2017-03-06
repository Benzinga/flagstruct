package flagstruct_test

import (
	"flag"

	"github.com/Benzinga/flagstruct"
)

// This snippet shows how to use flag and environment variable parsing with
// flagstruct.
func Example() {
	conf := struct {
		Compress bool   `flag:"z" usage:"whether or not to use compression" env:"COMPRESS"`
		OutputFn string `flag:"out" usage:"output ~filename~"`
	}{
		Compress: true,
	}

	// Setup enhanced usage help.
	// Note: this will hide flags not in the struct.
	flag.Usage = flagstruct.MakeStructUsage(&conf)

	// Set up flags based on structure.
	flagstruct.Struct(&conf)

	// Parse environment (optional.)
	// You can do this after flags to make env take precedence above flags.
	flagstruct.ParseEnv()

	// Parse flags.
	flagstruct.Parse()
}

// This snippet shows how to use MakeStructUsage to generate a nicer, ordered
// help output. This will exclude flags not in the struct, so using normal
// usage may be desirable in some cases.
func ExampleMakeStructUsage() {
	conf := struct {
		OptionA bool   `flag:"a" usage:"option a" env:"OP_A"`
		OptionB string `flag:"b" usage:"option b" env:"OP_B"`
		OptionC int64  `flag:"c" usage:"option c" env:"OP_C"`
	}{
		OptionB: "test",
	}

	flag.Usage = flagstruct.MakeStructUsage(&conf)
}
