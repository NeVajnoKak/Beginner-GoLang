package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NeVajnoKak/Beginner-GoLang/tree/main/third-project-bookstore/pkg/models"
	"github.com/NeVajnoKak/Beginner-GoLang/tree/main/third-project-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, "Error fetching books: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, "Error marshalling books: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid book ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	bookDetails, _, err := models.GetBookById(ID)
	if err != nil {
		http.Error(w, "Error fetching book: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Error marshalling book: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b, err := CreateBook.CreateBook()
	if err != nil {
		http.Error(w, "Error creating book: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(b)
	if err != nil {
		http.Error(w, "Error marshalling book: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid book ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	err = models.DeleteBook(ID)
	if err != nil {
		http.Error(w, "Error deleting book: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid book ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	bookDetails, db, err := models.GetBookById(ID)
	if err != nil {
		http.Error(w, "Error fetching book: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	if err := db.Save(&bookDetails).Error; err != nil {
		http.Error(w, "Error updating book: "+err.Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Error marshalling book: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
