package input

type RegisterUserInput struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type TopUpSaldoInput struct {
	Balance int `json:"balance"`
}
