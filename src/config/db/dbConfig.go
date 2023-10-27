package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql" //MySql Driver

	//Migration Imports Begin
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	//Migration Imports End
)

var userDatabase string
var passwordDatabase string
var databaseName string
var databaseHost string
var databasePort string
var connectionUrl string

// LoadConfigDB loads environment variables required for DB access configuration
func LoadConfigDB() {

	userDatabase = os.Getenv("USER_DATABASE")
	passwordDatabase = os.Getenv("PASSWORD_DATABASE")
	databaseName = os.Getenv("DATABASE_NAME")
	databaseHost = os.Getenv("DATABASE_HOST")
	databasePort = os.Getenv("DATABASE_PORT")
	connectionUrl = os.Getenv("CONNECTION_URL")

	createDateBaseIfNotExists()
}

func buildConnectionUrl() string {
	return fmt.Sprintf(connectionUrl, userDatabase, passwordDatabase, databaseHost, databasePort, databaseName)
}

// StartConnection open connection with database
func StartConnection() (*sql.DB, error) {

	connection, err := sql.Open("mysql", buildConnectionUrl())

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if err = connection.Ping(); err != nil {
		connection.Close()
		return nil, err
	}

	return connection, nil
}

// LoadMigration load and to run scripts sql
func LoadMigration() {

	timeoutMigration, err := strconv.ParseInt(os.Getenv("TIMEOUT_MIGRATION"), 10, 64)
	directoryMigration := os.Getenv("DIRECTORY_SCRITPS")

	if err != nil {
		log.Fatal(err)
	}

	connection, err := StartConnection()
	defer connection.Close()

	if err != nil {

		log.Fatal(err)
	}

	configMigration := &mysql.Config{
		MigrationsTable:  "SCHEMA_MIGRATIONS",
		DatabaseName:     databaseName,
		NoLock:           true,
		StatementTimeout: time.Second * time.Duration(timeoutMigration),
	}

	driver, err := mysql.WithInstance(connection, configMigration)

	if err != nil {
		log.Fatal(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(directoryMigration, "mysql", driver)

	if err != nil {
		log.Fatal(err)
	}

	// The instruction specifies whether the scripts are creation or deletion, greater than 0 is creation, less than 0 is deletion
	migration.Steps(1)

}

func createDateBaseIfNotExists() {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/", userDatabase, passwordDatabase, databaseHost, databasePort))

	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + databaseName)

	if err != nil {
		log.Fatal(err)
	}
}
