//go:generate go run github.com/99designs/gqlgen

package main

import (
	"bff/protocol"
	"context"
	"google.golang.org/grpc"
	"log"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

//type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

const (
	MovieService = "movie_service:9090"
)

func (r *queryResolver) ComingSoonMovies(ctx context.Context) ([]*MovieSummary, error) {
	conn, err := grpc.Dial(MovieService, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := protocol.NewMovieServiceClient(conn)
	message := &protocol.ComingSoonMoviesRequest{}
	res, err := client.ComingSoonMovies(context.TODO(), message)

	if res == nil {
		log.Print("nil ComingSoonMovies")
		return nil, nil
	}

	convIntArr := func(nums []int32) []int {
		retVal := make([]int, len(nums))

		for i, num := range nums {
			retVal[i] = int(num)
		}
		return retVal
	}

	response := make([]*MovieSummary, len(res.Summaries))
	for i, summary := range res.Summaries {
		response[i] = &MovieSummary{
			Adult:            summary.Adult,
			BackdropPath:     summary.BackdropPath,
			GenreIds:         convIntArr(summary.GenreIds),
			ID:               int(summary.Id),
			OriginalLanguage: summary.OriginalLanguage,
			OriginalTitle:    summary.OriginalTitle,
			Overview:         summary.Overview,
			PosterPath:       summary.PosterPath,
			ReleaseDate:      summary.ReleaseDate,
			Title:            summary.Title,
		}
	}

	return response, nil
}

func (r *queryResolver) HotMovies(ctx context.Context) ([]*MovieSummary, error) {
	conn, err := grpc.Dial(MovieService, grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := protocol.NewMovieServiceClient(conn)
	message := &protocol.HotMoviesRequest{}
	res, err := client.HotMovies(context.TODO(), message)

	if res == nil {
		log.Print("nil HotMovies")
		return nil, nil
	}

	convIntArr := func(nums []int32) []int {
		retVal := make([]int, len(nums))

		for i, num := range nums {
			retVal[i] = int(num)
		}
		return retVal
	}

	response := make([]*MovieSummary, len(res.Summaries))
	for i, summary := range res.Summaries {
		response[i] = &MovieSummary{
			Adult:            summary.Adult,
			BackdropPath:     summary.BackdropPath,
			GenreIds:         convIntArr(summary.GenreIds),
			ID:               int(summary.Id),
			OriginalLanguage: summary.OriginalLanguage,
			OriginalTitle:    summary.OriginalTitle,
			Overview:         summary.Overview,
			PosterPath:       summary.PosterPath,
			ReleaseDate:      summary.ReleaseDate,
			Title:            summary.Title,
		}
	}

	return response, nil
}
