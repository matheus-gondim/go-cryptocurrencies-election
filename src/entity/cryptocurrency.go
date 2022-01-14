package entity

import (
	"time"
)

type Cryptocurrency struct {
	ID        int64     `json:"id" gorm:"primarykey"`
	Name      string    `json:"name"`
	Symbol    string    `json:"symbol"`
	Like      int64     `json:"like"`
	Dislike   int64     `json:"dislike"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
