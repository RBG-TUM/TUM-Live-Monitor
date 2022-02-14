package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func logger(c *gin.Context) {
	start := time.Now()
	c.Next()

	var m string
	switch c.Request.Method {
	case http.MethodGet:
		m = getf(" GET  ")
	case http.MethodPost:
		m = postf(" POST ")
	default:
		m = defaultf(c.Request.Method)
	}
	status := c.Writer.Status()
	st := fmt.Sprintf(" %d ", status)
	if status >= 100 && status < 200 {
		st = http1xx(fmt.Sprintf(" %d ", status))
	} else if status >= 200 && status < 300 {
		st = http2xx(fmt.Sprintf(" %d ", status))
	} else if status >= 300 && status < 400 {
		st = http3xx(fmt.Sprintf(" %d ", status))
	} else if status >= 400 && status < 500 {
		st = http4xx(fmt.Sprintf(" %d ", status))
	} else if status >= 500 && status < 600 {
		st = http5xx(fmt.Sprintf(" %d ", status))
	}

	log.Infof("%s %s\t|\t%s\t|\t%dms", m, c.Request.RequestURI, st, time.Since(start).Milliseconds())
}
