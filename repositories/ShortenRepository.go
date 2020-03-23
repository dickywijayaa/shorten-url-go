package repositories

import (
	"../database"
	"../objects"

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

func (r *ShortenRepository) GetURLFromCode(code string) (string, error) {
	var data objects.Shorten
	query := r.DB.Table("shorten").Select("url").Where("shortcode=?", code).First(&data)

	return data.URL, query.Error
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