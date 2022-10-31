package http

import (
	"fmt"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	Dump      bool
	SkipPaths []string
}

func NewLogger() *Logger {
	l := &Logger{}
	l.Dump = false
	l.SkipPaths = []string{
		"/crash/dataUpload",
		"/sdk/dataUpload",
		"/log",
	}
	return l
}

func (l *Logger) HandlerFunc() gin.HandlerFunc {
	var skip map[string]struct{}
	if length := len(l.SkipPaths); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range l.SkipPaths {
			skip[path] = struct{}{}
		}
	}
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if _, ok := skip[path]; !ok {
			if l.Dump {
				fmt.Println("\n\n\n\n*********************************************")
			}
			log.Debug().Str("uri", c.Request.RequestURI).Str("method", c.Request.Method).Str("ip", c.ClientIP()).Msg("Request")
			if l.Dump {
				fmt.Println("\n\nRequest:")
				b, _ := httputil.DumpRequest(c.Request, true)
				fmt.Println(string(b))
			}
		}
	}
}
