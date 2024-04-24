# Book Test Web Service

This is a web service for managing books. It's written in Go and uses the Echo framework. Using Supabase as DB and Authentication services

This Service was intended to use PostgreSQL as it DB + Docker to Deploy, since its a bit difficult to Deploy and Configure it to Railway, i use Supabase Service instead

## Prerequisites

- Go (version 1.16 or higher)

## Setup

```sh
git clone https://github.com/gabeaventh/book_test.git
cd book_test
go mod download
```

Prepare Environtment Variables:
create a file named `.env` in the root folder

Replace the STRING inside <...> bracket with the Provided Info
or Ask the Repo Owner for it

```sh
SUPABASE_KEY:<KEY_STRING>
SUPABASE_URL:<URL_STRING>
```

Start the service:
```sh
go install github.com/cosmtrek/air@latest
air
```
The service will be available at http://localhost:8080.

API Endpoints
The service provides the following endpoints:

### User ###
```sh
POST /user/signup: Sign Up a New User OBJECT RESPONSE with JSON DATA
POST /user/signin: Sign In with a created user OBJECT RESPONSE with JSON DATA
POST /user/signout: Sign Out from authenticated user
```

User JSON DATA
```JSON
{
  "email": "email",
  "password": "password"
}
```
User OBJECT
```JSON
{
    "access_token": "",
    "token_type": "bearer",
    "expires_in": 3600,
    "refresh_token": "",
    "user": {
      "id": "uuid",
      "aud": "authenticated",
      "role": "authenticated",
      "email": "email@email.com",
      "invited_at": "0001-01-01T00:00:00Z",
      "confirmed_at": "2024-04-23T20:49:54.924674Z",
      "confirmation_sent_at": "0001-01-01T00:00:00Z",
      "app_metadata": {},
      "user_metadata": {
        "email": "email@email.com",
        "email_verified": false,
        "phone_verified": false,
        "sub": "uuid"
      },
      "created_at": "2024-04-23T20:49:54.91596Z",
      "updated_at": "2024-04-24T11:39:31.502661Z"
    },
    "provider_token": "",
    "provider_refresh_token": ""
  }
```

User OBJECT RESPONSE
```JSON
{
  "status": 200,
  "message": "User signed in successfully",
  "data": {User OBJECT}
}
```


### Book
> [!IMPORTANT]
Need Authentication Header
Please Refer to [This Section](#user "Goto user") on how to authenticate using POSTMAN / ThunderClient, then paste the access-token to Header
```sh
GET /book: Fetch all books OBJECT RESPONSE
POST /book: Create a new book OBJECT RESPONSE with JSON DATA
GET /book/:id: Fetch a book OJECT by ID
PUT /book/:id: Update a book OBJECT RESPONSE by ID with JSON DATA
DELETE /book/:id: Delete a book OBJECT RESPONSE by ID
```

Book JSON DATA
```JSON
{
  "title": "title",
  "author": "author",
  "published_date": "yyyy-mm-dd"
}
```

Book OBJECT
```JSON
{
  "id": 1,
  "title": "New Title 123",
  "author": "New Author12222355555",
  "published_date": "2024-04-01",
  "created_at": "2024-04-23T21:02:19.208139Z",
  "updated_at": "0001-01-01T00:00:00Z",
  "deleted_at": "0001-01-01T00:00:00Z"
}
```

Book OBJECT RESPONSE
```JSON
{
  "status": 200,
  "message": "Books fetched successfully",
  "data": {Book OBJECT}
}
```

List of Book OBJECT RESPONSE
```JSON
{
  "status": 200,
  "message": "Books fetched successfully",
  "data": [
    {Book OBJECT}
  ]
}
```

## Testing
To run the tests:

```sh
go test ./...
```