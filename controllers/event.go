package controllers

import (
	"6311_Project/models"
	"6311_Project/storage"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"sync"
)

type InterfaceEvent interface {
	AddEvent(c echo.Context) error
}
// readerEventType
// determine the the event
type readerEventType bool

// AddEvent
// create event in a concurrent process.
func AddEvent(t readerEventType, wg *sync.WaitGroup, in interface{}, DB storage.InterfaceDB) {
	if t {
		// add the event to readers event database
		bE := in.(*models.ReaderEvent)

		events , err := DB.EventSave(bE, "events", true)
		if err != nil {
			log.Fatal(err)
		}
	fmt.Println("we are in the thread, readers function", events)
	}else {
		// add to books event database
		bE := in.(*models.AuthorEvent)

		events , err := DB.EventSave(bE, "events", true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("we are in the thread, book function", events)
	}
	wg.Done()
}
