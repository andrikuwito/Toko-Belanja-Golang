package input

type CreateProduct struct {
	Title      string `json:"title" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
}

type UpdateProduct struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required"`
	Stock int    `json:"stock" binding:"required"`
}

type DeleteProduct struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required"`
	Stock int    `json:"stock" binding:"required"`
}

type IDProduct struct {
	ID int `uri:"id" binding:"required"`
}

type InputProduct struct {
	ID int `json:"id_product"`
}
