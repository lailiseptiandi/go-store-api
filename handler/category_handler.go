package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/golang-toko-online/entity"
	"github.com/lailiseptiandi/golang-toko-online/helpers"
	"github.com/lailiseptiandi/golang-toko-online/service"
)

func (h *globalHandler) GetCategory(c *gin.Context) {
	getCategory, err := h.globalService.GetCategory()
	if err != nil {
		response := helpers.APIResponse("category not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := service.FormatterCategory(getCategory)
	response := helpers.APIResponse("List Category", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *globalHandler) CreateCategory(c *gin.Context) {
	var input entity.CreateCategoryInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helpers.FormValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helpers.APIResponse("Create category failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newCategory, err := h.globalService.CreateCategory(input)
	if err != nil {
		response := helpers.APIResponse("Create category failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := service.FormatCategory(newCategory)
	response := helpers.APIResponse("Create category successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *globalHandler) DetailCategory(c *gin.Context) {
	var inputID entity.CategoryDetailInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		errors := helpers.FormValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helpers.APIResponse("Category Not Found", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	detailCategory, err := h.globalService.CategoryByID(inputID.ID)
	if err != nil {
		response := helpers.APIResponse("Category Not Found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := service.FormatCategory(detailCategory)
	response := helpers.APIResponse("Detail Category", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *globalHandler) UpdateCategory(c *gin.Context) {
	var inputID entity.CategoryDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helpers.APIResponse("Category Not Found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	category, err := h.globalService.CategoryByID(inputID.ID)

	if err != nil {
		response := helpers.APIResponse("Category Not Found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputCategory entity.CreateCategoryInput
	inputCategory.Name = category.Name
	inputCategory.Slug = category.Name

	err = c.ShouldBindJSON(&inputCategory)
	if err != nil {
		errors := helpers.FormValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helpers.APIResponse("Update Category Failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updateCategory, err := h.globalService.UpdateCategory(inputID, inputCategory)
	if err != nil {
		response := helpers.APIResponse("Update Category Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := service.FormatCategory(updateCategory)
	response := helpers.APIResponse("Update Category successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *globalHandler) DeleteGategory(c *gin.Context) {
	var inputID entity.CategoryDetailInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helpers.APIResponse("Category Not Found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	deleteCategory, _ := h.globalService.DeleteCategory(inputID.ID)

	formatter := service.FormatCategory(deleteCategory)
	response := helpers.APIResponse("Category Deleted", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
