<p align="center"><img src="https://cloud.githubusercontent.com/assets/9503891/8640486/b71678e6-28fa-11e5-8596-5fd6e63896d1.png" alt="Volatile Compress" title="Volatile Compress"><br><br></p>

Volatile Log is a handler for the [Core](https://github.com/volatile/core).  
It prints each request/response information (time, method, path, status, duration).

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
		core.Use(log.Handler)
	}

	core.Use(func(c *core.Context) {
		fmt.Fprint(c.ResponseWriter, "Hello, World!")
	})

	core.Run()
}
```

[![GoDoc](https://godoc.org/github.com/volatile/log?status.svg)](https://godoc.org/github.com/volatile/log)
