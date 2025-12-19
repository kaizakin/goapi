# Go API Project

A simple Go API project demonstrating a basic structure with `chi` router, middleware, and mock database integration.

## Description

This project implements a RESTful API service using Go. It features:

- **Router**: [go-chi/chi](https://github.com/go-chi/chi) for routing.
- **Logging**: [logrus](https://github.com/sirupsen/logrus) for structured logging.
- **Schema Validation**: [gorilla/schema](https://github.com/gorilla/schema) for request parameter decoding.
- **Architecture**: Clean architecture with separate handlers, middleware, and tools.

## Project Structure

```
.
├── api/                # API definitions and error handling
├── cmd/
│   └── api/            # Application entry point (main.go)
├── internal/
│   ├── handlers/       # HTTP handlers
│   ├── middleware/     # HTTP middleware (Authorization, etc.)
│   └── tools/          # Internal tools and database interfaces
├── go.mod              # Go module definition
└── README.md           # Project documentation
```

## API Endpoints

### Account

- **GET** `/account/coins`

  - Retrieves coin balance.
  - **Auth**: Requires `Authorization` header.
  - **Params**: `username` (query parameter).

  **Example Request:**

  ```bash
  curl -v -H "Authorization: ABC" "http://localhost:8000/account/coins?username=karthik"
  ```
