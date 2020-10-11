# dmrmonitor [![GoDoc](https://godoc.org/github.com/wilhelmstoll/dmrmonitor?status.svg)](https://godoc.org/github.com/wilhelmstoll/dmrmonitor) [![Build Status](https://travis-ci.org/wilhelmstoll/dmrmonitor.svg?branch=main)](https://travis-ci.org/wilhelmstoll/dmrmonitor)

Pure go implementation of access and process the dmr website.

## Installation

```
go get -u github.com/wilhelmstoll/dmrmonitor
```

## Example of usage

```go
package main

import (
	"fmt"

	"github.com/wilhelmstoll/dmrmonitor"
)

func main() {
	resp := dmrmonitor.Get()

	fmt.Println(resp.ActiveDmrEntries)
	fmt.Println(resp.FinishedDmrEntries)
}
```

## Reference

https://godoc.org/github.com/wilhelmstoll/dmrmonitor
