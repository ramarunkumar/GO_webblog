//------------------------------------------signup---------------------------------------------------//

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func showRegistrationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Register"}, "signup.html")

}

func register(c *gin.Context) {

	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/web_blog")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}

	name := c.PostForm("name")
	username := c.PostForm("username")
	password := c.PostForm("password")

	u := User{Name: name, Username: username, Password: password}
	fmt.Println(u)

	if _, err := registerNewUser(name, username, password); err == nil {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		fmt.Println(err)

		stmt := "INSERT INTO blog_user(name, username, password) VALUES ('" + name + "','" + username + "', '" + string(hashedPassword) + "')"
		fmt.Println("query ", stmt)
		var dbname string
		var dbusername string
		var dbpassword string

		row := db.QueryRow(stmt)
		if row != nil {
			fmt.Println("error", row)
		}
		fmt.Println(&dbname, &dbusername, &dbpassword)

		if row != nil {
			fmt.Println("inserted succesfully", row.Scan(&u))
		}
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": username + " " + "Successful registrated  & logged in again  "}, "login.html")

	} else {
		fmt.Println("reg error", err)
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"ErrorTitle": "Registration Failed",

			"ErrorMessage": err.Error()})

	}

}

func registerNewUser(name, username, password string) (*User, error) {

	if len(password) < 8 && !strings.Contains(password, "@") {
		return nil, errors.New("password must 8 character and symbol @")
	}

	if strings.TrimSpace(password) == "" {
		return nil, errors.New("the password can't be empty")
	} else if UsernameAvailable(username) {
		return nil, errors.New("the username isn't available")
	}

	u := User{Name: name, Username: username, Password: password}

	return &u, nil
}

func UsernameAvailable(username string) bool {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/web_blog")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}

	stmt := "SELECT username FROM blog_user WHERE username = ('" + username + "')"
	err = db.QueryRow(stmt).Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("isusername error", err)
		}
		return false
	}

	return true
}
