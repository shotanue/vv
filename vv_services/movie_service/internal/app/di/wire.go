//+build wireinject

package di

// The build tag makes sure the stub is not built in the final build.
import (
	"github.com/google/wire"
	"movie_service/internal/app/interface/read/in_memory"
	"movie_service/internal/app/interface/read/redis"
	"movie_service/internal/app/interface/read/tmdb"
	"movie_service/internal/app/interface/rpc/v1"
	"movie_service/internal/app/use_case"
)

func InitializeMovieService() (movieService v1.MovieService) {
	wire.Build(redis.NewRedisCache ,tmdb.NewQuery, use_case.NewInterActor, v1.NewMovieService)
	return
}


func InitializeTestMovieService() (movieService v1.MovieService) {
	wire.Build(in_memory.NewQuery, use_case.NewInterActor, v1.NewMovieService)
	return
}
