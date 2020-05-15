package bookrepo

import (
	"context"
	"go-app/domain/book"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Repo Interface
type Repo interface {
	CreateBook(book *book.Book) (*book.Book, error)
}

type bookRepo struct {
	db *mongo.Client
}

// NewBookRepo will instantiate Book Repository
func NewBookRepo(db *mongo.Client) Repo {
	return &bookRepo{
		db: db,
	}
}

func (b *bookRepo) CreateBook(book *book.Book) (*book.Book, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel() // releases resources if CreateBook completes before timeout elapses
	collection := b.db.Database("books-db").Collection("books")
	_, err := collection.InsertOne(ctx, *book)

	if err != nil {
		panic(err)
	}
	return book, nil
}
