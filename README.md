# Go API Project

This is a simple Go-based API server that responds with a JSON object containing the email address, current date and time, and a GitHub URL. The project is intended to demonstrate a basic REST API using Go's net/http package.

### Project Description

This project creates an API server that listens on port 8080 and responds to HTTP GET requests to the root URL (/). The server returns a JSON response with the following data:

- email: A hardcoded email address (inuoshios@gmail.com).
- current_datetime: The current date and time in RFC3339 format.
- github_url: A hardcoded GitHub profile URL (https://github.com/inuoshios/first).

### Setup Instructions

To run this project locally, follow these steps:

- Clone the repository (or copy the code into your project folder).
    ```sh
    git clone https://github.com/inuoshios/first
    ```
    ```sh
    cd first
    ```
- Run the server
    ```sh
    go run main.go
    ```

### API Documentation
#### Endpoint
- URL: /
- Method: GET
- Content-Type: application/json

#### Request Format
No request body is needed for this endpoint. It is a simple GET request.

#### Response Format
The response is a JSON object with the following fields:

```json
{
  "email": "inuoshios@gmail.com",
  "current_datetime": "2025-01-31T12:34:56Z",
  "github_url": "https://github.com/inuoshios/first"
}
```

### Example Usage
#### Request:

```sh
curl http://localhost:8080
```

### Backlink
https://hng.tech/hire/golang-developers
