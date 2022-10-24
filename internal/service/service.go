package service

import (
	"shorty/internal/dto"
	"shorty/internal/repository"
)

type Link interface {
	CreateShortLink(originURL string) (dto.Link, error)
	GetByShortURL(shortURL string) (dto.Link, error)
	ProcessURLClick(shortURL string) (dto.Link, error)
}

type Shortener interface {
	Shorten(str string) string
}

type Service struct {
	Link
	Shortener
}

func NewService(repos *repository.Repository) *Service {
	services := &Service{
		Shortener: NewShortenerService(),
	}
	services.Link = NewLinkService(repos.Link, services.Shortener)

	return services
}
