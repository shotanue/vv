// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package main

type MovieSummary struct {
	Adult            bool   `json:"adult"`
	BackdropPath     string `json:"backdrop_path"`
	GenreIds         []int  `json:"genre_ids"`
	ID               int    `json:"id"`
	OriginalLanguage string `json:"original_language"`
	OriginalTitle    string `json:"original_title"`
	Overview         string `json:"overview"`
	PosterPath       string `json:"poster_path"`
	ReleaseDate      string `json:"release_date"`
	Title            string `json:"title"`
}
