package controllers

import (
	"6311_Project/library"
	"github.com/stretchr/testify/assert"
	"go/types"
	"testing"
)



func TestNewBook(t *testing.T) {
	newReader := NewBook()
	assert.NotNil(t, newReader)
}

func TestParseObject(t *testing.T) {
	params := library.CreateBookParams{
		Name:    "Book Test",
		Author:  "Raynard",
		Pages:   7,
		Content: "qwerjblnlkon yf wesh ug se vvhjbhiu wcv uih",
	}
	book := parseObject(&params)
	//fmt.Println(book)

	assert.NotNil(t, book)
	assert.IsType(t,string("types.String") , params.Name )
	assert.IsType(t,string("types.String") , params.Content )
	assert.IsType(t,string("types.String")  , params.Author )
	assert.IsType(t,int8(types.Int8) , params.Pages)
}
