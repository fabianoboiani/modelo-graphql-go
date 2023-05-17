package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"modelo-graphql-go/graph/generated"
	"modelo-graphql-go/graph/resolver"
	"modelo-graphql-go/internal/api/book"
	"modelo-graphql-go/internal/configs"
	"net/http"
)

const defaultPort = "8080"

func Init() {
	conf := configs.GetConfig()
	port := conf.GetString("server.port")
	if port == "" {
		port = defaultPort
	}

	//Services and Repositorys
	bookRepository := book.NewRepository()
	bookService := book.NewService(bookRepository)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{BookService: bookService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
