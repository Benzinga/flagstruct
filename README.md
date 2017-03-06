# flagstruct [![Build Status](https://travis-ci.org/Benzinga/flagstruct.svg?branch=master)](https://travis-ci.org/Benzinga/flagstruct) [![Go Report Card](https://goreportcard.com/badge/github.com/Benzinga/flagstruct)](https://goreportcard.com/report/github.com/Benzinga/flagstruct) [![codecov](https://codecov.io/gh/Benzinga/flagstruct/branch/master/graph/badge.svg)](https://codecov.io/gh/Benzinga/flagstruct) [![GoDoc](https://godoc.org/github.com/Benzinga/flagstruct?status.svg)](https://godoc.org/github.com/Benzinga/flagstruct)
`flagstruct` is another library for parsing command line flags into structs.
Although packages named `flagstruct` already exist, this pattern emerged
coincidentally in some of our projects, and I decided to simply merge the
best parts into one library.

`flagstruct` has a few neat advantages:

  - Uses values as defaults; no need to encode these as struct tags.
  - Supports generating usage data with groups and specified ordering.
    To use grouping, simply use an unnamed `struct{}`-typed member as a
    separator.
  - Supports custom `flag.Value` types in structures, along with the built-in
    `flag.Value` types.
  - Implements FlagSets akin to Go's `flag.FlagSet`.
  - Implements environment variables through `env` flag.
  - Boolean special case is handled identically to Go's `flag` package.

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
