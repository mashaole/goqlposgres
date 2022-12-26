package main

import (
	"goqlposgress/graph"
	"goqlposgress/postgres"
	"log"
	"net/http"
	"os"

	customMiddleware "goqlposgress/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	DB := postgres.New(&pg.Options{
		User:     "PostgresDBUsername",
		Password: "PostgresDBPassword",
		Database: "meetmeup_dev",
	})
	defer DB.Close()
	DB.AddQueryHook(postgres.DBLogger{})
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	userRepo := postgres.UsersRepo{DB: DB}
	c := graph.Config{Resolvers: &graph.Resolver{
		MeetupsRepo: postgres.MeetupsRepo{DB: DB},
		UsersRepo:   postgres.UsersRepo{DB: DB},
	}}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(userRepo))

	queryHandler := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graph.DataloaderMiddleware(DB, queryHandler))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
