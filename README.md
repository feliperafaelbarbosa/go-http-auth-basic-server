# Basic HTTP Server with Authentication in Go

This project is a simple HTTP server written in Go that demonstrates basic HTTP authentication using the [go-http-auth](https://github.com/abbot/go-http-auth) library. The server requires a username and password to access its endpoints and serves static files from a specified directory.

## Features

- **Basic HTTP Authentication:** Requires a username and password for access.
- **Static File Server:** Serves files from a configurable directory.
- **Minimal and Clean Code:** Ideal for learning and portfolio demonstration.

## Usage

1. **Clone the repository:**
   ```bash
   git clone https://github.com/feliperafaelbarbosa/go-http-auth-basic-server.git
   cd go-http-auth-basic-server
   ```

2. **Run the server:**
   ```bash
   go run main.go <arg1> <arg2>
   ```

   > **Note:** The server expects exactly two arguments. Adjust the code or usage as needed for your scenario.

3. **Access the server:**
   Open your browser and navigate to [http://localhost:8080](http://localhost:8080). You will be prompted for credentials.

   - **Username:** felipe
   - **Password:** 12345

## Customization

- Change the `httpDir` variable to serve files from a different directory.
- Update the `Secret` function to add more users or change password hashes.

## Dependencies

- [go-http-auth](https://github.com/abbot/go-http-auth)
