package use_case

import (
	"movie_service/internal/app/domain/models"
	"movie_service/internal/app/interface/read"
)

type UseCase interface {
	ListHotMovies() []models.Summary
	ListComingSoonMovies() []models.Summary
}

// 一応UseCaseを抽象化してInterActorを作っておく
type interActor struct {
	query read.MovieQuery
}

func NewInterActor(query read.MovieQuery) UseCase {
	return interActor{query: query}
}

func (useCase interActor) ListHotMovies() []models.Summary {
	return useCase.query.FetchHotMovies()
}

func (useCase interActor) ListComingSoonMovies() []models.Summary {
	return useCase.query.FetchComingSoonMovies()
}
