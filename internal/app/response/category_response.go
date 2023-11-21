package response

import "time"

type CategoryResponse struct {
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ID          int       `json:"id"`
}
