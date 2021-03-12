package router

import (
	"6311_Project/config"
	"6311_Project/controllers"
	"context"
	"flag"
	"fmt"
	"github.com/labstack/echo"
)

type Control struct {
	debug bool
}
var ctx context.Context
type debugContext string
func New() *echo.Echo {
	c := control()

	e := echo.New()

	//j := debugContext("control")
	//ctx = context.WithValue(context.Background(), j, c.debug)
	//distance(ctx,j)

	secret, _ := config.LoadSecrets(c.debug)
	fmt.Println(secret.HmacSigningKey)

	auth := e.Group("/auth")



	author := controllers.InterfaceUsers(controllers.NewAuthor())
	reader := controllers.InterfaceUsers(controllers.NewReader())
	book := controllers.InterfaceUsers(controllers.NewBook())

	authors := auth.Group("/author")
	authors.POST("/signup", author.Create)
	authors.POST("/signin", author.Search)


	readers := auth.Group("/reader")
	readers.POST("/signup", reader.Create)
	readers.POST("/signin", reader.Search)


	books := authors.Group("/book")
	books.POST("/add", book.Create)
	books.GET("/search", book.Search)
	books.GET("/select", book.Select)

	return e
}




func distance(ctx context.Context, k debugContext)  {
	v := ctx.Value(k)

	if v != nil {
		fmt.Println("found", v)
	}else{
		fmt.Println("hello ", v)
	}
}




func control() *Control {
	c := Control{
		//debug: *flag.Bool("debug", true, "set debug characteristics")
	}
	flag.BoolVar(&c.debug, "debug", true, "set debug characteristics")
	flag.Parse()
	return &c
}