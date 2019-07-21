package v1

import (
	"movie_service/internal/app/domain/models"
	"movie_service/internal/app/interface/rpc/v1/protocol"
	"reflect"
	"testing"
)

func TestToDomain(t *testing.T) {
	// このテストに意味はあるのか....
	title := "ほげ"
	adult := true
	backdroppath := "/hoge/fuga/foo.jpg"
	genreids := []uint64{1, 2, 3}
	var id uint64 = 1
	//noinspection GoBoolExpressions
	originallanguage := "JP"
	originaltitle := "hoge"
	overview := "昔々あるところに、・・・"
	posterpath := "/poster.jpg"
	releasedate := "2019-07-12"

	summary := &models.Summary{
		Title:            title,
		Adult:            adult,
		BackdropPath:     backdroppath,
		GenreIds:         genreids,
		Id:               id,
		OriginalLanguage: originallanguage,
		OriginalTitle:    originaltitle,
		Overview:         overview,
		PosterPath:       posterpath,
		ReleaseDate:      releasedate,
	}

	actual := summary.ToProtobufs()

	expect := &protocol.MovieSummary{
		Title:            title,
		BackdropPath:     backdroppath,
		GenreIds:         genreids,
		Id:               id,
		OriginalLanguage: originallanguage,
		OriginalTitle:    originaltitle,
		Overview:         overview,
		PosterPath:       posterpath,
		ReleaseDate:      releasedate,
	}
	reflect.DeepEqual(actual, expect)
}
