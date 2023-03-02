package log

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GinLogger() gin.HandlerFunc {
	logger := NewLogger("gin")
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
		bodySize := c.Writer.Size()

		if statusCode >= http.StatusBadRequest {
			logger.Errorw(errorMessage,
				"method", method,
				"status", statusCode,
				"latency", latency,
				"path", path,
				"raw", raw,
				"body-size", bodySize,
				"client-ip", clientIP,
			)
		} else {
			logger.Infow("",
				"method", method,
				"status", statusCode,
				"latency", latency,
				"path", path,
				"raw", raw,
				"body-size", bodySize,
				"client-ip", clientIP,
			)
		}
	}
}
