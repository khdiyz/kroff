package handler

import (
	"errors"
	"kroff/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Description Upload file
// @Summary Upload file
// @Tags Storage
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 201 {object} response.IdResponse
// @Failure 400,401,404,500 {object} response.BaseResponse
// @Router /api/v1/admin/files [post]
// @Security ApiKeyAuth
func (h *Handler) uploadFile(c *gin.Context) {
	// Get the file from form data
	file, err := c.FormFile("file")
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, errors.New("invalid or missing file"))
		return
	}

	// Validate file size (e.g., 10MB limit)
	const maxSize = 10 << 20 // 10 MB
	if file.Size > maxSize {
		response.ErrorResponse(c, http.StatusBadRequest, errors.New("file size exceeds maximum limit of 10MB"))
		return
	}

	// Validate content type
	contentType := file.Header.Get("Content-Type")
	if !isAllowedContentType(contentType) {
		response.ErrorResponse(c, http.StatusBadRequest, errors.New("unsupported file type"))
		return
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, errors.New("failed to open uploaded file"))
		return
	}
	defer src.Close()

	// Upload file to storage service
	fileName, err := h.services.File.UploadFile(c.Request.Context(), src, file.Size, contentType)
	if err != nil {
		response.FromError(c, response.ServiceError(err, http.StatusInternalServerError))
		return
	}

	// Return success response with file ID
	c.JSON(http.StatusCreated, response.IdResponse{
		Id: fileName,
	})
}

func isAllowedContentType(contentType string) bool {
	allowedTypes := map[string]bool{
		"image/jpeg":               true,
		"image/png":                true,
		"image/gif":                true,
		"image/webp":               true,
		"application/pdf":          true,
		"image/svg+xml":            true,
		"image/x-icon":             true,
		"image/vnd.microsoft.icon": true,
	}
	return allowedTypes[contentType]
}
