package log

import (
	"os"
	"time"

	"log/slog"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
		logger := slog.New(handler)
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		timestamp := time.Now().Format("2006/01/02 - 15:04:05")

		logger.Info(
			"request_details",
			slog.String("method", method),
			slog.Int("status", status),
			slog.String("path", path),
			slog.Duration("latency", latency),
			slog.String("client_ip", clientIP),
			slog.String("timestamp", timestamp),
		)
	}
}
