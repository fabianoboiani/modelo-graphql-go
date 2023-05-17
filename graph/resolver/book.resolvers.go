package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"modelo-graphql-go/graph/generated"
	"modelo-graphql-go/internal/pkg/models"
)

// ID is the resolver for the id field.
func (r *bookResolver) ID(ctx context.Context, obj *models.Book) (string, error) {
	return obj.Id.Hex(), nil
}

// Category is the resolver for the category field.
func (r *bookResolver) Category(ctx context.Context, obj *models.Book) (int, error) {
	return int(obj.Category), nil
}

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

type bookResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *bookResolver) Author(ctx context.Context, obj *models.Book) (*models.Author, error) {
	return &obj.Author, nil
}
