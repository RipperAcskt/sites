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

	router.Use(h.Log())

	auth := router.Group("/auth")
	auth.POST("sing-up", h.SingUp)
	auth.POST("sing-in", h.SingIn)
	auth.POST("refresh", h.Refresh)

	router.GET("/profile", h.VerifyToken(), h.GetProfile)

	router.POST("/payment", h.VerifyToken(), h.CreatePayment)
	router.GET("/payment", h.VerifyToken(), h.GetPayment)

	router.POST("/flood", h.VerifyToken(), h.Start)

	return router
}
