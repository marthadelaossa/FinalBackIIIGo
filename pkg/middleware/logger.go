package middleware

import (
	"bytes"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// bodyLogWriter is a custom response writer with additional body logging capabilities.
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write overrides the Write method to log the response body.
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logger is a middleware that logs information about the incoming requests.
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Capture the request method, URL, and time
		method := ctx.Request.Method
		reqTime := time.Now()
		url := ctx.Request.URL
		var size int

		// Wrap the response writer with our custom body logger
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw

		// Process the request and move to the next middleware
		ctx.Next()

		// Capture the response size
		if ctx.Writer != nil {
			size = ctx.Writer.Size()
		}

		// Print request information
		fmt.Printf("\nPath:%s\nMethod:%s\nTime:%v\nRequest Size:%d\n", url, method, reqTime, size)

		// Print response body (if available)
		fmt.Printf("Response Body:%s\n", blw.body.String())
	}
}
