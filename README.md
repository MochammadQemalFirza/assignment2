# Assignment 2 Hacktiv8

A Golang RESTful API with Gin Framework and Postgresql database for Assignment 2

# Table of Contents

* [Features](#features)
* [Getting Started](#getting-started)
  * [Installation](#installation)
  * [Prequisited](#prequisited)
* [API Documentation](#api-documentation)

## Features
* Customer can create order items
* Customer can get order items
* Customer can update order items
* Customer can delete order items

## Getting Started
### Installation
* Clone this repository
```
https://github.com/MochammadQemalFirza/assignment2.git
```
* Setting .env with this template
```
#Database
DB_HOSTNAME="localhost"
DB_PORT="5432"
DB_USER="postgres"

DB_PASSWORD="your_paassword"

DB_DATABASE="your_database"
DB_SSLMODE="disable"

# Port
PORT = "8083"
```
* Run the server by typing this command on root folder project
```
go mod tidy
```
```
go run server.go
```

### Prequisited
* [Postgresql](https://www.postgresql.org/download/) - SQL Database.
* [Golang](https://go.dev/dl/) - Programming Language
