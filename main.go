package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := "8080"


	fs := http.FileServer(http.Dir("/home/felipe/devops/tmp"))

	fmt.Println("::::: Server started in: http://localhost:"+port, ":::: 🚀")

	log.Fatal(http.ListenAndServe(":"+port, fs))

}
