module product-api

go 1.20

require (
	// Dto Mapper
	// Doc: https://pkg.go.dev/github.com/devfeel/mapper#section-readme
	github.com/devfeel/mapper v0.7.13
	github.com/go-sql-driver/mysql v1.7.1 // Driver Mysql

	github.com/gorilla/mux v1.8.0 // Dependency to config routers application
	github.com/joho/godotenv v1.5.1 // Dependency to load env variables
	
	// Dependency to perform migration
	github.com/golang-migrate/migrate/v4 v4.16.2
	github.com/hashicorp/errwrap v1.1.0
	github.com/hashicorp/go-multierror v1.1.1
	go.uber.org/atomic v1.7.0
	
	// Go ORM Dependency
	// Doc: https://gorm.io/docs/
	github.com/jinzhu/inflection v1.0.0
	github.com/jinzhu/now v1.1.5 
	gorm.io/gorm v1.25.5 
	gorm.io/driver/mysql v1.5.2 // Driver orm to mysql
)
