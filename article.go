package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//------------------------------------------------dashboard--------------------------------------------------//

func dashboard(c *gin.Context) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("not connected....", err)
	} else {
		fmt.Println("connected...", db)
	}

	row, err := db.Query("SELECT * from blog_user,blogs where id = article_id ORDER BY blog_id Asc ")
	if err != nil {
		fmt.Println("insert Error", err)
	}

	emp := Data{}
	res := []Data{}

	for row.Next() {
		var article_id, blog_id int
		var title, content string
		var id int
		var name, username, password string
		err = row.Scan(&id, &name, &username, &password, &article_id, &blog_id, &title, &content)
		if err != nil {
			fmt.Println("scan error", err)
		}
		emp.Article_id = article_id
		emp.Blog_id = blog_id
		emp.Title = title
		emp.Content = content
		emp.ID = id
		emp.Name = name
		emp.Username = username
		emp.Password = password

		res = append(res, emp)

	}
	fmt.Println(len(res))

	render(c, gin.H{
		"title":   "Welcome to dashboard",
		"payload": res},
		"dashboard.html")

}

//--------------------------------------------------profile---------------------------------------------------//

func showArticleCreationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func showIndexPage(c *gin.Context) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("not connected....", err)
	} else {
		fmt.Println("connected...", db)
	}

	row, err := db.Query("SELECT * from blog_user,blogs where id = article_id")
	if err != nil {
		fmt.Println("insert Error", err)
	}

	emp := Data{}
	res := []Data{}

	for row.Next() {
		var article_id, blog_id int
		var title, content string
		var id int
		var name, username, password string
		err = row.Scan(&id, &name, &username, &password, &article_id, &blog_id, &title, &content)
		if err != nil {
			fmt.Println("scan error", err)
		}
		emp.Article_id = article_id
		emp.Blog_id = blog_id
		emp.Title = title
		emp.Content = content
		emp.ID = id
		emp.Name = name
		emp.Username = username
		emp.Password = password

		res = append(res, emp)

	}
	fmt.Println(len(res))

	render(c, gin.H{
		"title":   "Home Page",
		"payload": res}, "index.html")
}

//--------------------------------------------------createArticle-------------------------------------------------------//

func createArticle(c *gin.Context) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("not connected....", err)
	} else {
		fmt.Println("connected...", db)
	}
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := createNewArticle(title, content); err == nil {
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Println(title, content)
}

func createNewArticle(title, content string) (*Article, error) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("not connected....", err)
	} else {
		fmt.Println("connected...", db)
	}
	var articleList = []Article{}

	fmt.Println("title", title+"||"+"content", content)
	stmt := "INSERT INTO blogs(title, content) VALUES ('" + title + "','" + content + "') "
	fmt.Println(stmt)
	row := db.QueryRow(stmt).Scan(&title, &content)

	if row != nil {
		fmt.Println("inserted succesfully", row)
	}
	fmt.Println(row, db)

	a := Article{Blog_id: len(articleList) + 1, Title: title, Content: content}
	fmt.Println(db)
	return &a, nil

}

//-------------------------------------------update-------------------------------------------------------------

func showupdate(c *gin.Context) {
	fmt.Println("show")
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("not connected....", err)
	} else {
		fmt.Println("connected...", db)
	}
	row, err := db.Query("SELECT * from blog_user,blogs where id = article_id")
	if err != nil {
		fmt.Println("insert Error", err)
	}

	emp := Data{}
	res := []Data{}

	for row.Next() {
		var article_id, blog_id int
		var title, content string
		var id int
		var name, username, password string
		err = row.Scan(&id, &name, &username, &password, &article_id, &blog_id, &title, &content)
		if err != nil {
			fmt.Println("scan error", err)
		}
		emp.Article_id = article_id
		emp.Blog_id = blog_id
		emp.Title = title
		emp.Content = content
		emp.ID = id
		emp.Name = name
		emp.Username = username
		emp.Password = password

		res = append(res, emp)

	}
	fmt.Println(len(res))

	render(c, gin.H{
		"para":    "Update Article",
		"payload": emp}, "update.html")

}

func update(c *gin.Context) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("not connected....", err)
	} else {
		fmt.Println("connected...", db)
	}
	blog_id := c.PostForm("blog_id")
	fmt.Println("blog id", blog_id)
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := updateArticle(title, content, blog_id); err == nil {
		var d Data
		render(c, gin.H{
			"titles":  "update Successful",
			"data":    d,
			"payload": a}, "submission-successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Println(title, content)
}

func updateArticle(title, content, blog_id string) (*Article, error) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("not connected....", err)
	} else {
		fmt.Println("connected...", db)
	}

	fmt.Println("title:" + title + "|content:" + content)

	stmt := "UPDATE  blogs SET title='" + title + "',content='" + content + "' WHERE blog_id = '" + blog_id + "'"

	fmt.Println(stmt, blog_id)
	row := db.QueryRow(stmt).Scan(&blog_id)
	fmt.Println("blog_id:", blog_id, "title:", title, "content:", content)
	if row != nil {
		fmt.Println("update succesfully", row)
	}
	fmt.Println(row, db)

	fmt.Println(row)

	return &Article{}, nil
}

func deleteArticle(c *gin.Context) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("not connected....", err)
	} else {
		fmt.Println("connected...", db)
	}
	var blog_id int
	stmt := "DELETE blog_id FROM blogs WHERE blog_id = $1"

	row := db.QueryRow(stmt).Scan(&blog_id)
	if row != nil {
		fmt.Println("DELETED", row)
	}
}
