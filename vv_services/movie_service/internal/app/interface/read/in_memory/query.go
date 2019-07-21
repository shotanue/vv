package in_memory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"movie_service/internal/app/domain/models"
	"movie_service/internal/app/interface/read"
	"os"
	"path/filepath"
	"runtime"
)

type Query struct {
}

func (Query) FetchComingSoonMovies() []models.Summary {
	// ファイルをOpenする
	f, err := os.Open(basePath + "/coming_soon_movies.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 一気に全部読み取り
	byteArray, err := ioutil.ReadAll(f)
	PrintMyPath()
	return pluckMovieSummary(byteArray)
}

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func PrintMyPath() {
	fmt.Println(basePath)
}

func NewQuery() read.MovieQuery {
	return Query{}
}

func (Query) FetchHotMovies() []models.Summary {
	// ファイルをOpenする
	f, err := os.Open(basePath + "/popular_movies.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 一気に全部読み取り
	byteArray, err := ioutil.ReadAll(f)
	PrintMyPath()
	return pluckMovieSummary(byteArray)
}

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
