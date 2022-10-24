package service

import (
	"shorty/internal/dto"
	"shorty/internal/repository"
)

type LinkService struct {
	repo      repository.Link
	shortener Shortener
}

func NewLinkService(repo repository.Link, shortener Shortener) *LinkService {
	return &LinkService{
		repo:      repo,
		shortener: shortener,
	}
}

func (s *LinkService) CreateShortLink(originURL string) (dto.Link, error) {
	link, err := s.repo.Save(dto.Link{
		OriginURL: originURL,
		ShortURL:  s.shortener.Shorten(originURL),
	})
	if err != nil {
		return link, err
	}

	return link, nil
}

func (s *LinkService) GetByShortURL(shortURL string) (dto.Link, error) {
	link, err := s.repo.FindByShortUrl(shortURL)
	if err != nil {
		return link, err
	}

	return link, nil
}

func (s *LinkService) ProcessURLClick(shortURL string) (dto.Link, error) {
	link, err := s.repo.FindByShortUrl(shortURL)
	if err != nil {
		return link, err
	}

	link.RedirectCount++

	_, err = s.repo.Save(link)
	if err != nil {
		return link, err
	}

	return link, nil
}
