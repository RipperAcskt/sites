package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"sites/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"

	"sites/internal/service"
)

func (h *Handler) SingUp(c *gin.Context) {
	logger := getLogger(c)

	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.s.SingUp(c.Request.Context(), user)
	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		logger.Error("/users/auth/sing-up", zap.Error(fmt.Errorf("service sing up failed: %w", err)))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) SingIn(c *gin.Context) {
	logger := getLogger(c)

	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := h.s.SingIn(c.Request.Context(), user)
	if err != nil {
		if errors.Is(err, service.ErrUserDoesNotExists) || errors.Is(err, service.ErrIncorrectPassword) {

			c.JSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}
		logger.Error("/users/auth/sing-in", zap.Error(fmt.Errorf("service sing in failed: %w", err)))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  token.Access,
		"refresh_token": token.RT,
		"admin":         token.Admin,
	})
}

func (h *Handler) VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := getLogger(c)

		token := strings.Split(c.GetHeader("Authorization"), " ")
		if len(token) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Errorf("access token required").Error(),
			})
			return
		}
		accessToken := token[1]

		id, err := service.Verify(accessToken, h.Cfg)
		if err != nil {
			if errors.Is(err, service.ErrTokenExpired) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				return
			}
			if strings.Contains(err.Error(), jwt.ErrSignatureInvalid.Error()) {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"error": fmt.Errorf("wrong signature").Error(),
				})
				return
			}

			logger.Error("service verify failed", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Errorf("verify failed: %w", err).Error(),
			})
			return
		}

		c.Set("id", fmt.Sprint(id))

		c.Next()

	}
}

func (h *Handler) Refresh(c *gin.Context) {
	logger := getLogger(c)

	var ref struct {
		Refresh_token string `json:"refresh_token"`
	}
	if err := c.BindJSON(&ref); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := service.Verify(ref.Refresh_token, h.Cfg)
	if err != nil {
		if errors.Is(err, service.ErrTokenExpired) {

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		if strings.Contains(err.Error(), jwt.ErrSignatureInvalid.Error()) {

			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": fmt.Errorf("wrong signature").Error(),
			})
			return
		}

		logger.Error("/users/auth/refresh", zap.Error(fmt.Errorf("verify failed: %w", err)))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("verify rt failed: %w", err).Error(),
		})
		return
	}

	params := service.TokenParams{
		ID:                id,
		Type:              service.User,
		HS256_SECRET:      h.Cfg.HS256_SECRET,
		ACCESS_TOKEN_EXP:  h.Cfg.ACCESS_TOKEN_EXP,
		REFRESH_TOKEN_EXP: h.Cfg.REFRESH_TOKEN_EXP,
	}

	token, err := service.NewToken(params)
	if err != nil {
		if errors.Is(err, service.ErrUnknownType) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		logger.Error("/users/auth/refresh", zap.Error(fmt.Errorf("new token failed: %w", err)))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	exp := int((time.Duration(h.Cfg.REFRESH_TOKEN_EXP) * time.Hour * 24).Seconds())
	c.SetCookie("refresh_token", token.RT, exp, "/users/auth", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"access_token": token.Access,
	})
}
