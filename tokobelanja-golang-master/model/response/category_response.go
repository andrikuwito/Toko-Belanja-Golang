package response

import "time"

type CreateCategoryResponse struct {
	ID        int       `json:"id"`
	Type      string    `json:"types"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateCategoryResponse struct {
	ID        int       `json:"id"`
	Type      string    `json:"types"`
	UpdatedAt time.Time `json:"updated_at"`
}
