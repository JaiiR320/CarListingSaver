package types

import "time"

type Listing struct {
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	Price     int       `json:"price"`
	Title     string    `json:"title"`
	Mileage   int       `json:"mileage"`
	CreatedAt time.Time `json:"created_at"`
}

func NewListing(url string, price int, title string, mileage int) *Listing {
	return &Listing{
		Url:       url,
		Price:     price,
		Title:     title,
		Mileage:   mileage,
		CreatedAt: time.Now(),
	}
}
