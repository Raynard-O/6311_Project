package main

import (
	"6311_Project/router"
)

func main()  {
	e := router.New()
	e.Logger.Fatal(e.Start(":8081"))
}
