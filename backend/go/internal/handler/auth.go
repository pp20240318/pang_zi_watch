package handler

import (
	"net/http"

	"watch/backend/internal/auth"
	"watch/backend/internal/config"
	"watch/backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type changePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required,min=2"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

func LoginHandler(db *gorm.DB, cfg config.Config) gin.HandlerFunc {
	validate := validator.New()
	return func(c *gin.Context) {
		var req loginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		admin, err := service.Login(db, req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		token, err := auth.GenerateToken(cfg.JWTSecret, cfg.JWTExpiry, admin.ID, admin.Username, admin.Role.Name, admin.Role.Permissions)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "token error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func ChangePasswordHandler(db *gorm.DB) gin.HandlerFunc {
	validate := validator.New()
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(*auth.Claims)
		var req changePasswordRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := service.ChangePassword(db, claims.AdminID, req.OldPassword, req.NewPassword); err != nil {
			if err == service.ErrInvalidCredentials {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "old password incorrect"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "password changed successfully"})
	}
}
