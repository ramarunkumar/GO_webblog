package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

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

	row, err := db.Query("SELECT * from blogs ORDER BY blog_id ASC")
	if err != nil {
		fmt.Println("insert Error", err)
	}

	emp := Article{}
	res := []Article{}

	for row.Next() {
		var id, article_id, blog_id int
		var title, content string
		err = row.Scan(&article_id, &blog_id, &title, &content)
		if err != nil {
			fmt.Println("scan error", err)
		}
		emp.Article_id = id
		emp.Blog_id = blog_id
		emp.Title = title
		emp.Content = content
		res = append(res, emp)

	}
	fmt.Println(len(res))
	render(c, gin.H{
		"title":   "Home Page",
		"payload": res}, "index.html")
}

//---------------------------------------------------------------------------------------------------------

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

	articleList = append(articleList, a)
	fmt.Println(db)
	return &a, nil

}

//------------------------------------------------------------------------------------------------------//

func getArticle(c *gin.Context) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("not connected....", err)
	} else {
		fmt.Println("connected...", db)
	}
	emp := Article{}
	res := []Article{}
	stmt := "SELECT * from blogs "
	row := db.QueryRow(stmt)
	if err != nil {
		fmt.Println(" Error", row)
	}
	var id, article_id, blog_id int
	var title, content string
	err = row.Scan(&article_id, &blog_id, &title, &content)
	if err != nil {
		fmt.Println("scan error", err)
	}
	emp.Article_id = id
	emp.Blog_id = blog_id
	emp.Title = title
	emp.Content = content

	res = append(res, emp)

	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := getArticleByID(articleID); err == nil {
			render(c, gin.H{
				"title":   article.Title,
				"payload": res}, "article.html")

		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
	fmt.Println(db)
}

func getArticleByID(article_id int) (*Article, error) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println("not connected....", err)
	} else {
		fmt.Println("connected...", db)
	}
	a := Article{Article_id: article_id}

	if a.Article_id == article_id {
		return &a, nil

	}
	return nil, errors.New("Article not found")
}

//--------------------------------------------------------------------------------------------------------

// func deleteArticle(c *gin.Context) {
// 	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
// 	db, err := sql.Open("postgres", dbinfo)
// 	if err != nil {
// 		fmt.Println("not connected....", err)
// 	} else {
// 		fmt.Println("connected...", db)
// 	}
// 	var article_id string
// 	stmt := "DELETE FROM users WHERE id = $1;"

// 	row := db.QueryRow(stmt, article_id)
// 	if row != nil {
// 		fmt.Println("deleted", row)
// 	}

// }

// func updateArticle(c *gin.Context) {
// 	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, users, password, dbname)
// 	db, err := sql.Open("postgres", dbinfo)
// 	if err != nil {
// 		fmt.Println("not connected....", err)
// 	} else {
// 		fmt.Println("connected...", db)
// 	}
// 	title := c.PostForm("title")
// 	content := c.PostForm("content")
// 	article_id := c.PostForm("article_id")
// 	stmt := "UPDATE users SET title = $2, content = $3 WHERE article_id = $1;"

// 	row := db.QueryRow(stmt, title, content)
// 	if err != nil {
// 		panic(err.Error())
// 	} else {
// 		fmt.Println(row)
// 	}

// 	log.Println("UPDATE:  title:"+title+"| content: "+content, article_id)
// }
