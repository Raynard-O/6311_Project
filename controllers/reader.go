package controllers

import (
	"6311_Project/library"
	"6311_Project/library/reader"
	"6311_Project/models"
	"6311_Project/storage"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type Reader struct {
	title string
	DB *storage.InterfaceDB

}

func NewReader() *Reader {
	return &Reader{
		title: "readers",
	}
}

func (a *Reader) Create(c echo.Context) error {

	DB, err := storage.MongoInit("ICDE", a.title, context.Background())

	if err != nil {
		log.Fatal(err)
	}
	params := new(library.CreateUserParams)
	// bind params
	err = c.Bind(params)
	if err != nil {
		return BadRequestResponse(c, err.Error())
	}

	if params.Email == "" || params.Password == "" || params.ConfirmPassword == "" || params.FullName == "" || params.Username == "" {
		return BadRequestResponse(c, library.EMPTY)
	}
	//check if user details already exist
	if params.Email != "" {
		if  err = DB.FindOneUser(a.title, "email", params.Email, nil); err == nil {
			return BadRequestResponse(c, library.EmailTaken)
		}

		if _, err = DB.FindOne(a.title, "email", params.Email); err == nil {
			return BadRequestResponse(c, library.EmailTaken)
		}
		log.Println(err)
	}
	if params.Username != "" {
		if  err = DB.FindOneUser(a.title, "username", params.Username, nil); err == nil {
			return BadRequestResponse(c, library.EmailTaken)
		}
		if _, err = DB.FindOne(a.title, "email", params.Email); err == nil {
			return BadRequestResponse(c, library.UsernameTaken)
		}
	}
	hash, err := library.GenerateHash(params.Password)

	user := models.Reader{
		ID:     primitive.NewObjectID(),
		Fullname: params.FullName,
		Username: params.Username,
		Email:     params.Email,
		Password: hash,
		Books:  nil,
		Active: true,
	}

	users, err := DB.ReaderSave(user, a.title)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("reader successfully created")
	return reader.AuthorResponseData(c, &users, "", user.Email)

}


func  (a *Reader) Search(c echo.Context) error {
	DB, err := storage.MongoInit("ICDE", a.title, context.Background())

	if err != nil {
		log.Fatal(err)
	}
	params := new(library.LoginParams)
	err = c.Bind(params)
	if err != nil {
		return BadRequestResponse(c, err.Error())
	}

	userData := new(models.Reader)
	if params.Email != "" {
		fmt.Print(params.Email)
		DB.FindOneUser(a.title, "email", params.Email, &userData)
		//user, err := DB.FindOne("users", "email", params.Email)
		if err != nil {
			return BadRequestResponse(c, err.Error())
		}
		//userData = user
	} else if params.Username != "" {
		DB.FindOneUser(a.title, "username", params.Username, &userData)
		//user, err := DB.FindOne("users", "username", params.Email)
		if err != nil {
			return BadRequestResponse(c, err.Error())
		}
		//userData = user
	}

	passBool := library.CompareHashWithPassword(userData.Password, params.Password)
	log.Print(passBool)
	if !passBool {
		return BadRequestResponse(c, library.WrongPassword)
	}
	log.Println("reader successfully logged in")
	return reader.LoginUser(c, userData)
}

func  (a *Reader) Select(c echo.Context) error {
	return nil
}

func GetAllNotifications(c echo.Context) error{
	DB, err := storage.MongoInit("ICDE", "readers", context.Background())

	if err != nil {
		log.Fatal(err)
	}

	i, err := primitive.ObjectIDFromHex("60802bd690da4a5904739017")
	results := new(models.Reader)
	DB.FindByID(i, nil, &results)


	return c.JSONPretty(http.StatusOK, results.Notification, "")
}
