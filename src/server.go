package main

import (
	"fmt"
	"log"
	"net/http"
	"product-api/src/config"
	"product-api/src/config/db"
	"product-api/src/routers"
)

func main() {

	config.LoadEnvironment()
	db.LoadConfigDB()
	db.LoadMigration()

	port := config.GetPort()

	fmt.Println("Server Up")
	fmt.Println("Port", port)
	log.Fatal(http.ListenAndServe(port, routers.Routers()))

}
