package handler

import (
	"net/http"
	"strconv"

	"watch/backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := db.Model(&models.Product{})
		if brand := c.Query("brand"); brand != "" {
			q = q.Where("brand_code = ?", brand)
		}
		if enabled := c.Query("enabled"); enabled != "" {
			q = q.Where("enabled = ?", enabled == "true")
		}
		var products []models.Product
		q.Order("sort asc, id asc").Find(&products)
		out := make([]productResponse, 0, len(products))
		for _, p := range products {
			out = append(out, formatProduct(p))
		}
		c.JSON(http.StatusOK, gin.H{"data": out})
	}
}

func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		var p models.Product
		if err := db.First(&p, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": formatProduct(p)})
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req productPayload
		if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" || req.BrandID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		p := productFromPayload(req)
		if p.Specs == "" {
			p.Specs = defaultSpecsJSON()
		}
		if p.Variants == "" {
			p.Variants = defaultVariantsJSON()
		}
		if err := db.Create(&p).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": formatProduct(p)})
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		var p models.Product
		if err := db.First(&p, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		var req productPayload
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		if req.BrandID != "" {
			p.BrandCode = req.BrandID
		}
		if req.Name != "" {
			p.Name = req.Name
		}
		if req.Price > 0 {
			p.Price = req.Price
		}
		if req.Image != "" {
			p.Image = req.Image
		}
		if req.Gallery != nil {
			p.Gallery = mustJSONArray(req.Gallery)
		}
		if req.Rating > 0 {
			p.Rating = req.Rating
		}
		if req.Reviews > 0 {
			p.ReviewCount = req.Reviews
		}
		p.Tag = req.Tag
		if req.Description != "" {
			p.Description = req.Description
		}
		if req.Specs != nil {
			p.Specs = mustJSONObjects(req.Specs)
		}
		if req.Variants != nil {
			p.Variants = mustJSONObjects(req.Variants)
		}
		if req.Enabled != nil {
			p.Enabled = *req.Enabled
		}
		p.Sort = req.Sort
		db.Save(&p)
		c.JSON(http.StatusOK, gin.H{"data": formatProduct(p)})
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		db.Delete(&models.Product{}, id)
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}
