# Go Task REST API

A simple REST API built with Go's standard `net/http` package.

## Features

- Create tasks
- Get all tasks
- Get task by ID
- Update tasks
- Delete tasks
- Request logging middleware
- JSON request/response handling

## API Endpoints

POST   /tasks
GET    /tasks
GET    /tasks/{id}
PUT    /tasks/{id}
DELETE /tasks/{id}

## Run

go run .

Server runs on:

http://localhost:8080
