package models

import (
	"movie_service/internal/app/interface/rpc/v1/protocol"
)

type Summary struct {
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	GenreIds         []uint64 `json:"genre_ids"`
	Id               uint64   `json:"id"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    string   `json:"original_title"`
	Overview         string   `json:"overview"`
	PosterPath       string   `json:"poster_path"`
	ReleaseDate      string   `json:"release_date"`
	Title            string   `json:"title"`
}

func (summary Summary)ToProtobufs() *protocol.MovieSummary{
	return &protocol.MovieSummary{
		Title:            summary.Title,
		BackdropPath:     summary.BackdropPath,
		GenreIds:         summary.GenreIds,
		Id:               summary.Id,
		OriginalLanguage: summary.OriginalLanguage,
		OriginalTitle:    summary.OriginalTitle,
		Overview:         summary.Overview,
		PosterPath:       summary.PosterPath,
		ReleaseDate:      summary.ReleaseDate,
	}
}
