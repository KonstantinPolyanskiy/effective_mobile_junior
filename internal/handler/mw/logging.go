package mw

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Logging(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		l := log.With(
			zap.String("method", c.Request.Method),
			zap.String("path", c.FullPath()),
			zap.String("addr", c.Request.RemoteAddr),
		)

		c.Next()

		l.Info("request completed",
			zap.Duration("duration", time.Since(t)),
		)
	}
}
