package repositories

import (
	"fmt"
	"time"

	"github.com/dickywijayaa/shorten-url-go/database"
	"github.com/dickywijayaa/shorten-url-go/objects"

	"github.com/jinzhu/gorm"
)

type ShortenRepository struct {
	DB *gorm.DB
}

func ShortenRepositoryHandler() ShortenRepository {
	handler := ShortenRepository{
		DB: database.GetConnection(),
	}
	return handler
}

func (r *ShortenRepository) GetDetailsFromCode(code string) (objects.Shorten, error) {
	var data objects.Shorten
	query := r.DB.Table("shorten").Select("*").Where("shortcode=?", code).First(&data)

	return data, query.Error
}

func (r *ShortenRepository) StoreShortcode(data objects.Shorten) (bool, error) {
	query := r.DB.Table("shorten").Create(&data)
	if err := query.Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *ShortenRepository) CheckCodeExists(code string) (int, error) {
	var count int
	query := r.DB.Table("shorten").Where("shortcode=?", code).Count(&count)

	if err := query.Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (r *ShortenRepository) UpdateLastSeen(data objects.Shorten) {
	current_date := time.Now()
	data.RedirectCount = data.RedirectCount + 1
	data.LastSeenDate = &current_date

	query := r.DB.Table("shorten").Where("id = ?", data.ID).Update(data)
	if err := query.Error; err != nil {
		fmt.Println(err.Error())
	}

	return
}
