# Build a CRUD API with Golang

## Overview
This code essentially sets up a basic server for managing a list of movies, allowing clients to perform CRUD operations on the movie collection through HTTP requests.

## Features

- Route by a combination of paths and **HTTP methods** such as GET, POST, PUT and DELETE.
- Test web APIs with Postman. 

## Installation

### Prerequisites

- Go 1.20
- gorilla/mux v1.8.X


### Setup
To set up the routes which mean endpoints, you need to install the `gorilla/mux`. You can do this by running:
```
go get -u github.com/gorilla/mux
```
## Usage
Click on the `<> Code` button and download the file by selecting `Download ZIP`. Once the download is complete, open the downloaded file in your IDE such as GoLand.

To start the application, run the following command in the terminal:

```
go run main.go
```

Navigate to `http://localhost:8000/movies` in your web browser. Enter an {id} number after `8000/movies/`, the application will display the movie's information including id, isbn, title and director in dictionary format.

## Application's Files
- `main.go`: The main file that contains the CRUD application's implementation.
- `go.mod`: File contains all packages of a dependency which is required by the module.
- `go.sum`: File contains cryptographic hashes of the module.
- `GO-movies.postman_collection.json`: Automatically generated JSON file by Postman.



## References
"2. Build A CRUD API With Golang, Building 11 Projects", Akhil Sharma (2021), YouTube\
 Also the code includes Automated API test monitors by Postman and Newman.
