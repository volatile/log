/*
Package log is a handler for the Core.
It prints each request/response information (time, method, path, status, duration).

Example

"Hello, World!" example:

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
*/
package log
