package entity

import (
	"time"

	"github.com/matheus-gondim/go-cryptocurrencies-election/src/pb"
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

func (c *Cryptocurrency) FromCreate(dto *pb.CreateCryptocurrency) *Cryptocurrency {
	c.Name = dto.GetName()
	c.Symbol = dto.GetSymbol()

	return c
}

func (c *Cryptocurrency) ToOutput() *pb.Cryptocurrency {
	return &pb.Cryptocurrency{Id: c.ID, Name: c.Name, Symbol: c.Symbol, Like: c.Like, Dislike: c.Dislike}
}
