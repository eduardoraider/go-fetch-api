# Go Fetch API Using Mux for HTTP Routing

This project fetches data from an external API and returns the result as JSON.

## Using Mux for HTTP Routing

The Mux router for HTTP routing, provided by the "net/http" package, allows us to define routes and map them to specific handler functions.

For example, in the NewHandler function, I created a new Mux router using http.NewServeMux(). Then I used mux.HandleFunc to define a route ("/") and map it to the moviesHandler function, which handles incoming HTTP requests.

This allows us to easily define and manage multiple routes in our application.

## Usage

1. Run the HTTP server:

```
go run cmd/fetch-api/main.go
```

2. Test the server using cURL, Postman or any web browser:

```
curl http://localhost:8080/movies
```

You should see a JSON response containing movies fetched from the external API.

## Project Structure

- **cmd/fetch-api**: Contains the main.go file, which serves as the entry point of the application. It can import and utilize the api package to set up HTTP handlers.


- **internal/**: Houses internal packages that are specific to the project and should not be imported by other projects. Here, we have the api and external packages.


- **internal/api**: This package comprises HTTP handlers to handle requests. The handler.go file may contain functions to manage different routes and call functions from the external package.


- **internal/external**: This package encompasses logic to interact with external APIs. The movie_api.go file may contain functions to fetch movies from an external API.


- **pkg/**: Contains packages that can be imported by other projects. Here, we have the entity package.


- **pkg/entity**: Contains domain entities used throughout the project. The movie.go file may contain the definition of the Movie structure.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

by **Eduardo Raider** - Software Engineer