package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./storage.db")
	if err != nil {
		log.Fatal(err)
	}
	s := initStorage(db)
	e := Endpoint{&s}
	router := gin.Default()
	router.GET("/user/:name", e.getUser)
	router.POST("/user/:name", e.postUser)

	router.Run(":8080")
}
