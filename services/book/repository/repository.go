package repository

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/kenshin579/analyzing-go-web-server-tips/errors"
	"github.com/kenshin579/analyzing-go-web-server-tips/models"
)

var (
	errBookNotFound = errors.NotFound("book")
)

type Repository interface {
	GetBook(id bson.ObjectId) (*models.Book, error)
	CreateBook(book models.Book) error
}
