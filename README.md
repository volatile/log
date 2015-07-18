<p align="center"><img src="http://volatile.whitedevops.com/images/repositories/log/logo.png" alt="Volatile Log" title="Volatile Log"><br><br></p>

Volatile Log is a handler for the [Core](https://github.com/volatile/core).  
It prints each request/response information (time, method, path, status, duration) like this:

![Example](http://volatile.whitedevops.com/images/repositories/log/example.png)

## Installation

```Shell
$ go get -u github.com/volatile/log
```

## Usage

```Go
package main

import (
	"fmt"

	"github.com/volatile/core"
	"github.com/volatile/log"
)

func main() {
	if !core.Production {
		log.Use()
	}

	core.Use(func(c *core.Context) {
		fmt.Fprint(c.ResponseWriter, "Hello, World!")
	})

	core.Run()
}
```

[![GoDoc](https://godoc.org/github.com/volatile/log?status.svg)](https://godoc.org/github.com/volatile/log)
