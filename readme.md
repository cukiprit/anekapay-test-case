# Anekapay Test Case

## Overview

Anekapay Test Case is a simple RESTful API built with Go and Gin. You can do CRUD in this API.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the API](#running-the-api)
  - [Seeding the DB](#seeding-the-db)
- [API Endpoints](#api-endpoints)
- [Example Request](#example-request)

## Features

- Create, read, update, and delete animal records.
- SQLite database for persistent storage.

## Technologies

- Go (Golang)
- Gin Web Framework
- SQLite

## Getting Started

### Prerequisites

- Go (1.23.1 or later)
- SQLite3

### Instalation

1. Clone the repository:

```
git clone https://github.com/cukiprit/anekapay-test-case
cd anekapay-test-case
```

2. Install the required Go packages:

```
go mod tidy
```

3. Start the application

```
go run ./cmd/main.go
```

> The API will start on localhost:8080

## API Endpoints

| Method | Endpoints      | Description               |
| ------ | -------------- | ------------------------- |
| POST   | `/animals`     | Create a new animal       |
| GET    | `/animals`     | Get all animals           |
| GET    | `/animals/:id` | Get an animal by ID       |
| PUT    | `/animals`     | Update an existing animal |
| Delete | `/animals/:id` | Delete an animal by ID    |

## Example Request

**Create Animal**

```
curl -X POST http://localhost:8080/animals \
-H "Content-Type: application/json" \
-d '{
    "id": 1,
    "name": "Lion",
    "class": "Mammal",
    "legs": 4
}'
```

**Expected Response**

```
{
    "id": 1,
    "name": "Lion",
    "class": "Mammal",
    "legs": 4
}
```

Notes:

- If an animal with the same ID already exists, you might receive:

```
{
    "error": "Animal with this ID already exists."
}
```

**Get All Animals**

```
curl -X GET http://localhost:8080/animals
```

**Expected Response**

```
[
    {
        "id": 1,
        "name": "Lion",
        "class": "Mammal",
        "legs": 4
    },
    {
        "id": 2,
        "name": "Tiger",
        "class": "Mammal",
        "legs": 4
    }
    // More animal objects...
]
```

Notes:

- If no animals are found, the response will be:

```
{
    "error": "No animals found."
}
```
