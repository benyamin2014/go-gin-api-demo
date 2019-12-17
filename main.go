package main

import (
	db "jtapi/repo"
	router "jtapi/route"
)

func main() {

	defer db.SqlDB.Close()
	route := router.Setup()
	route.Run(":8080")
}
