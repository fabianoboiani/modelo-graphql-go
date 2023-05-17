package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"modelo-graphql-go/graph/generated"
	"modelo-graphql-go/graph/model"
	"modelo-graphql-go/internal/pkg/domain"
	"modelo-graphql-go/internal/pkg/models"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.BookMutationResponse, error) {
	newBook := &models.Book{
		Author:   models.Author{Name: input.Author},
		Category: domain.BookCategory(input.Category),
		Title:    input.Title,
	}

	if validateErr := validate.Struct(newBook); validateErr != nil {
		return nil, validateErr
	}

	if !domain.BookCategory.IsValid(newBook.Category) {
		return nil, errors.New("Category Invalid")
	}

	res, err := r.BookService.CreateBook(newBook)

	if err != nil {
		return nil, err
	}

	return &model.BookMutationResponse{
		Code:    http.StatusOK,
		Success: true,
		Message: "Book Created",
		Book:    res,
	}, nil
}

// UpdateBook is the resolver for the updateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id string, input model.NewBook) (*model.BookMutationResponse, error) {
	newBook := &models.Book{
		Author:   models.Author{Name: input.Author},
		Category: domain.BookCategory(input.Category),
		Title:    input.Title,
	}

	objId, _ := primitive.ObjectIDFromHex(id)
	newBook.Id = objId

	if validateErr := validate.Struct(newBook); validateErr != nil {
		return nil, validateErr
	}

	if !domain.BookCategory.IsValid(newBook.Category) {
		return nil, errors.New("Category Invalid")
	}

	res, err := r.BookService.EditBook(*newBook)

	if err != nil {
		return nil, err
	}

	return &model.BookMutationResponse{
		Code:    http.StatusOK,
		Success: true,
		Message: "Book Updated",
		Book:    res,
	}, nil
}

// DeleteBook is the resolver for the deleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id string) (*model.BookMutationResponse, error) {
	err := r.BookService.DeleteBook(id)

	if err != nil {
		return nil, err
	}

	return &model.BookMutationResponse{
		Code:    http.StatusOK,
		Success: true,
		Message: "Book Deleted",
		Book:    nil,
	}, nil
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*models.Book, error) {
	return r.BookService.GetAll()
}

// Status is the resolver for the status field.
func (r *queryResolver) Status(ctx context.Context) (*model.Health, error) {
	return &model.Health{Active: true}, nil
}

// GetBookByID is the resolver for the getBookById field.
func (r *queryResolver) GetBookByID(ctx context.Context, id string) (*model.BookMutationResponse, error) {
	res, err := r.BookService.GetBookById(id)

	if err != nil {
		return nil, err
	}

	return &model.BookMutationResponse{
		Code:    http.StatusOK,
		Success: true,
		Message: "Book found",
		Book:    res,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var validate = validator.New()
