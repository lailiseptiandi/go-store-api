package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/go-store-api/entity"
	"github.com/lailiseptiandi/go-store-api/helpers"
)

func (h *globalHandler) GetProduct(c *gin.Context) {
	product, err := h.globalService.GetProduct()
	if err != nil {
		response := helpers.APIResponse("product not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.APIResponse("list product", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
	return
}

func (h *globalHandler) CreateProduct(c *gin.Context) {
	var input entity.CreateProduct

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helpers.APIResponse("failed created product", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := h.globalService.CreateProduct(input)
	if err != nil {
		response := helpers.APIResponse("failed created product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("success created prodct", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
	return
}

func (h *globalHandler) UpdateProduct(c *gin.Context) {
	var input entity.ProductDetail
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helpers.APIResponse("product not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputProduct entity.CreateProduct
	err = c.ShouldBindJSON(&inputProduct)
	if err != nil {
		errors := helpers.FormValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helpers.APIResponse("failed updated product", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	product, err := h.globalService.UpdateProduct(input.ID, inputProduct)

	if err != nil {
		response := helpers.APIResponse("failed updated product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("success updated product", http.StatusBadRequest, "error", product)
	c.JSON(http.StatusBadRequest, response)
	return

}

func (h *globalHandler) ProductByID(c *gin.Context) {
	var inputID entity.ProductDetail
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helpers.FormValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helpers.APIResponse("product not found", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	product, err := h.globalService.ProductByID(inputID.ID)
	if err != nil {
		response := helpers.APIResponse("product not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}
	response := helpers.APIResponse("success get detail product", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
	return

}

func (h *globalHandler) DeleteProduct(c *gin.Context) {
	var input entity.ProductDetail
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helpers.APIResponse("product not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = h.globalService.DeleteProduct(input.ID)
	if err != nil {
		response := helpers.APIResponse("failed updated product", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.APIResponse("success deleted product", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
	return
}
