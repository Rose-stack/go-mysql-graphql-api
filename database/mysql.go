package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"go-graphql-api/dbmodel"
)

// a vriable to store database connection
var DBInstance *gorm.DB

// Var for error handling
var err error

// the db connection string
var CONNECTION_STRING string = "root:root@tcp(localhost:3306)/?charset=utf8&parseTime=True&loc=Local"

// connectiong to the db
func ConnectDB() {
	// pass the db connection string
	ConnectionURI := CONNECTION_STRING
	// check for db connection
	DBInstance, err = gorm.Open("mysql", ConnectionURI)
	if err != nil {
		fmt.Println(err)
		panic("Database connection attempt was unsuccessful.....")
	} else {
		fmt.Println("Database Connected successfully.....")
	}
	// log all dabase operations perfomed by this connection
	DBInstance.LogMode(true)
}

// Create a database
func CreateDB() {
	DBInstance.Exec("CREATE DATABASE IF NOT EXISTS Blog_Posts")
	DBInstance.Exec("USE Blog_Posts")
}

// migrate and sync the model to create a db table
func MigrateDB() {
	DBInstance.AutoMigrate(&dbmodel.Post{})
	fmt.Println("Database migration completed....")
}
