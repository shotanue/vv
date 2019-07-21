package configs

type TmdbConfig struct {
	TmdbToken string `split_words:"true" required:"true"`
}