package router

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
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

// TestNew
// checks for error in the server instance
func TestNew(t *testing.T) {
	e := New()

	//con
	assert.NotNil(t, e)
}
