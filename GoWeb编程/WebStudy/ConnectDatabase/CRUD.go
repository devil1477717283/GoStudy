package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	dsn := "root:devil147..@tcp(127.0.0.1:3306)/gwp"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}
func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id,content,author from posts limit ?", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id,content,author from posts where id=?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}
func (post *Post) Create() (err error) {
	_, err = Db.Exec("insert into posts(id,content,author) values(?,?,?)", post.Id, post.Content, post.Author)
	return
}
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id=?", post.Id)
	return
}
func (post *Post) Update() (err error) {
	stm, err := Db.Prepare("update posts set content=?,author=? where id=?")
	if err != nil {
		fmt.Println("预处理失败")
	}
	defer stm.Close()
	res, err := stm.Exec(post.Content, post.Author, post.Id)
	if err != nil {
		fmt.Println("获取结果失败")
		return
	}
	count, err := res.RowsAffected()
	if count < 0 {
		fmt.Println("更新失败")
	} else {
		fmt.Println("更新成功")
	}
	return
}
func main() {
	post := Post{Content: "Hello world!", Author: "Sau Sheong"}
	fmt.Println(post)
	post.Create()
	fmt.Println(post)
	readPost, _ := GetPost(post.Id)
	println("DEBUG:")
	fmt.Println(readPost)
	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Id = 1
	readPost.Create()
	posts, _ := Posts(2)
	println("DEBUG:")
	fmt.Println(posts)
	readPost.Delete()
}
