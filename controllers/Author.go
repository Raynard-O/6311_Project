package controllers

import (
	"6311_Project/library"
	"6311_Project/library/author"
	"6311_Project/models"
	"6311_Project/storage"
	"context"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

// Author
// class is a stack of author class objects
type Author struct {
	//Db
}
//
func NewAuthor() *Author {
	return &Author{}
}

func (a *Author) Create(c echo.Context) error {
	DB, err := storage.MongoInit("ICDE", "authors", context.Background())


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
		if  err = DB.FindOneUser("authors", "email", params.Email, nil); err == nil {
			return BadRequestResponse(c, library.EmailTaken)
		}

		if _, err = DB.FindOne("authors", "email", params.Email); err == nil {
			return BadRequestResponse(c, library.EmailTaken)
		}
		log.Println(err)
	}
	if params.Username != "" {
		if  err = DB.FindOneUser("authors", "username", params.Username, nil); err == nil {
			return BadRequestResponse(c, library.EmailTaken)
		}
		if _, err = DB.FindOne("authors", "email", params.Email); err == nil {
			return BadRequestResponse(c, library.UsernameTaken)
		}
	}
	hash, err := library.GenerateHash(params.Password)
	user := models.Author{
		ID:     primitive.NewObjectID(),
		Fullname: params.FullName,
		Username: params.Username,
		Email:     params.Email,
		Password: hash,
		Books:  nil,
		Active: true,
	}

	users, err := DB.AuthorSave(user, "authors")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("author successfully created")
	log.Println(users)
	return author.AuthorResponseData(c, &users, "", user.Email)

}

func  (a *Author) Search(c echo.Context) error {
	DB, err := storage.MongoInit("ICDE", "authors", context.Background())

	if err != nil {
		return InternalError(c, err.Error())
	}
	params := new(library.LoginParams)
	err = c.Bind(params)
	if err != nil {
		return InternalError(c, err.Error())
	}

	userData := new(models.Author)
	if params.Email != "" {
		//fmt.Println(params.Email)
		err = DB.FindOneUser("authors", "email", params.Email, &userData)
		//userDate2, err := DB.FindOne("users", "email", params.Email)
		if err != nil {
			return BadRequestResponse(c, err.Error())
		}
		//fmt.Println(userData )
	} else if params.Username != "" {
		err = DB.FindOneUser("authors", "username", params.Username, &userData)
		//user, err := DB.FindOne("users", "username", params.Email)
		if err != nil {
			return BadRequestResponse(c, err.Error())
		}

	}

	passBool := library.CompareHashWithPassword(userData.Password, params.Password)
	log.Print(passBool)
	if !passBool {
		return BadRequestResponse(c, library.WrongPassword)
	}
	log.Println("author successfully logged in")
	return author.LoginUser(c, userData)
}

func  (a *Author) Select(c echo.Context) error {
	return nil
}
