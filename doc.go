/*
Package log is a handler for the core (https://godoc.org/github.com/volatile/core).
It prints each request/response information (time, duration, status, method, path).

Usage

Use adds the handler to the default handlers stack:
	log.Use()

Make sure to include the handler above any other handler.
*/
package log
