package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/golang-toko-online/entity"
	"github.com/lailiseptiandi/golang-toko-online/helpers"
)

func (h *globalHandler) GetImageByProduct(c *gin.Context) {
	var input entity.ImageByProduct

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helpers.FormValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helpers.APIResponse("image product not found", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	imageProduct, err := h.globalService.GetImageByProduct(input.ID)
	if err != nil {
		response := helpers.APIResponse("image product not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}
	response := helpers.APIResponse("success get image by product", http.StatusOK, "success", imageProduct)
	c.JSON(http.StatusOK, response)
	return
}
