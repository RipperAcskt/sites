package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"sites/config"
	"sites/internal/service"
)

type Handler struct {
	s   *service.Service
	Cfg *config.Config
	log *zap.Logger
}

func New(s *service.Service, cfg *config.Config, log *zap.Logger) *Handler {
	return &Handler{s, cfg, log}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	router.Use(CORSMiddleware(), h.Log())

	auth := router.Group("/auth")
	auth.POST("sing-up", h.SingUp)
	auth.POST("sing-in", h.SingIn)
	auth.POST("refresh", h.Refresh)

	router.GET("/profile", h.VerifyToken(), h.GetProfile)

	router.POST("/payment", h.VerifyToken(), h.CreatePayment)
	router.GET("/payment", h.VerifyToken(), h.GetPayment)
	router.PUT("/payment/:id", h.VerifyToken(), h.CompletePayment)
	router.GET("/payments", h.VerifyToken(), h.GetPayments)

	router.GET("/tasks", h.VerifyToken(), h.GetTasks)
	router.PUT("/task/:id", h.VerifyToken(), h.StopTask)
	router.DELETE("/task/:id", h.VerifyToken(), h.DeleteTask)

	router.POST("/flood", h.VerifyToken(), h.Start)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
