package handler

import (
	"net/http"
	"strconv"

	"watch/backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type brandPayload struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Sort    int    `json:"sort"`
	Contact bool   `json:"contact"`
	Enabled *bool  `json:"enabled"`
}

func ListBrands(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var brands []models.Brand
		db.Order("sort asc, id asc").Find(&brands)
		c.JSON(http.StatusOK, gin.H{"data": brands})
	}
}

func CreateBrand(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req brandPayload
		if err := c.ShouldBindJSON(&req); err != nil || req.Code == "" || req.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		enabled := true
		if req.Enabled != nil {
			enabled = *req.Enabled
		}
		brand := models.Brand{Code: req.Code, Name: req.Name, Sort: req.Sort, Contact: req.Contact, Enabled: enabled}
		if err := db.Create(&brand).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "create failed, code may exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": brand})
	}
}

func UpdateBrand(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		var brand models.Brand
		if err := db.First(&brand, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		var req brandPayload
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		if req.Name != "" {
			brand.Name = req.Name
		}
		if req.Code != "" {
			brand.Code = req.Code
		}
		brand.Sort = req.Sort
		brand.Contact = req.Contact
		if req.Enabled != nil {
			brand.Enabled = *req.Enabled
		}
		db.Save(&brand)
		c.JSON(http.StatusOK, gin.H{"data": brand})
	}
}

func DeleteBrand(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		var brand models.Brand
		if err := db.First(&brand, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		var count int64
		db.Model(&models.Product{}).Where("brand_code = ?", brand.Code).Count(&count)
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "brand has products"})
			return
		}
		db.Delete(&brand)
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}
