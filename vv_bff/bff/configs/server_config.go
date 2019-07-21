package configs


type ServerConfig struct {
	DisableCors              bool   `default:"false"`
	ActivateGraphqlGui       bool   `default:"false"`
	Port                     string `envconfig:"PORT" default:"8080"`
	QueryEndpoint            string `default:"/query"`
	AccessControlAllowOrigin string `split_words:"true"`
}