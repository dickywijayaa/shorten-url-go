package objects

import (
	"time"

	_ "gopkg.in/go-playground/validator.v10"
)

type Shorten struct {
	ID			uint			`gorm:"id" json:"id"`
	URL 		string			`gorm:"url" json:"url"`
	Shortcode	string			`gorm:"shortcode" json:"shortcode"`
	CreatedAt	time.Time		`gorm:"created_at" json:"created_at"`
}

type ShortenRequest struct {
	URL			string 	`gorm:"url" json:"url" validate:"required"`
	Shortcode	string	`gorm:"shortcode" json:"shortcode" validate:"required"`
}