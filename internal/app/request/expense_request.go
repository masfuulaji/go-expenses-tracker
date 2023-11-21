package request

import "time"

type ExpenseRequest struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	CategoryId  string    `json:"category_id"`
	Incoming    float64   `json:"incoming"`
	Outgoing    float64   `json:"outgoing"`
	Balance     float64   `json:"balance"`
}
