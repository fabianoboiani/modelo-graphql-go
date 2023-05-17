package book

import (
	"context"
	"dev.azure.com/lojasrenner/datalab-commons.git/pkg/logger"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"modelo-graphql-go/internal/db"
	"modelo-graphql-go/internal/pkg/models"
	"time"
)

var bookCollection *mongo.Collection

type Repository interface {
	FindAll() ([]*models.Book, error)
	Insert(book *models.Book) (*models.Book, error)
	FindById(bookId string) (*models.Book, error)
	Update(book models.Book) (*models.Book, error)
	Delete(bookId string) error
}

type repository struct {
}

func NewRepository() *repository {
	bookCollection = db.GetCollection(db.Client, "library", "book")
	return new(repository)
}

func (r *repository) FindAll() ([]*models.Book, error) {
	logger.Debug("Teste debug logger")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var books []*models.Book

	defer cancel()

	results, err := bookCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleBook models.Book
		if err = results.Decode(&singleBook); err != nil {
			//c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: gin.H{"error": err}})
			return nil, err
		}
		books = append(books, &singleBook)
	}
	return books, nil
}

func (r *repository) Insert(newBook *models.Book) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	result, err := bookCollection.InsertOne(ctx, newBook)

	if err != nil {
		return nil, err
	}
	newBook.Id = result.InsertedID.(primitive.ObjectID)
	return newBook, nil
}

func (r *repository) FindById(bookId string) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var book models.Book
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(bookId)

	err := bookCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *repository) Update(book models.Book) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{"author": book.Author, "category": book.Category, "title": book.Title}

	result, err := bookCollection.UpdateOne(ctx, bson.M{"_id": book.Id}, bson.M{"$set": update})

	if err != nil {
		return nil, err
	}
	//get updated book details
	var updatedBook models.Book
	if result.MatchedCount == 1 {
		err := bookCollection.FindOne(ctx, bson.M{"_id": book.Id}).Decode(&updatedBook)

		if err != nil {
			return nil, err
		}
	}

	return &updatedBook, nil
}

func (r *repository) Delete(bookId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(bookId)

	result, err := bookCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return err
	}

	if result.DeletedCount < 1 {
		return errors.New("Book not found")
	}

	return nil
}
