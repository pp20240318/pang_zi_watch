package handler

import (
	"net/http"
	"strconv"

	"watch/backend/internal/auth"
	"watch/backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ListAdmins(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var admins []models.Admin
		if err := db.Preload("Role").Find(&admins).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": admins})
	}
}

type createAdminRequest struct {
	Username string `json:"username" validate:"required,min=2,max=64"`
	Password string `json:"password" validate:"required,min=6"`
	RoleID   uint   `json:"role_id" validate:"required"`
}

func CreateAdmin(db *gorm.DB) gin.HandlerFunc {
	validate := validator.New()
	return func(c *gin.Context) {
		var req createAdminRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var role models.Role
		if err := db.First(&role, req.RoleID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "role not found"})
			return
		}
		var existing models.Admin
		if err := db.Where("username = ?", req.Username).First(&existing).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
			return
		}
		hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "hash failed"})
			return
		}
		admin := models.Admin{Username: req.Username, Password: string(hashed), RoleID: req.RoleID}
		if err := db.Create(&admin).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		db.Preload("Role").First(&admin, admin.ID)
		c.JSON(http.StatusOK, gin.H{"data": admin})
	}
}

type updateAdminRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
}

func UpdateAdmin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		var admin models.Admin
		if err := db.First(&admin, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "admin not found"})
			return
		}
		var req updateAdminRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		if req.Username != "" && req.Username != admin.Username {
			var existing models.Admin
			if err := db.Where("username = ? AND id != ?", req.Username, id).First(&existing).Error; err == nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
				return
			}
			admin.Username = req.Username
		}
		if req.Password != "" {
			hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
			admin.Password = string(hashed)
		}
		if req.RoleID > 0 {
			admin.RoleID = req.RoleID
		}
		db.Save(&admin)
		db.Preload("Role").First(&admin, admin.ID)
		c.JSON(http.StatusOK, gin.H{"data": admin})
	}
}

func DeleteAdmin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(*auth.Claims)
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		if claims.AdminID == uint(id) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete yourself"})
			return
		}
		if err := db.Delete(&models.Admin{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}

func ListRoles(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var roles []models.Role
		db.Find(&roles)
		c.JSON(http.StatusOK, gin.H{"data": roles})
	}
}

type roleRequest struct {
	Name        string `json:"name" validate:"required"`
	Permissions string `json:"permissions" validate:"required"`
}

func CreateRole(db *gorm.DB) gin.HandlerFunc {
	validate := validator.New()
	return func(c *gin.Context) {
		var req roleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		role := models.Role{Name: req.Name, Permissions: req.Permissions}
		if err := db.Create(&role).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": role})
	}
}

func UpdateRole(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		var role models.Role
		if err := db.First(&role, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
			return
		}
		var req roleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		if req.Name != "" {
			role.Name = req.Name
		}
		if req.Permissions != "" {
			role.Permissions = req.Permissions
		}
		db.Save(&role)
		c.JSON(http.StatusOK, gin.H{"data": role})
	}
}

func DeleteRole(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		var count int64
		db.Model(&models.Admin{}).Where("role_id = ?", id).Count(&count)
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "role in use"})
			return
		}
		db.Delete(&models.Role{}, id)
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}
