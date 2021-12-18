package main

import (
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Name     string
	Username string
	Password string `bson:"password"`
}

type Article struct {
	Blog_id    int
	Article_id int
	Title      string
	Content    string
	created_on time.Time
}
type Blog struct {
	User
	Article
}

var userList = []User{
	{ID: 1, Name: "ram", Username: "ram", Password: "123"},
	{ID: 2, Name: "arun", Username: "arun", Password: "123"},
}

var articleList = []Article{
	{

		Blog_id:    1,
		Article_id: 1,
		Title:      "Google",
		Content:    "Google LLC is an American multinational technology company that specializes in Internet-related services and products, which include online advertising technologies, a search engine, cloud computing, software, and hardware.",
	},
	{
		Blog_id:    2,
		Article_id: 2,
		Title:      "Facebook",
		Content:    "Meta Platforms, Inc., doing business as Meta and formerly known as Facebook, Inc., is a multinational technology conglomerate based in Menlo Park, California. The company is the parent organization of Facebook, Instagram, and WhatsApp, among other subsidiaries.",
	},
}
