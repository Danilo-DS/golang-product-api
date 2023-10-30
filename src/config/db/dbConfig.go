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

//var newScripts int

var versionBeforeMigration uint
var versionAfterMigration uint

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

	log.Println("Starting Migrations")

	timeoutMigration, err := strconv.ParseInt(os.Getenv("TIMEOUT_MIGRATION"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	newScripts, err := strconv.ParseInt(os.Getenv("NUMBER_SCRIPTS_TO_RUN"), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	directoryMigration := os.Getenv("DIRECTORY_SCRITPS")

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

	runMigration(int(newScripts), migration)

	log.Println("Finishing Migrations")
}

func runMigration(numberScripts int, migration *migrate.Migrate) {

	migrationErr := migration.Steps(numberScripts) // The instruction specifies whether the scripts are creation or deletion, greater than 0 is creation, less than 0 is deletion

	_, dirty, err := migration.Version() // Get current version migration

	if err != nil {
		log.Fatal("Unexpected Error:", err)
	}

	if dirty {
		log.Println("Migration fail:", migrationErr)

		if err = rollbackMigration(numberScripts, migration); err != nil {
			log.Fatal("Rollback migration fail:", err)
		}

		log.Fatal("Rollback scripts executed successfully")
	}
}

func rollbackMigration(numberScripts int, migration *migrate.Migrate) error {
	version, _, err := migration.Version() // Get current version migration

	if err != nil {
		return err
	}

	err = migration.Force(int(version)) // Set migration version to latest version success

	if err != nil {
		return err
	}

	return migration.Steps(-numberScripts)
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
