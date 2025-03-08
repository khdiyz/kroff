package handler

import (
	"kroff/pkg/models"
	"kroff/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginResponse struct {
	Token string `json:"token"`
}

// @Summary Login
// @Description Login to the system
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login request"
// @Success 200 {object} loginResponse
// @Failure 400,401,500 {object} response.BaseResponse
// @Router /api/v1/auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var input models.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	token, err := h.services.Authorization.Login(input)
	if err != nil {
		response.FromError(c, err)
		return
	}

	c.JSON(http.StatusOK, loginResponse{Token: token})
}
