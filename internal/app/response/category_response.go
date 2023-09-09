package response

import "time"

type CategoryResponse struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    CreatedAt time.Time `db:"created_at" json:"created_at"`
    UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
