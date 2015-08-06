/*
Package log is a handler for the Core.
It prints each request/response information (time, duration, status, method, path).

Usage

Example:

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
