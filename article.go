package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//------------------------------------------------dashboard--------------------------------------------------//

func dashboard(c *gin.Context) {
	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/web_blog")
	if err != nil {
		fmt.Println("could not connect to database: ", err)
	}
	id := c.Query("id")
	article := c.Query("article")
	//row, err := db.Query("SELECT * FROM blog_user,blogs WHERE blog_user.id='" + id + "' AND blogs.article_id='" + article_id + "';")

	row, err := db.Query("SELECT * from blog_user,blogs Where '" + id + "'= '" + article + "' ORDER BY blog_id Asc ")
	if err != nil {
		fmt.Println("insert Error", err)
	}

	// emp := Data{}
	res := []Data{}

	for row.Next() {
		emp := Data{}

		err = row.Scan(&emp.ID, &emp.Name, &emp.Username, &emp.Password, &emp.Article_id, &emp.Blog_id, &emp.Title, &emp.Content)
		if err != nil {
			fmt.Println("scan error", err)
		}

		res = append(res, emp)

	}
	// fmt.Println(res)
	fmt.Println("id", id, "article:", article, id)
	render(c, gin.H{
		"title":   "Welcome to dashboard",
		"article": article,
		"id":      id,
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
		log.Fatalf("could not connect to database: %v", err)
	}

	row, err := db.Query("SELECT * from blogs ORDER BY blog_id ASC ")
	if err != nil {
		fmt.Println("insert Error", err)
	}

	// emp := Data{}
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

	} else {
		fmt.Println("connected...", db)
	}
	var articleList = []Article{}

	fmt.Println("title", title+"||"+"content", content)
	stmt := "INSERT INTO blogs(title, content) VALUES ($1,$2)"
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
	fmt.Println("show")

	db, err := sql.Open("postgres", "postgres://postgres:qwerty123@localhost:5432/web_blog")
	if err != nil {
		fmt.Println("could not connect to database: ", err)

	} else {
		fmt.Println("connected...", db)
	}
	row, err := db.Query("SELECT * from blogs")
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
		fmt.Println(blog_id)

		res = append(res, emp)
	}

	fmt.Println("emp", len(res))

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
	} else {
		fmt.Println("connected...", db)
	}
	blog := c.Request.URL.Query()
	if blog != nil {
		fmt.Print(blog)
	}

	blog_id := c.PostForm("blog_id")

	title := c.PostForm("title")
	content := c.PostForm("content")
	u := Data{Title: title, Content: content}

	stmt := "UPDATE  blogs SET title= $1, content=$2 WHERE blog_id = '" + blog_id + "'"

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
		log.Fatalf("could not connect to database: ", err)
	} else {
		fmt.Println("connected...", db)
	}
	fmt.Println("Delete")

	blog_id := c.Query("blog_id")

	fmt.Println("blog id", blog_id)

	row := db.QueryRow("DELETE FROM blogs WHERE blog_id = '" + blog_id + "'").Scan(&blog_id)
	if row != nil {
		fmt.Println(row)
		render(c, gin.H{
			"payload": row}, "delete.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Print("Deleted")

}
