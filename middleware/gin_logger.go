package middleware

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type GinLogger struct {
	logger *slog.Logger
}

func (gl GinLogger) Middleware(ctx *gin.Context) {
	ts := make([]byte, 30)

	gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			path := param.Path

			return fmt.Sprintf("%v |%3d| %13v | %15s |%-7s %#v requestID=%s, userID=%s\n",
				string(param.TimeStamp.AppendFormat(ts[:0], "2006/01/02 - 15:04:05")),
				param.StatusCode,
				param.Latency,
				param.ClientIP,
				param.Method,
				path,
				ctx.GetString("hash"),
				ctx.GetString("userID"),
			)
		},
	})(ctx)
}

func NewGinLogger(logger *slog.Logger) *GinLogger {
	return &GinLogger{logger: logger}
}
