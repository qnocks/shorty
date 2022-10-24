package repository

import (
	"encoding/json"
	"shorty/internal/dto"
	"shorty/pkg/db/redis"
)

type LinkRepository struct {
	redisClient *redis.Client
}

func NewLinkRepository(redisClient *redis.Client) *LinkRepository {
	return &LinkRepository{redisClient: redisClient}
}

func (r *LinkRepository) Save(link dto.Link) (dto.Link, error) {
	bytes, err := json.Marshal(link)
	if err != nil {
		return link, err
	}

	_, err = r.redisClient.Redis.Set(r.redisClient.Ctx, link.ShortURL, bytes, 0).Result()
	if err != nil {
		return link, err
	}

	return link, nil
}

func (r *LinkRepository) FindByShortUrl(shortURL string) (dto.Link, error) {
	var link dto.Link
	bytes := r.redisClient.Redis.Get(r.redisClient.Ctx, shortURL).Val()
	err := json.Unmarshal([]byte(bytes), &link)
	if err != nil {
		return link, err
	}

	return link, nil
}
