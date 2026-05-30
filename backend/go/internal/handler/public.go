package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"watch/backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PublicListBanners(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var banners []models.Banner
		db.Where("enabled = ?", true).Order("sort asc, id asc").Find(&banners)
		c.JSON(http.StatusOK, gin.H{"data": banners})
	}
}

type publicBrand struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Count   int64  `json:"count,omitempty"`
	Contact bool   `json:"contact,omitempty"`
}

func PublicListPages(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pages []models.ContentPage
		q := db.Where("enabled = ?", true).Order("category asc, sort asc, id asc")
		if cat := c.Query("category"); cat != "" {
			q = q.Where("category = ?", cat)
		}
		q.Select("id, slug, title, category, sort").Find(&pages)
		c.JSON(http.StatusOK, gin.H{"data": pages})
	}
}

func PublicGetPage(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")
		var page models.ContentPage
		if err := db.Where("slug = ? AND enabled = ?", slug, true).First(&page).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": page})
	}
}

func PublicListBrands(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var totalProducts int64
		db.Model(&models.Product{}).Where("enabled = ?", true).Count(&totalProducts)

		out := []publicBrand{{ID: "all", Name: "全部系列", Count: totalProducts}}

		var brands []models.Brand
		db.Where("enabled = ?", true).Order("sort asc, id asc").Find(&brands)
		for _, b := range brands {
			var count int64
			db.Model(&models.Product{}).Where("brand_code = ? AND enabled = ?", b.Code, true).Count(&count)
			item := publicBrand{ID: b.Code, Name: b.Name, Count: count, Contact: b.Contact}
			out = append(out, item)
		}
		c.JSON(http.StatusOK, gin.H{"data": out})
	}
}

func PublicListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := db.Where("enabled = ?", true)
		if brand := c.Query("brand"); brand != "" && brand != "all" {
			q = q.Where("brand_code = ?", brand)
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

func PublicGetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
		var p models.Product
		if err := db.Where("enabled = ?", true).First(&p, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": formatProduct(p)})
	}
}

type createOrderRequest struct {
	ProductID     uint   `json:"productId"`
	VariantID     string `json:"variantId"`
	VariantLabel  string `json:"variantLabel"`
	PayAmount     float64 `json:"payAmount"`
	Customer      struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
		Remark  string `json:"remark"`
	} `json:"customer"`
}

func PublicCreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createOrderRequest
		if err := c.ShouldBindJSON(&req); err != nil || req.ProductID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		var p models.Product
		if err := db.Where("enabled = ?", true).First(&p, req.ProductID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
			return
		}
		payAmount := req.PayAmount
		if payAmount <= 0 {
			payAmount = p.Price
		}
		order := models.Order{
			OrderNo:         fmt.Sprintf("PZ%d", time.Now().UnixMilli()),
			Status:          models.OrderStatusPending,
			ProductID:       p.ID,
			ProductName:     p.Name,
			ProductImage:    p.Image,
			ProductPrice:    p.Price,
			VariantID:       req.VariantID,
			VariantLabel:    req.VariantLabel,
			PayAmount:       payAmount,
			CustomerName:    req.Customer.Name,
			CustomerEmail:   req.Customer.Email,
			CustomerPhone:   req.Customer.Phone,
			CustomerAddress: req.Customer.Address,
			CustomerRemark:  req.Customer.Remark,
		}
		if err := db.Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": publicOrderResponse(order)})
	}
}

type payOrderRequest struct {
	PaymentMethod string `json:"paymentMethod"`
}

func PublicPayOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderNo := c.Param("orderNo")
		var order models.Order
		if err := db.Where("order_no = ?", orderNo).First(&order).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
			return
		}
		if order.Status != models.OrderStatusPending {
			c.JSON(http.StatusBadRequest, gin.H{"error": "order already processed"})
			return
		}
		var req payOrderRequest
		c.ShouldBindJSON(&req)
		now := time.Now()
		order.Status = models.OrderStatusPaid
		order.PaymentMethod = req.PaymentMethod
		order.PaidAt = &now
		db.Save(&order)
		c.JSON(http.StatusOK, gin.H{"data": publicOrderResponse(order)})
	}
}

func publicOrderResponse(o models.Order) gin.H {
	return gin.H{
		"orderNo": o.OrderNo,
		"status":  o.Status,
		"product": gin.H{
			"id":    o.ProductID,
			"name":  o.ProductName,
			"image": o.ProductImage,
			"price": o.ProductPrice,
		},
		"payAmount": o.PayAmount,
		"customer": gin.H{
			"name":    o.CustomerName,
			"email":   o.CustomerEmail,
			"phone":   o.CustomerPhone,
			"address": o.CustomerAddress,
			"remark":  o.CustomerRemark,
		},
		"paymentMethod": o.PaymentMethod,
		"paidAt":        o.PaidAt,
		"createdAt":     o.CreatedAt,
	}
}
