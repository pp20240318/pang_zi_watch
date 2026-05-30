package models

import "time"

type Admin struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:64" json:"username"`
	Password  string    `json:"-"`
	RoleID    uint      `json:"role_id"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;size:64" json:"name"`
	Permissions string    `gorm:"type:text" json:"permissions"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Brand struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"uniqueIndex;size:64" json:"code"`
	Name      string    `gorm:"size:128" json:"name"`
	Sort      int       `gorm:"default:0" json:"sort"`
	Contact   bool      `gorm:"default:false" json:"contact"`
	Enabled   bool      `gorm:"default:true" json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	BrandCode    string    `gorm:"size:64;index" json:"brandId"`
	Name         string    `gorm:"size:256" json:"name"`
	Price        float64   `json:"price"`
	Image        string    `gorm:"size:512" json:"image"`
	Gallery      string    `gorm:"type:text" json:"gallery"`
	Rating       float64   `gorm:"default:0" json:"rating"`
	ReviewCount  int       `gorm:"default:0" json:"reviews"`
	Tag          string    `gorm:"size:32" json:"tag"`
	Description  string    `gorm:"type:text" json:"description"`
	Specs        string    `gorm:"type:text" json:"specs"`
	Variants     string    `gorm:"type:text" json:"variants"`
	Enabled      bool      `gorm:"default:true" json:"enabled"`
	Sort         int       `gorm:"default:0" json:"sort"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ContentPage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Slug      string    `gorm:"uniqueIndex;size:64" json:"slug"`
	Title     string    `gorm:"size:128" json:"title"`
	Category  string    `gorm:"size:64;index" json:"category"`
	Body      string    `gorm:"type:longtext" json:"body"`
	Sort      int       `gorm:"default:0" json:"sort"`
	Enabled   bool      `gorm:"default:true" json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Banner struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"size:128" json:"title"`
	Subtitle  string    `gorm:"size:256" json:"subtitle"`
	Image     string    `gorm:"size:512" json:"image"`
	CTA       string    `gorm:"size:64" json:"cta"`
	Link      string    `gorm:"size:512" json:"link"`
	Sort      int       `gorm:"default:0" json:"sort"`
	Enabled   bool      `gorm:"default:true" json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	OrderNo         string     `gorm:"uniqueIndex;size:32" json:"order_no"`
	Status          string     `gorm:"size:32;index" json:"status"`
	ProductID       uint       `json:"product_id"`
	ProductName     string     `gorm:"size:256" json:"product_name"`
	ProductImage    string     `gorm:"size:512" json:"product_image"`
	ProductPrice    float64    `json:"product_price"`
	VariantID       string     `gorm:"size:64" json:"variant_id"`
	VariantLabel    string     `gorm:"size:128" json:"variant_label"`
	PayAmount       float64    `json:"pay_amount"`
	CustomerName    string     `gorm:"size:64" json:"customer_name"`
	CustomerEmail   string     `gorm:"size:128" json:"customer_email"`
	CustomerPhone   string     `gorm:"size:32" json:"customer_phone"`
	CustomerAddress string     `gorm:"type:text" json:"customer_address"`
	CustomerRemark  string     `gorm:"type:text" json:"customer_remark"`
	PaymentMethod   string     `gorm:"size:32" json:"payment_method"`
	PaidAt          *time.Time `json:"paid_at"`
	AdminRemark     string     `gorm:"type:text" json:"admin_remark"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

const (
	OrderStatusPending   = "pending"
	OrderStatusPaid      = "paid"
	OrderStatusShipped   = "shipped"
	OrderStatusCompleted = "completed"
	OrderStatusCancelled = "cancelled"
)
