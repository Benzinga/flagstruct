# flagstruct [![Build Status](https://travis-ci.org/Benzinga/flagstruct.svg?branch=master)](https://travis-ci.org/Benzinga/flagstruct) [![Go Report Card](https://goreportcard.com/badge/github.com/Benzinga/flagstruct)](https://goreportcard.com/report/github.com/Benzinga/flagstruct) [![codecov](https://codecov.io/gh/Benzinga/flagstruct/branch/master/graph/badge.svg)](https://codecov.io/gh/Benzinga/flagstruct) [![GoDoc](https://godoc.org/github.com/Benzinga/flagstruct?status.svg)](https://godoc.org/github.com/Benzinga/flagstruct)
`flagstruct` is (another) library for parsing command line flags into structs.
Although packages named `flagstruct` already exist, this pattern emerged
coincidentally (without checking to see if it existed prior) in some projects,
and I decided to simply merge the best parts into one library.

`flagstruct` has a few neat advantages:

  - Uses values as defaults; no need for strings.
  - Supports easily pretty-printing a structure
  - Supports custom `flag.Value` types in structures
  - A useful amount of interoperability with `flag`, allowing you to mix code
    that uses both.
  - Support for FlagSet.
  - Optionally supports environment variable parsing.

## Usage
A quick example follows:

```go
package main

import (
    "flag"
    "github.com/Benzinga/flagstruct"
)

var conf = struct {
    Compress bool   `flag:"z" usage:"whether or not to use compression" env:"COMPRESS"`
    OutputFn string `flag:"out" usage:"output ~filename~"`
}{
    Compress: true,
}

func main() {
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
```
