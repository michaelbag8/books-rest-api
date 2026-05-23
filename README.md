# Books REST API

A simple and lightweight RESTful API built with Go and the native `net/http` package for managing books.

This project demonstrates core backend development concepts using pure Golang without external frameworks.

---

## Features

* CRUD operations for books
* RESTful API design
* JSON request and response handling
* Route handling with `net/http`
* Request validation
* Error handling
* Clean project structure
* In-memory or database storage support

---

## Tech Stack

* Golang
* net/http
* JSON


---

## Project Structure

```bash id="8uv58u"
books-rest-api/
├── handlers/
├── models/
├── routes/
├── utils/
├── main.go
└── go.mod
```

---

## Getting Started

### Prerequisites

* Go 1.22+

---

## Installation

### Clone the repository

```bash id="9mxn7g"
git clone https://github.com/michaelbag8/books-rest-api.git
```

### Navigate into the project directory

```bash id="rxt3fz"
cd books-rest-api
```

### Install dependencies

```bash id="6rkhm9"
go mod tidy
```

### Run the application

```bash id="8mp6j6"
go run main.go
```

The server will start on:

```bash id="6q5r0l"
http://localhost:8080
```

---

## API Endpoints

| Method | Endpoint    | Description       |
| ------ | ----------- | ----------------- |
| GET    | /books      | Get all books     |
| GET    | /books/{id} | Get a single book |
| POST   | /books      | Create a new book |
| PUT    | /books/{id} | Update a book     |
| DELETE | /books/{id} | Delete a book     |

---

## Example Book Object

```json id="0okj31"
{
  "id": 1,
  "title": "The Go Programming Language",
  "author": "Alan Donovan",
  "published_year": 2015
}
```

---

## Running Tests

```bash id="jgrnr2"
go test ./...
```

---

## Future Improvements

* JWT authentication
* Database integration
* Pagination & filtering
* Docker support
* Swagger documentation
* Middleware implementation

---

## Author

Michael BAG

Built with ❤️ using pure Golang and the `net/http` package.
