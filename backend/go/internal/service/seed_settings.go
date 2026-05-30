package service

import (
	"log"

	"watch/backend/internal/models"

	"gorm.io/gorm"
)

func EnsureSiteSettings(db *gorm.DB) {
	var count int64
	db.Model(&models.SiteSetting{}).Count(&count)
	if count > 0 {
		return
	}

	log.Println("seeding site settings...")
	s := models.SiteSetting{
		ID:               1,
		SiteName:         "胖子腕表",
		SiteSlogan:       "PANGZI WATCHES",
		BrandDescription: "专注正品名表销售，提供全国联保与专业售后服务。让每一秒都值得珍藏。",
		ServicePhone:     "400-888-6688",
		ServiceHours:     "9:00 - 21:00",
		Copyright:        "© 2026 胖子腕表 Pangzi Watches. 保留所有权利.",
		IcpNumber:        "京ICP备xxxxxxxx号",
		IcpLink:          "",
		ContactEmail:     "service@pangziwatch.com",
		BusinessEmail:    "business@pangziwatch.com",
		CompanyAddress:   "北京市朝阳区建国门外大街 1 号国贸写字楼 A 座 18 层",
	}
	db.Create(&s)
	log.Println("site settings seed completed")
}
