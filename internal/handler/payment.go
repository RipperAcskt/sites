package handler

import (
	"errors"
	"fmt"
	"net/http"
	"sites/internal/model"
	"sites/internal/service"
	"strconv"

	"github.com/02amanag/coinbase-commerce-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) CreatePayment(c *gin.Context) {
	logger := getLogger(c)

	var chargeModel model.Charge

	if err := c.BindJSON(&chargeModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	charge := coinbase.ChargeParam{
		Name:         chargeModel.Name,
		Description:  chargeModel.Description,
		Local_price:  chargeModel.Local_price,
		Pricing_type: "fixed_price",
		Metadata:     chargeModel.Metadata,
		Redirect_url: chargeModel.Redirect_url,
		Cancel_url:   chargeModel.Cancel_url,
	}

	k, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "verefy token failed",
		})
	}

	id, _ := strconv.ParseUint(k.(string), 10, 64)
	err := h.s.CreateCharge(c.Request.Context(), charge, id)
	if err != nil {
		if errors.Is(err, service.ErrUserDoesNotExists) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		logger.Error("/payment/:id", zap.Error(fmt.Errorf("create charge failed: %w", err)))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("create charge failed: %w", err).Error(),
		})
		return
	}
}

func (h *Handler) GetPayment(c *gin.Context) {
	logger := getLogger(c)

	k, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "verefy token failed",
		})
	}

	b, err := h.s.GetCharge(c.Request.Context(), k.(string))
	if err != nil {
		logger.Error("/payment", zap.Error(fmt.Errorf("get charge failed: %w", err)))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("get charge failed: %w", err).Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": b,
	})

}