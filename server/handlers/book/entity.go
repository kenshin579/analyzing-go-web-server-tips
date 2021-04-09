package book

import "github.com/kenshin579/analyzing-go-web-server-tips/models"

type getBookResponse struct {
	Book *models.Book `json:"book"`
}

type createBookResponse struct {
	Book *models.Book `json:"book"`
}
