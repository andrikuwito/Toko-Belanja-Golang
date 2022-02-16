package controller

import (
	"net/http"
	"tokobelanja-golang/helper"
	"tokobelanja-golang/model/input"
	"tokobelanja-golang/model/response"
	"tokobelanja-golang/service"

	"github.com/gin-gonic/gin"
)

type transactionHistoryController struct {
	transactionHistoryService service.TransactionHistoryService
	userService               service.UserService
	productService            service.ProductService
}

func NewTransactionHistoryController(transactionHistoryService service.TransactionHistoryService, userService service.UserService, productService service.ProductService) *transactionHistoryController {
	return &transactionHistoryController{transactionHistoryService, userService, productService}
}

func (h *transactionHistoryController) NewTransaction(c *gin.Context) {
	var input input.InputTransaction

	err := c.ShouldBindJSON(&input)

	currentUser := c.MustGet("currentUser").(int)

	if err != nil {
		resp := helper.APIResponse("error", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	newTransaction, err := h.transactionHistoryService.CreateTransaction(input, currentUser)

	if err != nil {
		resp := helper.APIResponse("error", err.Error())
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	product, err := h.productService.GetProductByID(input.ProductID)

	billResponse := response.NewTransactionBillResponse{
		TotalPrice:   newTransaction.TotalPrice,
		Quantity:     input.Quantity,
		ProductTitle: product.Title,
	}

	newTransactionResponse := response.NewTransactionResponse{
		Message:         "You have successfully purchased the product",
		TransactionBill: billResponse,
	}

	resp := helper.APIResponse("success", newTransactionResponse)
	c.JSON(http.StatusCreated, resp)
}

func (h *transactionHistoryController) GetMyTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	err := c.ShouldBind(&currentUser)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transactions, err := h.transactionHistoryService.GetMyTransaction(currentUser)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("success", transactions)
	c.JSON(http.StatusOK, response)
	return
}

func (h *transactionHistoryController) GetUserTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	err := c.ShouldBind(&currentUser)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transactions, err := h.transactionHistoryService.GetUserTransaction(currentUser)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("success", transactions)
	c.JSON(http.StatusOK, response)
	return

}
