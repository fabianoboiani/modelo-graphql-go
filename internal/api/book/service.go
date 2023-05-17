package book

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"modelo-graphql-go/internal/pkg/models"
)

type Service interface {
	GetAll() ([]*models.Book, error)
	CreateBook(book *models.Book) (*models.Book, error)
	GetBookById(bookId string) (*models.Book, error)
	EditBook(book models.Book) (*models.Book, error)
	DeleteBook(bookId string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]*models.Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) CreateBook(book *models.Book) (*models.Book, error) {

	newBook := models.Book{
		Id:       primitive.NewObjectID(),
		Title:    book.Title,
		Category: book.Category,
		Author:   book.Author,
	}

	result, err := s.repository.Insert(&newBook)

	return result, err
}

func (s *service) GetBookById(bookId string) (*models.Book, error) {
	book, err := s.repository.FindById(bookId)
	return book, err
}

func (s *service) EditBook(book models.Book) (*models.Book, error) {
	updatedBook, err := s.repository.Update(book)
	return updatedBook, err
}

func (s *service) DeleteBook(bookId string) error {
	err := s.repository.Delete(bookId)
	return err
}
