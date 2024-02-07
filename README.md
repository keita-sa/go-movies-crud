# go movies crud

## Overview
This code essentially sets up a basic server for managing a list of movies, allowing clients to perform CRUD operations on the movie collection through HTTP requests.

## Features

- Route by a combination of paths and **HTTP methods** such as GET, POST, PUT and DELETE.
- Test web APIs with Postman. 

## Installation

### Prerequisites

- Go 1.20


### Setup
To set up the routes which mean endpoints, you need to install the `gorilla/mux`. You can do this by running:
```
go get -u github.com/gorilla/mux
```
## Usage
Click on the `<>　Code` button and download the file by selecting `Download ZIP`. Once the download is complete, open the downloaded file in your IDE such as GoLand.


Navigate to `http://localhost:8000/movies` in your web browser. Enter an {id} number after `8000/movies/`, the application will display the movie's information including id, isbn, title and director in dictionary format.





### References;
"2. Build A CRUD API With Golang, Building 11 Projects", Akhil Sharma (2021), YouTube\
 Also the code includes Automated API test monitors by Postman and Newman.
