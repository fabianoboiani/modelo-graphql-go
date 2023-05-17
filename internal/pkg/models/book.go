package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"modelo-graphql-go/internal/pkg/domain"
)

type Book struct {
	Id       primitive.ObjectID  `json:"id,omitempty" bson:"_id"`
	Author   Author              `json:"author,omitempty" validate:"required"`
	Category domain.BookCategory `json:"category,omitempty" validate:"required"`
	Title    string              `json:"title,omitempty" validate:"required"`
}

type Author struct {
	Name string `json:"author,omitempty" validate:"required"`
}
