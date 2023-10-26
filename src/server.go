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

	fmt.Println("Server Up")
	log.Fatal(http.ListenAndServe(config.GetPort(), routers.Routers()))

}
