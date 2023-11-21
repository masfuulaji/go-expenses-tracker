package response

import "time"

type ExpenseResponse struct {
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	Description string    `json:"description"`
	CategoryId  string    `json:"category_id"`
	ID          int       `json:"id"`
	Incoming    float64   `json:"incoming"`
	Outgoing    float64   `json:"outgoing"`
	Balance     float64   `json:"balance"`
}
