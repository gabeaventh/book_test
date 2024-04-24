# Book Test Web Service

This is a web service for managing books. It's written in Go and uses the Echo framework. Using Supabase as DB and Authentication services

## Prerequisites

- Go (version 1.16 or higher)

## Setup

```sh
git clone <repository-url>
cd <project-directory>
go mod download
```
Start the service:
```sh
go install github.com/cosmtrek/air@latest
air
```
The service will be available at http://localhost:8080.

API Endpoints
The service provides the following endpoints:

Book
```sh
GET /book: Fetch all books
POST /book: Create a new book with JSON DATA
GET /book/:id: Fetch a book by ID
PUT /book/:id: Update a book by ID with JSON DATA
DELETE /book/:id: Delete a book by ID
```

Book JSON DATA
```JSON
{
  "title": "title",
  "author": "author",
  "published_date": "yyyy-mm-dd"
}
```

User
```sh
POST /user/signup: Sign Up a New User with JSON DATA
POST /user/signin: Sign In with a created user with JSON DATA
POST /user/signout: Sign Out from authenticated user
```

User JSON DATA
```JSON
{
  "email": "email",
  "password": "password"
}
```

## Testing
To run the tests:

```sh
go test ./...
```