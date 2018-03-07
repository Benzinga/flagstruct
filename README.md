[![GoDoc](https://godoc.org/github.com/Benzinga/flagstruct?status.svg)](https://godoc.org/github.com/Benzinga/flagstruct) [![Build Status](https://travis-ci.org/Benzinga/flagstruct.svg?branch=master)](https://travis-ci.org/Benzinga/flagstruct) [![Go Report Card](https://goreportcard.com/badge/github.com/Benzinga/flagstruct)](https://goreportcard.com/report/github.com/Benzinga/flagstruct) [![codecov](https://codecov.io/gh/Benzinga/flagstruct/branch/master/graph/badge.svg)](https://codecov.io/gh/Benzinga/flagstruct)

# flagstruct :checkered_flag:
`flagstruct` is another library for parsing command line flags into structs.
Although packages named `flagstruct` already exist, this pattern emerged
coincidentally in some of our projects, and we decided to simply merge the
best parts into one library. We noticed the name was already taken too late. "_Any resemblance to actual persons, living or dead, or actual events is purely coincidental._"

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

# Getting Started
`flagstruct` is a library. To make use of it, you need to write software that imports it. An example is included below that you can use to play around with flagstruct.

## Prerequisites
`flagstruct` is built in the Go programming language. If you are new to Go, you will need to [install Go](https://golang.org/dl/).

There are no other dependencies, but you may want to configure your text editor for Go if you have not done so.

## Acquiring
Next, you'll want to `go get` flagstruct, like so:

```sh
go get github.com/Benzinga/flagstruct
```

If your `$GOPATH` is configured, and git is setup to know your credentials, in a few moments the command should complete with no output. The repository will exist under `$GOPATH/src/github.com/Benzinga/flagstruct`. It cannot be moved from this location.

Hint: If you've never used Go before, your `$GOPATH` will be under the `go` folder of your user directory.

## Example
A quick example follows:

```go
package main

import (
    "flag"
    "fmt"

    "github.com/Benzinga/flagstruct"
)

var conf = struct {
    Compress bool   `flag:"z" usage:"whether or not to use compression" env:"COMPRESS"`
    OutputFn string `flag:"out" usage:"output ~filename~"`
}{
    Compress: true,
}

func main() {
    // Parse flags and environment based off of flagstruct.
    flagstruct.Configure(&conf)

    // Print out the flags.
    fmt.Printf("Flags: %+v\n", conf)
}
```

> **To use this example:**
> Save to a file with a `.go` extension, like `test.go`, then use `go run test.go [args]` to play with the argument parsing. Do not put the example file into the flagstruct directory as it will get confused with the flagstruct source code.
