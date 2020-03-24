package objects

import (
	"time"
)

type Shorten struct {
	ID			uint			`gorm:"id" json:"id"`
	URL 		string			`gorm:"url" json:"url"`
	Shortcode	string			`gorm:"shortcode" json:"shortcode"`
	CreatedAt	time.Time		`gorm:"created_at" json:"created_at"`
}

type ShortenRequest struct {
	URL			string 	`gorm:"url" json:"url"`
	Shortcode	string	`gorm:"shortcode" json:"shortcode"`
}