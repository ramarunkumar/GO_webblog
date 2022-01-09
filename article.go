package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//------------------------------------------------dashboard--------------------------------------------------//

func dashboard(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/web_blog")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}

	row, err := db.Query("SELECT * from blog_user,blog WHERE id=1 AND article_id=1 ORDER BY blog_id ASC")
	if err != nil {
		fmt.Println("insert Error", err)
	}

	res := []Data{}

	for row.Next() {
		emp := Data{}

		err = row.Scan(&emp.ID, &emp.Name, &emp.Username, &emp.Password, &emp.Article_id, &emp.Blog_id, &emp.Title, &emp.Content)
		if err != nil {
			fmt.Println("scan error", err)
		}

		res = append(res, emp)
	}
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

	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/web_blog")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}

	row, err := db.Query("SELECT * from blog ORDER BY blog_id ASC ")
	if err != nil {
		fmt.Println("insert Error", err)
	}

	res := []Data{}

	for row.Next() {
		emp := Data{}

		err = row.Scan(&emp.Article_id, &emp.Blog_id, &emp.Title, &emp.Content)
		if err != nil {
			fmt.Println("scan error", err)
		}

		res = append(res, emp)

	}
	fmt.Println(len(res))

	render(c, gin.H{
		"title":   "Home Page",
		"payload": res}, "index.html")
}

//--------------------------------------------------createArticle-------------------------------------------------------//

func createArticle(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/web_blog")
	if err != nil {
		fmt.Println("could not connect to database: ", err)

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
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/web_blog")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	var articleList = []Article{}

	fmt.Println("title", title+"||"+"content", content)
	stmt := "INSERT INTO blog(title, content) VALUES ($1,$2)"
	fmt.Println(stmt)
	row := db.QueryRow(stmt, title, content).Scan(&title, &content)

	if row != nil {
		fmt.Println("inserted succesfully", row)
	}
	fmt.Println(row, db)

	a := Article{Blog_id: len(articleList) + 1, Title: title, Content: content}
	fmt.Println(db)
	return &a, nil

}

//-------------------------------------------update-------------------------------------------------------------

func editArticle(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/web_blog")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	row, err := db.Query("SELECT * from blog")
	if err != nil {
		fmt.Println("insert Error", err)
	}

	res := []Article{}

	blog_id := c.Query("blog_id")
	title := c.Query("title")
	content := c.Query("content")

	fmt.Println("blog id", blog_id)
	for row.Next() {
		emp := Article{}

		err = row.Scan(&emp.Article_id, &emp.Blog_id, &emp.Title, &emp.Content)
		if err != nil {
			fmt.Println("scan error", err)
		}

		res = append(res, emp)
	}
	fmt.Println(blog_id)

	render(c, gin.H{

		"blog":    blog_id,
		"title":   title,
		"content": content,

		"payload": res}, "update.html")
}

//------------------------------------------------update-------------------------------------------------------------//
func update(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/web_blog")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	blog := c.Request.URL.Query()
	if blog != nil {
		fmt.Print(blog)
	}

	blog_id := c.PostForm("blog_id")

	title := c.PostForm("title")
	content := c.PostForm("content")
	u := Data{Title: title, Content: content}

	stmt := "UPDATE  blog SET title= $1, content=$2 WHERE blog_id = '" + blog_id + "'"

	row := db.QueryRow(stmt, title, content).Scan(&blog_id)

	fmt.Println("blog_id:", blog_id, "title:", title, "content:", content)
	if row != nil {
		fmt.Println("update succesfully", row)
		render(c, gin.H{
			"titles": "blog:" + blog_id + " title " + title + "Content" + content + "   update Successful",

			"payload": u}, "submission-successful.html")

	} else {

		c.HTML(http.StatusNotFound, "errors/404", nil)
		fmt.Println("not inserted...")

	}

}

func deleteArticle(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/web_blog")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}

	fmt.Println("Delete")

	blog_id := c.Query("blog_id")

	fmt.Println("blog id", blog_id)

	row := db.QueryRow("DELETE FROM blog WHERE blog_id = '" + blog_id + "'").Scan(&blog_id)
	if row != nil {
		fmt.Println(row)
		render(c, gin.H{
			"payload": row}, "delete.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Print("Deleted")

}
