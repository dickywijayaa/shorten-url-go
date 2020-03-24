package services

import (
	"github.com/dickywijayaa/shorten-url-go/repositories"
	"github.com/dickywijayaa/shorten-url-go/objects"

	"errors"
	"time"
)

type ShortenService struct {
	repository repositories.ShortenRepository
}

func ShortenServiceHandler() ShortenService {
	handler := ShortenService{
		repository: repositories.ShortenRepositoryHandler(),
	}

	return handler
}

func (s *ShortenService) FetchURLByCode(code string) (string, error) {
	result, err := s.repository.GetURLFromCode(code)
	return result, err
}

func (s *ShortenService) StoreShortenURL(url string, code string) (objects.Shorten, error) {
	var temp objects.Shorten
	countCheckURL, err := s.repository.CheckCodeExists(code)
	if err != nil {
		return temp, err
	}

	if countCheckURL > 0 {
		return temp, errors.New("code already exists")
	}

	data := objects.Shorten{
		URL: url,
		Shortcode: code,
		CreatedAt: time.Now(),
	}

	result, err := s.repository.StoreShortcode(data)

	if result == true {
		return data, err
	}

	return temp, err
}