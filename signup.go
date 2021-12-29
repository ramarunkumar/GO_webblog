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

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if db != nil {
		fmt.Println("register database error", err)
	} else {
		fmt.Println("register database no error", db)
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
			"title": username + " " + "Successful registrated  & logged in   "}, "login-successful.html")

	} else {
		fmt.Println("reg error", err)
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"ErrorTitle": "Registration Failed",

			"ErrorMessage": err.Error()})

	}

}

func registerNewUser(name, username, password string) (*User, error) {

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println(" db error", err)
	} else {
		fmt.Println(" no error", db)
	}

	if len(password) < 8 && !strings.Contains(password, "@") {
		return nil, errors.New("password must 8 character and symbol @")
	}

	if strings.TrimSpace(password) == "" {
		return nil, errors.New("the password can't be empty")
	} else if UsernameAvailable(username) {
		return nil, errors.New("the username isn't available")
	}

	u := User{Name: name, Username: username, Password: password}
	fmt.Println("registernew user", db)

	return &u, nil
}

func UsernameAvailable(username string) bool {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		fmt.Println(" db error", err)
	} else {
		fmt.Println("  no error", db)
	}
	// SELECT password FROM public.users WHERE username='ram';

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
