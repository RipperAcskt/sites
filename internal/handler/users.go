package handler

import (
	"errors"
	"fmt"
	"net/http"

	"sites/internal/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) GetProfile(c *gin.Context) {
	logger := getLogger(c)

	k, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "verefy token failed",
		})
	}
	user, p, err := h.s.GetProfile(c.Request.Context(), k.(string))
	if err != nil {
		if errors.Is(err, service.ErrUserDoesNotExists) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		logger.Error("/users/profile/{id}", zap.Error(fmt.Errorf("get profile failed: %w", err)))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("get profile failed: %w", err).Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":     user,
		"payments": p,
	})
}
