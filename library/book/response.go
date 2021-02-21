package book

import (
	"6311_Project/models"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)



type BookResponse struct {
	ID       primitive.ObjectID
	Name        string `json:"full_name"`
	Author        []string `json:"author"`
	Pages	int8 	`json:"pages"`
	Content string	`json:"content"`
}
type Response struct {
	User    BookResponse `json:"user"`
	Token   string       `json:"token"`
	Success bool         `json:"success"`
}

func BookResponseData(c echo.Context, book *models.Book, token string, channel string) error {
	response := Response{
		User: BookResponse{
			ID:    book.BookID  ,
			Name:    book.BookName,
			Author:  book.Authors,
			Pages:   book.Pages,
			Content: book.Content,
		},
		Success: true,
	}
	return c.JSONPretty(http.StatusOK, response, "")
}
