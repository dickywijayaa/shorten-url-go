package objects

import (
	"time"
)

type Shorten struct {
	ID            uint       `gorm:"id" json:"id"`
	URL           string     `gorm:"url" json:"url"`
	Shortcode     string     `gorm:"shortcode" json:"shortcode"`
	RedirectCount int        `gorm:"redirect_count" json:"redirect_count"`
	CreatedAt     time.Time  `gorm:"created_at" json:"created_at"`
	LastSeenDate  *time.Time `gorm:"last_seen_date" json:"last_seen_date,omitempty"`
}

type ShortenRequest struct {
	URL       string `gorm:"url" json:"url"`
	Shortcode string `gorm:"shortcode" json:"shortcode"`
}
