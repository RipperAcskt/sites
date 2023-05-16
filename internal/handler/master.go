package handler

import (
	"errors"
	"fmt"
	"net/http"
	"sites/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) Start(c *gin.Context) {
	logger := getLogger(c)

	var model struct {
		Email string  `json:"Email"`
		Price float64 `json:"price"`
		Time  int     `json:"Time"`
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

	err := h.s.CreateMaster(c.Request.Context(), id.(string), model.Email, model.Price, model.Time)
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

func (h *Handler) GetTasks(c *gin.Context) {
	k, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "verefy token failed",
		})
		return
	}
	tasks := h.s.GetTasks(k.(string))
	var info []struct {
		Id     int
		Status string
		Email  string
		Time   int
	}
	var tmp struct {
		Id     int
		Status string
		Email  string
		Time   int
	}
	for i, task := range tasks {
		tmp.Id = i
		tmp.Status = task.Status
		tmp.Email = task.Email
		tmp.Time = task.Time
		info = append(info, tmp)
	}

	c.JSON(http.StatusOK, gin.H{
		"userID": k,
		"tasks:": info,
	})
}

func (h *Handler) StopTask(c *gin.Context) {
	k, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "verefy token failed",
		})
		return
	}

	var status struct {
		Status string `json:"Status"`
	}

	if err := c.BindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	orderID := c.Param("id")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	if status.Status == "stop" {
		err = h.s.StopTask(k.(string), orderIDInt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}
		c.Status(http.StatusOK)
		return
	}
	if status.Status == "start" {
		tasks := h.s.GetTasks(k.(string))
		if len(tasks) <= orderIDInt {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": service.ErrNoTask,
			})
			return
		}
		h.s.MasterProgram(&tasks[orderIDInt], tasks[orderIDInt].Email, tasks[orderIDInt].Time)
	}

}

func (h *Handler) DeleteTask(c *gin.Context) {
	k, ok := c.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "verefy token failed",
		})
		return
	}

	orderID := c.Param("id")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	h.s.DeleteTask(k.(string), orderIDInt)

}
