package services

import (
	"../repositories"
	"../objects"

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

func (s *ShortenService) StoreShortenURL(url string, code string) (bool, error) {
	data := objects.Shorten{
		URL: url,
		Shortcode: code,
		CreatedAt: time.Now(),
	}

	result, err := s.repository.StoreShortcode(data)
	return result, err
}