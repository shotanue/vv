package read

import "movie_service/internal/app/domain/models"

type MovieQuery interface {

	// 人気の映画のリストを返す
	// 何を持って人気かはtmdbの仕様になるので、おそらく外国での興行やtmdbのサイトでの独自注目度だったりするかもしれない
	FetchHotMovies() []models.Summary

	// 近日公開の映画のリストを返す
	// 日本の映画のみだと3件だったりして寂しい
	FetchComingSoonMovies() []models.Summary
}
