package read

import "movie_service/internal/app/domain/models"

type Cache interface {
	FetchSummaries(key string) ([]models.Summary, error)
	SetSummaries(key string, summaries []models.Summary)
}

// キャッシュに使用するキーの一覧
const (
	HotMovies = "hot-movies"
	ComingSoonMovies = "coming-soon-movies"
)
