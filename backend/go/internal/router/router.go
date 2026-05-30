package router

import (
	"net/http"

	"watch/backend/internal/config"
	"watch/backend/internal/handler"
	"watch/backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, cfg config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.Static("/uploads", "./uploads")

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	{
		api.POST("/login", handler.LoginHandler(db, cfg))

		public := api.Group("/public")
		{
			public.GET("/banners", handler.PublicListBanners(db))
			public.GET("/brands", handler.PublicListBrands(db))
			public.GET("/products", handler.PublicListProducts(db))
			public.GET("/products/:id", handler.PublicGetProduct(db))
			public.POST("/orders", handler.PublicCreateOrder(db))
			public.POST("/orders/:orderNo/pay", handler.PublicPayOrder(db))
		}

		authGroup := api.Group("")
		authGroup.Use(middleware.AuthRequired(cfg))
		{
			authGroup.POST("/change-password", handler.ChangePasswordHandler(db))
			authGroup.POST("/upload", handler.UploadHandler())

			authGroup.GET("/dashboard", middleware.RequirePermission("dashboard:view"), handler.DashboardStats(db))

			brands := authGroup.Group("/brands")
			brands.GET("", middleware.RequireAnyPermission("brand:read", "brand:write"), handler.ListBrands(db))
			brands.POST("", middleware.RequirePermission("brand:write"), handler.CreateBrand(db))
			brands.PUT("/:id", middleware.RequirePermission("brand:write"), handler.UpdateBrand(db))
			brands.DELETE("/:id", middleware.RequirePermission("brand:delete"), handler.DeleteBrand(db))

			products := authGroup.Group("/products")
			products.GET("", middleware.RequireAnyPermission("product:read", "product:write"), handler.ListProducts(db))
			products.GET("/:id", middleware.RequireAnyPermission("product:read", "product:write"), handler.GetProduct(db))
			products.POST("", middleware.RequirePermission("product:write"), handler.CreateProduct(db))
			products.PUT("/:id", middleware.RequirePermission("product:write"), handler.UpdateProduct(db))
			products.DELETE("/:id", middleware.RequirePermission("product:delete"), handler.DeleteProduct(db))

			banners := authGroup.Group("/banners")
			banners.GET("", middleware.RequireAnyPermission("banner:read", "banner:write"), handler.ListBanners(db))
			banners.POST("", middleware.RequirePermission("banner:write"), handler.CreateBanner(db))
			banners.PUT("/:id", middleware.RequirePermission("banner:write"), handler.UpdateBanner(db))
			banners.DELETE("/:id", middleware.RequirePermission("banner:delete"), handler.DeleteBanner(db))

			orders := authGroup.Group("/orders")
			orders.GET("", middleware.RequireAnyPermission("order:read", "order:write"), handler.ListOrders(db))
			orders.GET("/:id", middleware.RequireAnyPermission("order:read", "order:write"), handler.GetOrder(db))
			orders.PATCH("/:id", middleware.RequirePermission("order:write"), handler.UpdateOrder(db))

			admins := authGroup.Group("/admins")
			admins.Use(middleware.RequirePermission("admin:manage"))
			admins.GET("", handler.ListAdmins(db))
			admins.POST("", handler.CreateAdmin(db))
			admins.PUT("/:id", handler.UpdateAdmin(db))
			admins.DELETE("/:id", handler.DeleteAdmin(db))

			roles := authGroup.Group("/roles")
			roles.Use(middleware.RequirePermission("admin:roles"))
			roles.GET("", handler.ListRoles(db))
			roles.POST("", handler.CreateRole(db))
			roles.PUT("/:id", handler.UpdateRole(db))
			roles.DELETE("/:id", handler.DeleteRole(db))
		}
	}

	return r
}
