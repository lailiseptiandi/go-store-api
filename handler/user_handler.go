package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/golang-toko-online/entity"
	"github.com/lailiseptiandi/golang-toko-online/helpers"
	"github.com/lailiseptiandi/golang-toko-online/service"
)

func (h *globalHandler) RegisterUser(c *gin.Context) {
	var input entity.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helpers.FormValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helpers.APIResponse("Register account failed", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUser, err := h.globalService.RegisterUser(input)
	if err != nil {

		response := helpers.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {

		response := helpers.APIResponse("Generate token failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	user := service.FormatUser(newUser, token)
	response := helpers.APIResponse("Register account succesfully", http.StatusOK, "error", user)

	c.JSON(http.StatusOK, response)

}

func (h *globalHandler) LoginUser(c *gin.Context) {
	var input entity.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helpers.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginUser, err := h.globalService.LoginUser(input)
	if err != nil {
		response := helpers.APIResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loginUser.ID)

	if err != nil {

		response := helpers.APIResponse("Generate token failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginFormat := service.FormatUser(loginUser, token)
	response := helpers.APIResponse("Login successfully", http.StatusOK, "success", loginFormat)
	c.JSON(http.StatusOK, response)

}

func (h *globalHandler) CheckEmailAvailable(c *gin.Context) {

	var input entity.CheckEmailUser

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helpers.APIResponse("Email checking failed", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	isEmailAvailable, err := h.globalService.CheckEmailUser(input)
	if err != nil {
		errorMessage := gin.H{"error": "server error"}
		response := helpers.APIResponse("Email checking failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	data := gin.H{"is_available": isEmailAvailable}
	var metaMessage string

	if isEmailAvailable {
		metaMessage = "Email is available"
	} else {
		metaMessage = "Email has been registered"
	}

	response := helpers.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *globalHandler) UpdateUser(c *gin.Context) {
	var inputID entity.GetUserID

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helpers.APIResponse("User not found", http.StatusBadRequest, "error", errors.New("failed"))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := h.globalService.GetUserByID(inputID.ID)

	if err != nil {
		response := helpers.APIResponse("User not found", http.StatusBadRequest, "error", errors.New("failed"))
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputUser entity.RegisterUserInput
	inputUser.Password = user.Password
	inputUser.Email = user.Email

	err = c.ShouldBindJSON(&inputUser)
	if err != nil {
		errors := helpers.FormValidationError(err)
		errorMessage := gin.H{"error": errors}
		response := helpers.APIResponse("Update user failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updateUser, err := h.globalService.UpdateUser(inputID, inputUser)

	if err != nil {
		response := helpers.APIResponse("Update user failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(updateUser.ID)
	if err != nil {

		response := helpers.APIResponse("Generate token failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formateUser := service.FormatUser(updateUser, token)
	response := helpers.APIResponse("User successfully updated", http.StatusOK, "success", formateUser)
	c.JSON(http.StatusOK, response)
}

func (h *globalHandler) GetUserByID(c *gin.Context) {
	var inputID entity.GetUserID

	err := c.ShouldBindUri(&inputID)

	if err != nil {
		response := helpers.APIResponse("User not found", http.StatusBadRequest, "error", errors.New("failed"))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := h.globalService.GetUserByID(inputID.ID)

	if err != nil {
		response := helpers.APIResponse("User not found", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatUser := service.FormatDetailUser(user)
	response := helpers.APIResponse("List User", http.StatusBadRequest, "success", formatUser)
	c.JSON(http.StatusBadRequest, response)
}
