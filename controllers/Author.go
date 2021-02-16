package controllers

import "github.com/labstack/echo"

// Author
// class is a stack of author class objects
type Author struct {
	//Db
}

func NewAuthor() *Author {
	return &Author{}
}

func (a *Author) CreateUser(c echo.Context) error {
	return nil
}
