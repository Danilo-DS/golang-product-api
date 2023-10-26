module product-api

go 1.20

require (
	github.com/go-sql-driver/mysql v1.7.1 // Driver Mysql

	github.com/joho/godotenv v1.5.1 // Dependency to load env variables
	
	github.com/gorilla/mux v1.8.0 // Dependency to config routers application
	
	// Dependency to perform migration
	github.com/golang-migrate/migrate/v4 v4.16.2
	github.com/hashicorp/errwrap v1.1.0
	github.com/hashicorp/go-multierror v1.1.1
	go.uber.org/atomic v1.7.0

	// Dto Mapper
	// https://pkg.go.dev/github.com/devfeel/mapper#section-readme
	github.com/devfeel/mapper v0.7.13
)

