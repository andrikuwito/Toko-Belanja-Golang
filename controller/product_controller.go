package controller

import (
	"net/http"
	"tokobelanja-golang/helper"
	"tokobelanja-golang/model/input"
	"tokobelanja-golang/model/response"
	"tokobelanja-golang/service"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService service.ProductService
	userService    service.UserService
}

func NewProductController(productService service.ProductService, userService service.UserService) *productController {
	return &productController{productService, userService}
}

func (h *productController) CreateProduct(c *gin.Context) {
	var input input.CreateProduct

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

	newProduct, err := h.productService.CreateProduct(input)

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helper.APIResponse("error", error_message)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// success created user
	Response := response.CreateProductResponse{
		ID:        newProduct.ID,
		Title:     newProduct.Title,
		Price:     newProduct.Price,
		Stock:     newProduct.Stock,
		CreatedAt: newProduct.CreatedAt,
	}

	resp := helper.APIResponse("success", Response)
	c.JSON(http.StatusCreated, resp)

}

func (h *productController) UpdateProduct(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var inputUpdate input.UpdateProduct
	var IDInput input.IDProduct

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

	_, err = h.productService.UpdateProduct(IDInput.ID, inputUpdate)

	if err != nil {
		// errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	productUpdated, err := h.productService.GetProductByID(IDInput.ID)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	productResponse := response.UpdateProductResponse{
		ID:        productUpdated.ID,
		Title:     productUpdated.Title,
		Price:     productUpdated.Price,
		Stock:     productUpdated.Stock,
		UpdatedAt: productUpdated.UpdatedAt,
	}

	response := helper.APIResponse("ok", productResponse)
	c.JSON(http.StatusOK, response)
	return
}

func (h *productController) DeleteProduct(c *gin.Context) {

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

	_, err = h.productService.DeleteProduct(Deleted.ID)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	message := "products deleted"

	response := helper.APIResponse("ok", message)
	c.JSON(http.StatusOK, response)
}

func (h *productController) GetAllProduct(c *gin.Context) {

	allProduct, err := h.productService.GetProduct()

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helper.APIResponse("error", error_message)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := helper.APIResponse("success", allProduct)
	c.JSON(http.StatusOK, resp)
}

func (h *productController) GetProductByID(c *gin.Context) {

	var idproduct input.InputProduct

	product, err := h.productService.GetProductByID(idproduct.ID)

	if err != nil {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	response := helper.APIResponse("ok", product)
	c.JSON(http.StatusOK, response)
	return
}
