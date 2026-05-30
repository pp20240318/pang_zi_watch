package handler

import (
	"net/http"
	"strconv"

	"watch/backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := db.Model(&models.Order{})
		if status := c.Query("status"); status != "" {
			q = q.Where("status = ?", status)
		}
		if kw := c.Query("q"); kw != "" {
			like := "%" + kw + "%"
			q = q.Where("order_no LIKE ? OR customer_phone LIKE ? OR customer_email LIKE ?", like, like, like)
		}
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))
		if page < 1 {
			page = 1
		}
		if size < 1 || size > 100 {
			size = 20
		}
		var total int64
		q.Count(&total)
		var orders []models.Order
		q.Order("id desc").Offset((page - 1) * size).Limit(size).Find(&orders)
		c.JSON(http.StatusOK, gin.H{"data": orders, "total": total, "page": page, "size": size})
	}
}

func GetOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		var order models.Order
		if err := db.First(&order, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": order})
	}
}

type updateOrderRequest struct {
	Status      string `json:"status"`
	AdminRemark string `json:"admin_remark"`
}

func UpdateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		var order models.Order
		if err := db.First(&order, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		var req updateOrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		if req.Status != "" {
			order.Status = req.Status
		}
		if req.AdminRemark != "" {
			order.AdminRemark = req.AdminRemark
		}
		db.Save(&order)
		c.JSON(http.StatusOK, gin.H{"data": order})
	}
}
