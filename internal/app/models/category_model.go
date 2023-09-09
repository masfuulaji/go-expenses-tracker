package models

import (
	"time"

	"github.com/lib/pq"
)

type Category struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    CreatedAt time.Time `db:"created_at" json:"created_at"`
    UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
    DeletedAt pq.NullTime `db:"deleted_at" json:"deleted_at"`
}
