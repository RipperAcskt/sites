package handler

import (
	"errors"
	"fmt"
	"net/http"
	"sites/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) Start(c *gin.Context) {
	logger := getLogger(c)

	var model struct {
		Email string  `json:"email"`
		Price float64 `json:"price"`
		Time  int     `json:"time"`
	}

	if err := c.BindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "verefy token failed",
		})
		return
	}

	err := h.s.Start(c.Request.Context(), id.(string), model.Email, model.Price, model.Time)
	if err != nil {
		if errors.Is(err, service.ErrLowBalance) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		logger.Error("/payment", zap.Error(fmt.Errorf("start failed: %w", err)))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("start failed: %w", err).Error(),
		})
		return
	}
}
