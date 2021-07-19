package main

import (
	"fmt"

	"github.com/emurray647/run_log/core/db"
	"github.com/emurray647/run_log/network"
)

func main() {
	fmt.Println("Running")

	dbConnection := db.DbConnection{}
	err := dbConnection.Open()
	if err != nil {
		panic(err.Error())
	}
	defer dbConnection.Close()

	err = dbConnection.CreateTables()
	if err != nil {
		panic(err.Error())
	}

	network.HandleRequests(&dbConnection)

}
