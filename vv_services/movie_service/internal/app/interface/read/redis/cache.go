package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"movie_service/internal/app/domain/models"
	"movie_service/internal/app/interface/read"
	"time"
)

type Cache struct {
	redis redis.Client
}

func NewRedisCache() read.Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Cache{
		redis: *client,
	}
}

type NoCacheError struct {
}

func (*NoCacheError) Error() string {
	return "no cache found"
}

type SummaryContainer struct {
	data []models.Summary
}

func (cache Cache) FetchSummaries(key string) ([]models.Summary, error) {
	summary, err := cache.redis.Get(key).Result()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	if summary == "" {
		return nil, &NoCacheError{}
	}

	var container []models.Summary
	_ = json.Unmarshal([]byte(summary), &container)
	return container, nil
}

func (cache Cache) SetSummaries(key string, summaries []models.Summary) {
	hotMoviesJson, err := json.Marshal(summaries)
	if err != nil {
		log.Print(err)
		return
	}
	cache.redis.Set(key, hotMoviesJson, 6*time.Hour)
}
