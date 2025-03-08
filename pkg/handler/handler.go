package handler

import (
	"kroff/docs"
	"kroff/pkg/service"
	"kroff/utils/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Services
	logger   *logger.Logger
}

func NewHandlers(services *service.Services, logger *logger.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.HandleMethodNotAllowed = true
	router.Use(corsMiddleware())

	h.setupSwagger(router)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/auth/login", h.login)

		admin := v1.Group("/admin" /*h.adminIdentity()*/)
		{
			category := admin.Group("/categories")
			{
				category.POST("", h.createCategory)
				category.GET("", h.getAllCategories)
				category.GET("/:id", h.getCategoryByID)
				category.PUT("/:id", h.updateCategory)
				category.DELETE("/:id", h.deleteCategory)
			}

			product := admin.Group("/products")
			{
				product.POST("", h.createProduct)
				product.GET("", h.getProducts)
				product.GET("/:id", h.getProduct)
				product.PUT("/:id", h.updateProduct)
				product.DELETE("/:id", h.deleteProduct)
			}

			admin.POST("/files", h.uploadFile)
		}

		category := v1.Group("/categories")
		{
			category.GET("", h.getCategoriesPublic)
		}

		product := v1.Group("/products")
		{
			product.GET("", h.getProductsPublic)
		}
	}

	return router
}

func (h *Handler) setupSwagger(router *gin.Engine) {
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler), func(ctx *gin.Context) {
		docs.SwaggerInfo.Host = ctx.Request.Host
		if ctx.Request.TLS != nil {
			docs.SwaggerInfo.Schemes = []string{"https"}
		}
	})
}
