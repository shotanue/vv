package tmdb

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"io/ioutil"
	"log"
	"movie_service/configs"
	"movie_service/internal/app/domain/models"
	"movie_service/internal/app/interface/read"
	"movie_service/internal/app/interface/read/redis"
	"net/http"
)

type Query struct {
	cache    read.Cache
	apiToken string
}

func NewQuery(cache read.Cache) read.MovieQuery {
	conf := configs.TmdbConfig{}
	if err := envconfig.Process("", &conf); err != nil {
		log.Fatal(err)
	}

	return Query{cache, conf.TmdbToken}
}

func (query Query) FetchHotMovies() [] models.Summary {
	// キャッシュがあるかまず確認
	hotMovies, err := query.cache.FetchSummaries(read.HotMovies)
	if err == nil {
		return hotMovies
	}
	switch e := err.(type) {
	case *redis.NoCacheError:
		// キャッシュがないだけなので、それ自体は特に問題ない
	default:
		log.Print(e)
	}

	// キャッシュがなければAPIを叩いてデータを取りに行く
	// ほんとはバッチとかで定期的にキャッシュなりDBにデータをいれてあげるのが良いのだろうけど、ひとまずAPIを直接叩いて取りに行く
	url := fmt.Sprintf(
		"https://api.themoviedb.org/3/movie/popular?api_key=%s&language=ja-JP&page=1&region=JP",
		query.apiToken,
	)

	summaries := pluckMovieSummary(get(url))

	// 取ったデータはキャッシュに入れる。
	go func() {
		query.cache.SetSummaries(read.HotMovies, summaries)
	}()

	return summaries
}

func (query Query) FetchComingSoonMovies() []models.Summary {
	comingSoonMovies, err := query.cache.FetchSummaries(read.ComingSoonMovies)
	if err == nil {
		return comingSoonMovies
	}
	switch e := err.(type) {
	case *redis.NoCacheError:
		// キャッシュがないだけなので問題ない
	default:
		log.Print(e)
	}

	// APIを叩いてデータを取ってくる
	url := fmt.Sprintf(
		"https://api.themoviedb.org/3/movie/upcoming?api_key=%s&language=ja-JP&page=1&region=JP",
		query.apiToken,
	)

	summaries := pluckMovieSummary(get(url))

	// 取ったデータはキャッシュに入れる
	go func() {
		query.cache.SetSummaries(read.ComingSoonMovies, summaries)
	}()

	return summaries
}

// APIを叩くhelper
func get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return byteArray
}

// tmdbのAPIのレスポンスからメインのデータを抜き出す
func pluckMovieSummary(byteArray []byte) []models.Summary {
	type Response struct {
		Page    int8             `json:"page"`
		Results []models.Summary `json:"results"`
	}

	var response Response
	if err := json.Unmarshal(byteArray, &response); err != nil {
		log.Fatal(err)
	}
	return response.Results
}
