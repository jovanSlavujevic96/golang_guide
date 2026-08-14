# Project Description

A Go-powered "Event Booking" REST API

- GET /events -> Get a list of available events
- GET /events/{id} -> Get a list of available events
- POST /events -> Create a new bookable event **(Auth Required)**
- PUT /events/{id} -> Update an event **(Auth Required)** **(Only by creator)**
- DELETE /events{id} -> Delete an event **(Auth Required)** **(Only by creator)**
- POST /signup -> Create a new user
- POST /login -> Authenticate user **(Auth Token (JWT))**
- POST /events/{id}/register -> register user for event **(Auth Required)**
- DELETE /events/{id}/register -> cancel registration **(Auth Required)**

## Installing the GIN Framework

```
$ go get -u github.com/gin-gonic/gin
$ go mod tidy
```

## Installing SQLite
```
$ go get github.com/mattn/go-sqlite3
```

## Installing JWT
```
$ go get -u github.com/golang-jwt/jwt/v5
```
