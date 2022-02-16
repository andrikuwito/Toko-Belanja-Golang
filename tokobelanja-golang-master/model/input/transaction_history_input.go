package input

type InputTransaction struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type InputUserTransaction struct {
	ID int `json:"id_user"`
}
