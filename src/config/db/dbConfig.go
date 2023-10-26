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

func LoadConfigDB() {

	userDatabase = os.Getenv("USER_DATABASE")
	passwordDatabase = os.Getenv("PASSWORD_DATABASE")
	databaseName = os.Getenv("DATABASE_NAME")

	createDateBaseIfNotExists()
}

func StartConnection() (*sql.DB, error) {

	connection, err := sql.Open("mysql", "root:root@/productdb?charset=utf8&parseTime=True&loc=Local&multiStatements=true")

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

	migration.Steps(2)

}

func createDateBaseIfNotExists() {

	db, err := sql.Open("mysql", "root:root@/")

	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + databaseName)

	if err != nil {
		log.Fatal(err)
	}
}
