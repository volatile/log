package log

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/volatile/core"
	"github.com/volatile/core/coreutil"
	"github.com/whitedevops/colors"
)

// Use tells the core to use this handler.
func Use() {
	core.Use(func(c *core.Context) {
		start := time.Now()

		c.Next()

		log.Printf(colors.ResetAll+" %s  %s  %s  %s"+colors.ResetAll, fmtMethod(c), fmtPath(c), fmtStatus(c), fmtDuration(start))
	})
}

func fmtMethod(c *core.Context) string {
	s := colors.ResetAll

	switch strings.ToUpper(c.Request.Method) {
	case "GET":
		s += colors.Green + "GET"
	case "POST":
		s += colors.Cyan + "POST"
	case "PUT":
		s += colors.Blue + "PUT"
	case "PATCH":
		s += colors.Blue + "PATCH"
	case "DELETE":
		s += colors.Red + "DELETE"
	case "HEAD":
		s += colors.Magenta + "HEAD"
	case "OPTIONS":
		s += colors.Magenta + "OPTIONS"
	case "TRACE":
		s += colors.Magenta + "TRACE"
	default:
		s += colors.Red + colors.Blink + "UNKNOWN"
	}

	return s + colors.ResetAll
}

func fmtPath(c *core.Context) string {
	return colors.ResetAll + c.Request.URL.String()
}

func fmtStatus(c *core.Context) string {
	code := coreutil.ResponseStatus(c.ResponseWriter)

	s := colors.ResetAll

	switch {
	case code >= 200 && code <= 299:
		s += colors.Green
	case code >= 300 && code <= 399:
		s += colors.Cyan
	case code >= 400 && code <= 499:
		s += colors.Yellow
	default:
		s += colors.Red + colors.Blink
	}

	return s + strconv.Itoa(code) + colors.ResetAll
}

func fmtDuration(start time.Time) string {
	return colors.ResetAll + colors.Dim + time.Since(start).String() + colors.ResetAll
}
