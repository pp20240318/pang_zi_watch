package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
			return
		}
		ext := strings.ToLower(filepath.Ext(file.Filename))
		switch ext {
		case ".jpg", ".jpeg", ".png", ".gif", ".webp":
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file type"})
			return
		}
		uploadDir := "uploads"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "create upload dir failed"})
			return
		}
		filename := time.Now().Format("20060102_150405") + "_" + strings.ReplaceAll(file.Filename, " ", "_")
		savePath := filepath.Join(uploadDir, filename)
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "save file failed"})
			return
		}
		url := "/" + filepath.ToSlash(filepath.Join(uploadDir, filename))
		c.JSON(http.StatusOK, gin.H{"url": url})
	}
}
