package v1

import (
	"context"
	"movie_service/internal/app/domain/models"
	"movie_service/internal/app/interface/rpc/v1/protocol"
	"movie_service/internal/app/use_case"
)

type MovieService struct {
	useCase use_case.UseCase
}

func NewMovieService(useCase use_case.UseCase) MovieService {
	return MovieService{
		useCase: useCase,
	}
}

func (s MovieService) HotMovies(context.Context, *protocol.HotMoviesRequest) (*protocol.HotMoviesResponse, error) {
	return &protocol.HotMoviesResponse{
		Summaries: summariesToProtobufs(s.useCase.ListHotMovies()),
	}, nil
}

func (s MovieService) ComingSoonMovies(context.Context, *protocol.ComingSoonMoviesRequest) (*protocol.ComingSoonMoviesResponse, error) {
	return &protocol.ComingSoonMoviesResponse{
		Summaries: summariesToProtobufs(s.useCase.ListComingSoonMovies()),
	}, nil
}

func (MovieService) MovieDetail(context.Context, *protocol.MovieDetailRequest) (*protocol.MovieDetailResponse, error) {
	panic("implement me")
}

func summariesToProtobufs(summaries []models.Summary) []*protocol.MovieSummary {
	movieSummary := make([]*protocol.MovieSummary, len(summaries))
	for i, summary := range summaries {
		movieSummary[i] = summary.ToProtobufs()
	}
	return movieSummary
}
