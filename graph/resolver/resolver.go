package resolver

//go:generate go run github.com/99designs/gqlgen generate
import (
	"modelo-graphql-go/internal/api/book"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	BookService book.Service
}
