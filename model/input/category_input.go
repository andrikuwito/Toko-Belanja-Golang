package input

type CreateCategory struct {
	Type string `json:"type" binding:"required"`
}

type UpdateCategory struct {
	Type string `json:"type" binding:"required"`
}

type IDCategory struct {
	ID int `uri:"id" binding:"required"`
}

type DeleteCategory struct {
	ID int `uri:"id" binding:"required"`
}
