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

	row, err := db.Query("SELECT * from blog_user,blogs Where id= article_id ORDER BY blog_id Asc ")
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
		"title": "Welcome to dashboard",

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

	row, err := db.Query("SELECT * from blog_user,blogs where id = article_id ORDER by blog_id ASC")
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

func editArticle(c *gin.Context) {
	fmt.Println("show")

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("not connected....", err)
	} else {
		fmt.Println("connected...", db)
	}
	row, err := db.Query("SELECT * from blogs")
	if err != nil {
		fmt.Println("insert Error", err)
	}

	emp := Article{}
	res := []Article{}
	blog := c.Request.URL.Query()
	if blog != nil {
		fmt.Print("Blog id: ", blog)
	}

	for row.Next() {

		var blog_id, article_id int
		var title, content string

		err = row.Scan(&article_id, &blog_id, &title, &content)
		if err != nil {
			fmt.Println("scan error", err)
		}
		fmt.Println(blog_id)
		emp.Blog_id = blog_id
		emp.Title = title
		emp.Content = content
		emp.Article_id = article_id
		res = append(res, emp)

	}

	render(c, gin.H{

		"blog":    blog,
		"payload": res}, "update.html")
}

//------------------------------------------------update-------------------------------------------------------------//
func update(c *gin.Context) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		fmt.Println("not connected....", err)
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
	fmt.Println(u)

	stmt := "UPDATE  blogs SET title='" + title + "',content='" + content + "' WHERE blog_id = '" + blog_id + "'"

	fmt.Println(stmt, blog)

	row := db.QueryRow(stmt).Scan(&blog_id)

	fmt.Println("blog_id:", blog_id, "title:", title, "content:", content)
	if row != nil {
		fmt.Println("update succesfully", row)
		render(c, gin.H{
			"titles": "blog:" + blog_id + " title " + title + "Content" + content + "   update Successful",

			"payload": u}, "submission-successful.html")
		fmt.Println(title, content)
	} else {

		c.HTML(http.StatusNotFound, "errors/404", nil)
		fmt.Println("not inserted...")

	}

}

func deleteArticle(c *gin.Context) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("not connected....", err)
	} else {
		fmt.Println("connected...", db)
	}
	fmt.Println("Delete")
	var blog_id int

	row := db.QueryRow("DELETE FROM blogs WHERE blog_id = 3;").Scan(&blog_id)
	if row != nil {

		render(c, gin.H{
			"payload": row}, "delete.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Print("Deleted")

}
