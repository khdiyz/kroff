package handler

import (
	"kroff/pkg/models"
	"kroff/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create category
// @Description Create a new category
// @Tags Category
// @Accept json
// @Produce json
// @Param category body models.CreateCategory true "Category details"
// @Success 201 {object} response.IdResponse
// @Failure 400,500 {object} response.BaseResponse
// @Router /api/v1/admin/categories [post]
// @Security ApiKeyAuth
func (h *Handler) createCategory(c *gin.Context) {
	var input models.CreateCategory
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	id, err := h.services.Category.CreateCategory(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.IdResponse{Id: strconv.FormatInt(id, 10)})
}

// @Summary Get all categories
// @Description Get all categories
// @Tags Category
// @Produce json
// @Success 200 {object} []models.Category
// @Failure 500 {object} response.BaseResponse
// @Router /api/v1/admin/categories [get]
// @Security ApiKeyAuth
func (h *Handler) getAllCategories(c *gin.Context) {
	categories, err := h.services.Category.GetAllCategories()
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, categories)
}

// @Summary Get category by ID
// @Description Get category by ID
// @Tags Category
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} models.Category
// @Failure 400,404,500 {object} response.BaseResponse
// @Router /api/v1/admin/categories/{id} [get]
// @Security ApiKeyAuth
func (h *Handler) getCategoryByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	category, err := h.services.Category.GetCategoryByID(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, category)
}

// @Summary Update category
// @Description Update category by ID
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param category body models.UpdateCategory true "Category details"
// @Success 200 {object} response.BaseResponse
// @Failure 400,404,500 {object} response.BaseResponse
// @Router /api/v1/admin/categories/{id} [put]
// @Security ApiKeyAuth
func (h *Handler) updateCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var input models.UpdateCategory
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	input.ID = id

	err = h.services.Category.UpdateCategory(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{Message: "Category updated successfully"})
}

// @Summary Delete category
// @Description Delete category by ID
// @Tags Category
// @Param id path string true "Category ID"
// @Success 200 {object} response.BaseResponse
// @Failure 400,404,500 {object} response.BaseResponse
// @Router /api/v1/admin/categories/{id} [delete]
// @Security ApiKeyAuth
func (h *Handler) deleteCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	err = h.services.Category.DeleteCategory(id)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{Message: "Category deleted successfully"})
}

// @Summary Get all categories public
// @Description Get all categories public
// @Tags Category
// @Produce json
// @Param Accept-Language header string false "Accept-Language" Enums(uz, ru)
// @Success 200 {object} []models.CategoryPublic
// @Failure 400,500 {object} response.BaseResponse
// @Router /api/v1/categories [get]
func (h *Handler) getCategoriesPublic(c *gin.Context) {
	categories, err := h.services.Category.GetAllCategoriesPublic(c.GetHeader("Accept-Language"))
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, categories)
}
