package database

import "database/sql"

func InitDB() *sql.DB {
	connectionString := "root:admin@tcp(localhost:33006)/northwind"
	databaseConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return databaseConnection
}
