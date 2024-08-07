package middleware

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type RLogger struct {}
func (r * RLogger)getDurationInMilliseconds(start time.Time) float64 {
	duration := time.Since(start)

	return float64(duration.Nanoseconds()) / 1000000
}

func (rl *RLogger) LogRequest(c *gin.Context) {
    method := c.Request.Method
    startTime := time.Now()

    // Capture the request body
    var requestBody []byte
    if c.Request.Body != nil {
        bodyBytes, err := io.ReadAll(c.Request.Body)
        if err == nil {
            requestBody = bodyBytes
            // Restore the request body
            c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
        }
    }

    // Proceed with the request
    c.Next()

    // Log request details after response is finished
    c.Writer.WriteHeaderNow()
    duration := rl.getDurationInMilliseconds(startTime)
    log.Printf("Request fulfilled: method=%s, url=%s, duration=%.2fms, status=%d",
        method, c.Request.URL.Path, duration, c.Writer.Status())

    // Log request body and query parameters if they exist
    if len(requestBody) > 0 {
        log.Printf("Request body: %s", string(requestBody))
    }
    if len(c.Request.URL.Query()) > 0 {
        log.Printf("Query parameters: %v", c.Request.URL.Query())
    }

    // Log headers, omitting specific ones
    headers := make(map[string]string)
    for k, v := range c.Request.Header {
        if k != "User-Agent" && k != "Authorization" && k != "Accept" &&
            k != "Accept-Encoding" && k != "Accept-Language" && k != "If-None-Match" {
            headers[k] = v[0]
        }
    }
    log.Printf("Headers: %v", headers)
}