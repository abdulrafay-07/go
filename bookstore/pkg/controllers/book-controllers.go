package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abdulrafay-07/go/bookstore/pkg/models"
	"github.com/abdulrafay-07/go/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()

	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// retrieve bookId
	bookId := vars["bookId"]

	// convert to int
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book id.")
	}

	bookDetails, _ := models.GetBookById(Id)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}

	utils.ParseBody(r, book)

	b := book.CreateBook()

	res, _ := json.Marshal(b)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book id.")
	}

	deletedBook := models.DeleteBook(Id)

	res, _ := json.Marshal(deletedBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book = &models.Book{}

	utils.ParseBody(r, book)

	vars := mux.Vars(r)

	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing book id.")
	}

	bookDetails, db := models.GetBookById(Id)

	// update data
	if book.Name != "" {
		bookDetails.Name = book.Name
	}
	if book.Author != "" {
		bookDetails.Author = book.Author
	}
	if book.Publication != "" {
		bookDetails.Publication = book.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
