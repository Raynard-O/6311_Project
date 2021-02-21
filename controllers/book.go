package controllers

import (
	"6311_Project/library"
	bookLibrary "6311_Project/library/book"
	"6311_Project/models"
	"6311_Project/storage"
	"context"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"sync"
	"time"
)

type books struct {
	title string
	wg sync.WaitGroup
}

func NewBook() *books {
	return &books{title: "books"}
}

func (b *books) Create(c echo.Context) error {
	DB, err := storage.MongoInit("ICDE", b.title, context.Background())

	if err != nil {
		log.Fatal(err)
	}
	params := new(library.CreateBookParams)
	// bind params
	err = c.Bind(params)
	if err != nil {
		return BadRequestResponse(c, err.Error())
	}

	book := parseObject(params)

	books, err := DB.BookSave(book, b.title)
	// add event to ICDE database
	//created go routines
	b.wg.Add(1)
	go func() {
		AddEvent(readerEventType(false), &b.wg)
	}()

	if err != nil {
		log.Fatal(err.Error())
	}
	//  wait for process to end
	b.wg.Wait()
	log.Println("reader successfully created")
	return bookLibrary.BookResponseData(c, &books, "", books.BookName)
}

func parseObject( params *library.CreateBookParams) *models.Book {

	aut := []string{}
	book := models.Book{
		BookID:     primitive.NewObjectID(),
		BookName:   params.Name,
		Pages:      params.Pages,
		Authors:    append(aut, params.Author),
		UploadDate: time.Now(),
		Content:    params.Content,
	}
	return &book
}

func (b *books) Login(c echo.Context) error {
	return nil
}