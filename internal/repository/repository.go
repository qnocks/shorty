package repository

import (
	"shorty/internal/dto"
	"shorty/pkg/db/redis"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type Link interface {
	Save(link dto.Link) (dto.Link, error)
	FindByShortUrl(shortURL string) (dto.Link, error)
}

type Repository struct {
	Link
}

func NewRepository(redis *redis.Client) *Repository {
	return &Repository{Link: NewLinkRepository(redis)}
}
