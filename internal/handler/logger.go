package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (h *Handler) Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		log, _ := zap.NewProduction(zap.Fields(zap.String("url", c.Request.URL.Path), zap.String("method", c.Request.Method), zap.Any("uuid", uuid.New()), zap.String("request time", time.Now().String())))
		c.Set("logger", *log)
		c.Next()

		log.Info("request", zap.String("time", time.Since(start).String()))
	}
}

func getLogger(c *gin.Context) *zap.Logger {
	tmp, _ := c.Get("logger")
	logger := tmp.(zap.Logger)
	return &logger
}
