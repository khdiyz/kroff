package handler

import (
	"kroff/pkg/models"
	"kroff/utils/helper"
	"kroff/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.CreateProduct true "Product details"
// @Success 201 {object} response.IdResponse
// @Failure 400,500 {object} response.BaseResponse
// @Router /api/v1/admin/products [post]
// @Security ApiKeyAuth
func (h *Handler) createProduct(c *gin.Context) {
	var input models.CreateProduct
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	id, err := h.services.Product.CreateProduct(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.IdResponse{Id: strconv.FormatInt(id, 10)})
}

// @Summary Get all products
// @Description Get all products
// @Tags products
// @Produce json
// @Success 200 {object} []models.Product
// @Failure 400,500 {object} response.BaseResponse
// @Router /api/v1/admin/products [get]
// @Security ApiKeyAuth
func (h *Handler) getProducts(c *gin.Context) {
	products, err := h.services.Product.GetAllProducts()
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

// @Summary Get product by ID
// @Description Get product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 400,404,500 {object} response.BaseResponse
// @Router /api/v1/admin/products/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	product, err := h.services.Product.GetProductByID(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary Update product
// @Description Update product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body models.UpdateProduct true "Product details"
// @Success 200 {object} response.BaseResponse
// @Failure 400,404,500 {object} response.BaseResponse
// @Router /api/v1/admin/products/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var input models.UpdateProduct
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	input.ID = id

	err = h.services.Product.UpdateProduct(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{Message: "Product updated successfully"})
}

// @Summary Delete product
// @Description Delete product by ID
// @Tags products
// @Param id path string true "Product ID"
// @Success 200 {object} response.BaseResponse
// @Failure 400,404,500 {object} response.BaseResponse
// @Router /api/v1/admin/products/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	err = h.services.Product.DeleteProduct(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{Message: "Product deleted successfully"})
}

type getProductsPublicResponse struct {
	Data       []models.ProductPublic `json:"data"`
	Pagination models.Pagination      `json:"pagination"`
}

// @Summary Get all products public
// @Description Get all products public
// @Tags products
// @Produce json
// @Param categoryId query string false "Category ID"
// @Param page query int true "page" default(1)
// @Param limit query int true "page limit" default(10)
// @Param Accept-Language header string false "Accept-Language" Enums(uz, ru)
// @Success 200 {object} getProductsPublicResponse
// @Failure 400,500 {object} response.BaseResponse
// @Router /api/v1/products [get]
func (h *Handler) getProductsPublic(c *gin.Context) {
	pagination, err := helper.ListPagination(c)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	filters := make(map[string]any)

	categoryId := c.Query("categoryId")
	if categoryId != "" {
		categoryIdInt, err := strconv.ParseInt(categoryId, 10, 64)
		if err != nil {
			response.ErrorResponse(c, http.StatusBadRequest, err)
			return
		}
		filters["category_id"] = categoryIdInt
	}

	options := models.FilterOptions{
		Filters: filters,
		Limit:   pagination.Limit,
		Page:    pagination.Page,
	}

	products, totalCount, err := h.services.Product.GetAllProductsPublic(c.GetHeader("Accept-Language"), options)
	if err != nil {
		response.FromError(c, err)
		return
	}

	helper.UpdatePagination(&pagination, totalCount)

	c.JSON(http.StatusOK, getProductsPublicResponse{
		Data:       products,
		Pagination: pagination,
	})
}
