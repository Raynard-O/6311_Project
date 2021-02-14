package router

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"sync"
	"testing"
	"time"
)
func process(index int , wg *sync.WaitGroup, e *echo.Echo) {
	fmt.Printf("Hitting endpoint  %+v\n",  "item")
	time.Sleep(time.Duration(index) * time.Second)
	e.Shutdown(nil)
	wg.Done()
}
func TestNew(t *testing.T) {
	e := New()
	wg := sync.WaitGroup{}
	wg.Add(1)
	c := make(chan error)
	go func() {
		//err := error(nil)
		if err := e.Start(":6311"); err != nil{
			c <- err
		}
		wg.Done()
	}()

	log.Println(<-c)
	wg.Wait()

	t.Errorf("Error starting server: %v", <-c)
}
