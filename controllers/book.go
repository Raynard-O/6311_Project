package controllers

import (
	"6311_Project/library"
	bookLibrary "6311_Project/library/book"
	"6311_Project/models"
	"6311_Project/storage"
	"6311_Project/thirdparty"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strconv"
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
	log.Println(params)
	book := parseObject(params)

	books, err := DB.BookSave(book, b.title)
	log.Println("book and author event  successfully created")
	// add event to ICDE database
	e := &models.AuthorEvent{
		EventID: primitive.NewObjectID(),
		BookTitle: books.BookName,
		Authors:  books.Authors,
		EventType: "author",
		Event:  models.Events{
			EventTime: time.Time{},
		},
	}

	//created go routines
	b.wg.Add(1)
	go func() {
		AddEvent(readerEventType(false), &b.wg,e, DB)
	}()
	err = thirdparty.Recommend(c, books.Authors)
	if err != nil {
		return BadRequestResponse(c, err.Error())
	}
	//  wait for process to end
	b.wg.Wait()
	log.Println("book and author event  successfully created")
	return bookLibrary.BookResponseData(c, &books, "", books.BookName)
}


func parseObject( params *library.CreateBookParams) *models.Book {

	p, err := strconv.Atoi(params.Pages)

	if err != nil {
		log.Fatal(err)
	}
	book := models.Book{
		ID:     primitive.NewObjectID(),
		BookName:   params.Name,
		Pages:      int8(p),
		Authors:    params.Author,
		UploadDate: time.Now(),
		Content:    params.Content,
	}
	return &book
}

func (b *books) Search(c echo.Context) error {

	DB, err := storage.MongoInit("ICDE", "books", context.Background())

	if err != nil {
		return InternalError(c, err.Error())
	}
	params := new(library.SearchBookParams)
	err = c.Bind(params)
	if err != nil {
		return InternalError(c, err.Error())
	}
	field := make(map[string]interface{})
	filter := make(map[string]interface{})
	filter["$regex"] = params.Name

	field["book_name"] = filter

	var output interface{}
	fmt.Println(field)
	err = DB.FindMany(field, nil, nil, 0, 0, &output)
	if err != nil {
		return BadRequestResponse(c, err.Error())
	}
	c.Set("id", "gdjhbcvschbsb")

	return c.JSONPretty(200, output,"")
}


func (b *books) Select(c echo.Context) error {
	DB, err := storage.MongoInit("ICDE", "books", context.Background())

	if err != nil {
		return InternalError(c, err.Error())
	}

	id := c.QueryParam("id")
	OID, err  := primitive.ObjectIDFromHex(id)

	if err != nil {
		return BadRequestResponse(c, err.Error())
	}
	book := new(models.Book)

	err = DB.FindByID(OID, nil, &book )
	if err != nil {
		return BadRequestResponse(c, err.Error())
	}

	i, err := primitive.ObjectIDFromHex("60802bd690da4a5904739017")

	fmt.Printf("context is : %v\n", i)
	e := &models.ReaderEvent{
		EventID:    primitive.NewObjectID(),
		BookTitle:  book.BookName,
		BookID: book.ID,
		Authors: book.Authors,
		UserID:     i,
		EventType: "reader",
		Event:      models.Events{
			EventTime: time.Time{},
		},
	}
	//created go routines
	b.wg.Add(1)
	go func() {
		AddEvent(readerEventType(true), &b.wg,e, DB)
	}()

	if err != nil {
		log.Fatal(err.Error())
	}
	//  wait for process to end
	b.wg.Wait()
	log.Println("book and reader event  successfully selected")
	return c.JSONPretty(200, book, "")
}


func  GetuserBooks( c echo.Context) error{
	DB, err := storage.MongoInit("ICDE", "readerEventType", context.Background())

	if err != nil {
		return InternalError(c, err.Error())
	}
	m := make(map[string]interface{})
	filter := make(map[string]interface{})
	filter["$regex"] = ""
	m["book_title"] = filter
	log.Println(m)
	var output interface{}
	err = DB.FindManyHistory(m, nil, nil, 4, 0, &output)

	return c.JSONPretty(http.StatusOK, output, "")
}
// when logged in, send the id to the frontend, let the frontend embed it in the ip everytime it sends a requst so we know who is sending the request