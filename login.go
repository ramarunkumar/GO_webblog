//------------------------------------------login---------------------------------------------------//

package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func showLoginPage(c *gin.Context) {

	render(c, gin.H{
		"title": "Login",
	}, "login.html")

}

func login(c *gin.Context) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if db != nil {
		fmt.Println("login db error", err)
	} else {
		fmt.Println("login db no error", db)
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, password)

	if isUserValid(username, "qwerty123", password) {

		fmt.Println("user is valid")
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		emp := Article{}
		res := []Article{}
		res = append(res, emp)
		render(c, gin.H{
			"title": username,

			"data": res},
			"login-successful.html")

	} else {

		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid username and password provided"})
	}

}

func isUserValid(username, pgPassword string, uPwd string) bool {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, pgPassword, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if db != nil {
		fmt.Println("isuservalid db ", db)
	} else {
		fmt.Println("isuservalid db error", err)
	}

	u := User{Username: username, Password: uPwd}

	var tmp User

	stmt := "SELECT password FROM blog_user WHERE username='" + username + "'"
	fmt.Println("stamt", stmt)
	row := db.QueryRow(stmt)

	fmt.Println(row)

	err = row.Scan(&tmp.Password)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
	}

	fmt.Println(tmp.Username, tmp.Password)
	if err = bcrypt.CompareHashAndPassword([]byte(tmp.Password), []byte(u.Password)); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//--------------------------logout----------------------------------------
func logout(c *gin.Context) {

	c.SetCookie("token", "", -1, "", "", false, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
