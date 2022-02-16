package controller

import (
	"fmt"
	"net/http"
	"tokobelanja-golang/helper"
	"tokobelanja-golang/middleware"
	"tokobelanja-golang/model/input"
	"tokobelanja-golang/model/response"
	"tokobelanja-golang/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input input.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helper.APIResponse("error", error_message)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	newUser, err := h.userService.CreateUser(input)

	if err != nil {
		error_message := gin.H{
			"error": err.Error(),
		}

		resp := helper.APIResponse("error", error_message)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// success created user
	userResponse := response.UserRegisterResponse{
		ID:        newUser.ID,
		FullName:  newUser.FullName,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
	}

	resp := helper.APIResponse("success", userResponse)
	c.JSON(http.StatusCreated, resp)

}

func (h *userController) Login(c *gin.Context) {
	var input input.LoginUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send to services
	// get user by email
	user, err := h.userService.GetUserByEmail(input.Email)

	if err != nil {
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// return when user not found!
	if user.ID == 0 {
		errorMessages := "User not found!"
		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusNotFound, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		response := helper.APIResponse("failed", "password not match!")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// lets create token!
	jwtService := middleware.NewService()
	token, err := jwtService.GenerateToken(user.ID)

	// return the token!
	response := helper.APIResponse("ok", gin.H{
		"token": token,
	})
	c.JSON(http.StatusOK, response)
	return
}

func (h *userController) TopUpSaldo(c *gin.Context) {
	var input input.TopUpSaldoInput
	currentUser := c.MustGet("currentUser").(int)

	err := c.ShouldBindJSON(&input)

	if err != nil {
		// errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userTopup, err := h.userService.TopUp(input.Balance, currentUser)

	if err != nil {
		// errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("ok", gin.H{
		"message": fmt.Sprintf("Your balance has been successfully updated to Rp. %d", userTopup.Balance),
	})
	c.JSON(http.StatusOK, response)
	return

}

func (h *userController) UpdateUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var inputUserUpdate input.UpdateUserInput

	err := c.ShouldBindJSON(&inputUserUpdate)

	if err != nil {
		// errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.userService.UpdateUser(currentUser, inputUserUpdate)

	if err != nil {
		// errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userUpdated, err := h.userService.GetUserByID(currentUser)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userResponse := response.UserUpdateResponse{
		ID:        userUpdated.ID,
		FullName:  userUpdated.FullName,
		Email:     userUpdated.Email,
		UpdatedAt: userUpdated.UpdatedAt,
	}

	response := helper.APIResponse("ok", userResponse)
	c.JSON(http.StatusOK, response)
	return
}

func (h *userController) DeleteUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	_, err := h.userService.DeleteUser(currentUser)

	if err != nil {
		// errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("ok", "Your account has been successfully deleted!")
	c.JSON(http.StatusOK, response)
	return
}
