package handler

import (
	"net/http"
	"time"

	"watch/backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DashboardStats(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var productCount, brandCount, bannerCount int64
		db.Model(&models.Product{}).Count(&productCount)
		db.Model(&models.Brand{}).Count(&brandCount)
		db.Model(&models.Banner{}).Count(&bannerCount)

		var pendingOrders, paidOrders int64
		db.Model(&models.Order{}).Where("status = ?", models.OrderStatusPending).Count(&pendingOrders)
		db.Model(&models.Order{}).Where("status = ?", models.OrderStatusPaid).Count(&paidOrders)

		start := time.Now().Truncate(24 * time.Hour)
		var todayOrders int64
		var todaySales float64
		db.Model(&models.Order{}).Where("created_at >= ?", start).Count(&todayOrders)
		db.Model(&models.Order{}).Where("status IN ? AND paid_at >= ?", []string{models.OrderStatusPaid, models.OrderStatusShipped, models.OrderStatusCompleted}, start).
			Select("COALESCE(SUM(pay_amount),0)").Scan(&todaySales)

		c.JSON(http.StatusOK, gin.H{
			"product_count":  productCount,
			"brand_count":    brandCount,
			"banner_count":   bannerCount,
			"pending_orders": pendingOrders,
			"paid_orders":    paidOrders,
			"today_orders":   todayOrders,
			"today_sales":    todaySales,
		})
	}
}
