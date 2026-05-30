package handler

import (
	"net/http"
	"strconv"

	"watch/backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type bannerPayload struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Image    string `json:"image"`
	CTA      string `json:"cta"`
	Link     string `json:"link"`
	Sort     int    `json:"sort"`
	Enabled  *bool  `json:"enabled"`
}

func ListBanners(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var banners []models.Banner
		db.Order("sort asc, id asc").Find(&banners)
		c.JSON(http.StatusOK, gin.H{"data": banners})
	}
}

func CreateBanner(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req bannerPayload
		if err := c.ShouldBindJSON(&req); err != nil || req.Title == "" || req.Image == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		enabled := true
		if req.Enabled != nil {
			enabled = *req.Enabled
		}
		b := models.Banner{
			Title: req.Title, Subtitle: req.Subtitle, Image: req.Image,
			CTA: req.CTA, Link: req.Link, Sort: req.Sort, Enabled: enabled,
		}
		db.Create(&b)
		c.JSON(http.StatusOK, gin.H{"data": b})
	}
}

func UpdateBanner(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		var b models.Banner
		if err := db.First(&b, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		var req bannerPayload
		c.ShouldBindJSON(&req)
		if req.Title != "" {
			b.Title = req.Title
		}
		b.Subtitle = req.Subtitle
		if req.Image != "" {
			b.Image = req.Image
		}
		b.CTA = req.CTA
		b.Link = req.Link
		b.Sort = req.Sort
		if req.Enabled != nil {
			b.Enabled = *req.Enabled
		}
		db.Save(&b)
		c.JSON(http.StatusOK, gin.H{"data": b})
	}
}

func DeleteBanner(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		db.Delete(&models.Banner{}, id)
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}
