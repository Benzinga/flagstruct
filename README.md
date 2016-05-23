# flagstruct [![Build Status](http://54.84.4.72:8000/api/badges/Benzinga/flagstruct/status.svg)](http://54.84.4.72:8000/Benzinga/flagstruct)
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

## Usage
A quick example follows:

```go
package main

import (
    "flag"
    "github.com/Benzinga/flagstruct"
)

var conf = struct {
    Compress bool   `flag:"z" usage:"whether or not to use compression"`
    OutputFn string `flag:"out" usage:"output ~filename~"`
}{
    Compress: true,
}

func main() {
    // Setup enhanced usage help
    flag.Usage = flagstruct.MakeStructUsage(&conf)

    // Parse flags.
    flagstruct.ParseStruct(&conf)
}
```
