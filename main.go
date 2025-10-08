package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "felipe" {
		// password is "12345"
		return "$1$LEQnFZBF$0OGKwyCrxIzqJBw8qd72e0"
	}
	return ""
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Please check your arguments.")
		os.Exit(1)
	}

	httpDir := "/home/felipe/tmp"
	port := "8080"
	authenticator := auth.NewBasicAuthenticator("example.com", Secret)

	fmt.Println("::::: Server started in: http://localhost:"+port, ":::: 🚀")

	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		fmt.Fprintf(w, "<html><body><h1>Hello, %s!</h1></body></html>", r.Username)
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	}))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
