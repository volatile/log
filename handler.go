package log

import (
	"fmt"
	"log"
	"time"

	"github.com/volatile/core"
	"github.com/volatile/core/coreutil"
	"github.com/whitedevops/colors"
)

// Use tells the core to use this handler.
func Use() {
	core.Use(func(c *core.Context) {
		start := time.Now()
		path := c.Request.URL.Path // Keep original request path in case of http.StripPrefix.

		c.Next()

		log.Printf(colors.ResetAll+"  %s   %s   %s  %s", fmtDuration(start), fmtStatus(c), fmtMethod(c), fmtPath(path))
	})
}

func fmtDuration(start time.Time) string {
	return fmt.Sprintf("%s%s%13s%s", colors.ResetAll, colors.ResetAll+colors.Dim, time.Since(start), colors.ResetAll)
}

func fmtStatus(c *core.Context) string {
	code := coreutil.ResponseStatus(c)

	color := colors.White

	switch {
	case code >= 200 && code <= 299:
		color += colors.BackgroundGreen
	case code >= 300 && code <= 399:
		color += colors.BackgroundCyan
	case code >= 400 && code <= 499:
		color += colors.BackgroundYellow
	default:
		color += colors.BackgroundRed
	}

	return fmt.Sprintf("%s%s %3d %s", colors.ResetAll, color, code, colors.ResetAll)
}

func fmtMethod(c *core.Context) string {
	var color string

	switch c.Request.Method {
	case "GET":
		color += colors.Green
	case "POST":
		color += colors.Cyan
	case "PUT", "PATCH":
		color += colors.Blue
	case "DELETE":
		color += colors.Red
	}

	return fmt.Sprintf("%s%s%s%s", colors.ResetAll, color, c.Request.Method, colors.ResetAll)
}

func fmtPath(path string) string {
	return fmt.Sprintf("%s%s%s%s", colors.ResetAll, colors.Dim, path, colors.ResetAll)
}
