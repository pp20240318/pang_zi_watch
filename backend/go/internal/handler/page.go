package handler

import (
	"net/http"
	"strconv"

	"watch/backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type pagePayload struct {
	Slug     string `json:"slug"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Body     string `json:"body"`
	Sort     int    `json:"sort"`
	Enabled  *bool  `json:"enabled"`
}

func ListPages(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pages []models.ContentPage
		q := db.Order("category asc, sort asc, id asc")
		if cat := c.Query("category"); cat != "" {
			q = q.Where("category = ?", cat)
		}
		q.Find(&pages)
		c.JSON(http.StatusOK, gin.H{"data": pages})
	}
}

func CreatePage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req pagePayload
		if err := c.ShouldBindJSON(&req); err != nil || req.Slug == "" || req.Title == "" || req.Category == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		enabled := true
		if req.Enabled != nil {
			enabled = *req.Enabled
		}
		p := models.ContentPage{
			Slug: req.Slug, Title: req.Title, Category: req.Category,
			Body: req.Body, Sort: req.Sort, Enabled: enabled,
		}
		if err := db.Create(&p).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "slug already exists"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": p})
	}
}

func UpdatePage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		var p models.ContentPage
		if err := db.First(&p, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		var req pagePayload
		c.ShouldBindJSON(&req)
		if req.Slug != "" {
			p.Slug = req.Slug
		}
		if req.Title != "" {
			p.Title = req.Title
		}
		if req.Category != "" {
			p.Category = req.Category
		}
		p.Body = req.Body
		p.Sort = req.Sort
		if req.Enabled != nil {
			p.Enabled = *req.Enabled
		}
		if err := db.Save(&p).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "slug already exists"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": p})
	}
}

func DeletePage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		db.Delete(&models.ContentPage{}, id)
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}
