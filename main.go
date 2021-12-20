package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var err error

const (
	host     = "localhost"
	port     = 5432
	users    = "postgres"
	password = "qwerty123"
	dbname   = "web_blog"
)

var r *gin.Engine

func main() {

	r = gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	initializeRoutes()
	r.Run()
}

func dbCon() (db *sql.DB) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("not connected....", err)
	} else {
		fmt.Println("connected...", db)
	}
	return db

}

func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
