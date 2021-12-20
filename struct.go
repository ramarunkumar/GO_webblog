package main

import (
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
	Created_on string
}
type Data struct {
	ID         int
	Name       string
	Username   string
	Password   string `bson:"password"`
	Blog_id    int
	Article_id int
	Title      string
	Content    string
	Created_on string
}
