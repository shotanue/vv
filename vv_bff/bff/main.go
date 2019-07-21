package main

import (
	"bff/configs"
	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

func main() {
	var serverConfig configs.ServerConfig
	if err := envconfig.Process("", &serverConfig); err != nil {
		log.Fatal(err)
	}

	port := serverConfig.Port
	endpoint := serverConfig.QueryEndpoint

	if serverConfig.ActivateGraphqlGui {
		http.Handle("/", handler.Playground("GraphQL playground", endpoint))
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	}

	h := handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{}}))

	http.Handle(endpoint, withHeaders(h, serverConfig))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func withHeaders(h http.Handler, config configs.ServerConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		{
			var origin string
			if config.DisableCors {
				origin = "*"
			} else {
				if config.AccessControlAllowOrigin == "" {
					log.Print("${ACCESS_CONTROL_ALLOW_ORIGIN} is required.")
				}

				origin = config.AccessControlAllowOrigin
			}
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")

		// I added this for another handler of mine,
		// but I do not think this is necessary for GraphQL's handler
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}
