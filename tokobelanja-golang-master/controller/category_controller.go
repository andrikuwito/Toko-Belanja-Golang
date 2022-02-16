package controller

import (
	"net/http"
	"tokobelanja-golang/helper"
	"tokobelanja-golang/model/input"
	"tokobelanja-golang/model/response"
	"tokobelanja-golang/service"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryService service.CategoryService
	userService     service.UserService
}

func NewCategoryController(categoryService service.CategoryService, userService service.UserService) *categoryController {
	return &categoryController{categoryService, userService}
}

func (h *categoryController) CreateCategory(c *gin.Context) {
	var input input.CreateCategory

	err := c.ShouldBindJSON(&input)

	if err != nil {

		resp := helper.APIResponse("error", err.Error())

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	currentUser := c.MustGet("currentUser").(int)

	userResult, err := h.userService.GetUserByID(currentUser)

	if userResult.Role != "admin" {
		resp := helper.APIResponse("error", "Unauthorized User!")
		c.JSON(http.StatusUnauthorized, resp)
		return
	}

	newCategory, err := h.categoryService.CreateCategory(input)

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helper.APIResponse("error", error_message)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// success created user
	Response := response.CreateCategoryResponse{
		ID:        newCategory.ID,
		Type:      newCategory.Type,
		CreatedAt: newCategory.CreatedAt,
	}

	resp := helper.APIResponse("success", Response)
	c.JSON(http.StatusCreated, resp)

}

func (h *categoryController) UpdateCategory(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var inputUpdate input.UpdateCategory
	var IDInput input.IDCategory

	err := c.ShouldBindJSON(&inputUpdate)
	err = c.ShouldBindUri(&IDInput)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userResult, err := h.userService.GetUserByID(currentUser)

	if userResult.Role != "admin" {
		resp := helper.APIResponse("error", "Unauthorized User!")
		c.JSON(http.StatusUnauthorized, resp)
		return
	}

	_, err = h.categoryService.UpdateCategory(IDInput.ID, inputUpdate)

	if err != nil {
		// errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	categoryUpdated, err := h.categoryService.GetCategoryByID(currentUser)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userResponse := response.UpdateCategoryResponse{
		ID:        categoryUpdated.ID,
		Type:      categoryUpdated.Type,
		UpdatedAt: categoryUpdated.UpdatedAt,
	}

	response := helper.APIResponse("ok", userResponse)
	c.JSON(http.StatusOK, response)
	return
}

func (h *categoryController) DeleteCategory(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	var Deleted input.DeleteCategory
	err := c.ShouldBindUri(&Deleted)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userResult, err := h.userService.GetUserByID(currentUser)

	if userResult.Role != "admin" {
		resp := helper.APIResponse("error", "Unauthorized User!")
		c.JSON(http.StatusUnauthorized, resp)
		return
	}

	_, err = h.categoryService.DeleteCategory(Deleted.ID)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	message := "Category deleted"

	response := helper.APIResponse("ok", message)
	c.JSON(http.StatusOK, response)
}

func (h *categoryController) GetAllCategory(c *gin.Context) {

	allCategory, err := h.categoryService.GetCategory()

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helper.APIResponse("error", error_message)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helper.APIResponse("success", allCategory)
	c.JSON(http.StatusOK, resp)
}
