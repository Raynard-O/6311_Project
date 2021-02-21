package controllers

import (
	"fmt"
	"github.com/labstack/echo"
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
func AddEvent(t readerEventType, wg *sync.WaitGroup) {
	//check the kind of event
	if t {
		// add the event to readers event database
	fmt.Println("we are in the thread, readers function")
	}else {
		// add to books event database
		fmt.Println("we are in the thread, book function")
	}
	wg.Done()
}