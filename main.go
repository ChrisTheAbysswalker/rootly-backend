package main

import (
	"github.com/ChrisTheAbysswalker/rootly-backend/db"
)

func main() {
	db := db.DbConnection()
	RunServer(db)
}