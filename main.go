package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

// this will serve as my database for now
var books = []Book{
	{ID: 1, Title: "The Go Programming Language", Author: "Alan Donovan", Pages: 380},
	{ID: 2, Title: "Clean Code", Author: "Robert Martin", Pages: 431},
	{ID: 3, Title: "Start With Why", Author: "Simon Sinek", Pages: 401},
	{ID: 4, Title: "Breaking Out", Author: "Michael Bag", Pages: 500},
}

// since we already have four books our next will be 5
var nextID int = 5

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)

}

// getAllBooks to get the list of all the books
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}

	writeJSON(w, http.StatusOK, books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}

	var req Book
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	if req.Title == "" || req.Author == "" || req.Pages == 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "title, author and pages are required",
		})
		return
	}

	req.ID = nextID
	nextID++
	books = append(books, req)
	writeJSON(w, http.StatusCreated, req)

}

// update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}

	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid id convsersion",
		})
		return
	}

	var updatedBook Book

	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "bad request book",
		})
		return
	}

	for i, b := range books {
		if b.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook

			writeJSON(w, http.StatusOK, updatedBook)
			return
		}
	}
	writeJSON(w, http.StatusNotFound, map[string]string{
		"error": "book not found",
	})
}

func getBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return
	}

	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid URl",
		})
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
		return
	}

	for _, b := range books {
		if b.ID == id {
			writeJSON(w, http.StatusOK, b)
			return
		}
	}

	writeJSON(w, http.StatusNotFound, map[string]string{
		"error": "book not found",
	})

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

	}

}

// to handle all books method
func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAllBooks(w, r)
	case http.MethodPost:
		createBook(w, r)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
	}

}

// // to handle single book method
func bookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getBook(w, r)
	case http.MethodPut:
		updateBook(w, r)
	case http.MethodDelete:
		deleteBook(w, r)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
	}

}

func main() {
	http.HandleFunc("/books", booksHandler)
	//http.HandleFunc("/books", updateBook)
	http.HandleFunc("/books/", getBook)
	http.HandleFunc("/book/", updateBook)

	fmt.Println("serving is runing at http://localhost:8080/")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server failed to start", err)
		return
	}
}
