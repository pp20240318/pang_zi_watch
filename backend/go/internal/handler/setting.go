package handler

import (
	"net/http"

	"watch/backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type siteSettingPayload struct {
	SiteName         string `json:"siteName"`
	SiteSlogan       string `json:"siteSlogan"`
	BrandDescription string `json:"brandDescription"`
	ServicePhone     string `json:"servicePhone"`
	ServiceHours     string `json:"serviceHours"`
	Copyright        string `json:"copyright"`
	IcpNumber        string `json:"icpNumber"`
	IcpLink          string `json:"icpLink"`
	ContactEmail     string `json:"contactEmail"`
	BusinessEmail    string `json:"businessEmail"`
	CompanyAddress   string `json:"companyAddress"`
}

func getSiteSetting(db *gorm.DB) (models.SiteSetting, error) {
	var s models.SiteSetting
	err := db.First(&s, 1).Error
	return s, err
}

func GetSiteSettings(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		s, err := getSiteSetting(db)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "settings not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": s})
	}
}

func UpdateSiteSettings(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		s, err := getSiteSetting(db)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "settings not found"})
			return
		}
		var req siteSettingPayload
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		s.SiteName = req.SiteName
		s.SiteSlogan = req.SiteSlogan
		s.BrandDescription = req.BrandDescription
		s.ServicePhone = req.ServicePhone
		s.ServiceHours = req.ServiceHours
		s.Copyright = req.Copyright
		s.IcpNumber = req.IcpNumber
		s.IcpLink = req.IcpLink
		s.ContactEmail = req.ContactEmail
		s.BusinessEmail = req.BusinessEmail
		s.CompanyAddress = req.CompanyAddress
		db.Save(&s)
		c.JSON(http.StatusOK, gin.H{"data": s})
	}
}

func PublicGetSiteSettings(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		s, err := getSiteSetting(db)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "settings not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": s})
	}
}
