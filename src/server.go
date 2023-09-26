package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Server Up")
	log.Fatal(http.ListenAndServe(":3000", nil))

}
