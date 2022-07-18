package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	dsn := "root:devil147..@tcp(127.0.0.1:3306)/gwp"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}
func retrieve(id int) (post Post, err error) {
	post = Post{}
	err = db.QueryRow("select id,content,author from posts where id=?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}
func (post *Post) create() (err error) {
	_, err = db.Exec("insert into posts(id,content,author) values(?,?,?)", post.Id, post.Content, post.Author)
	return
}
func (post *Post) delete() (err error) {
	_, err = db.Exec("delete from posts where id=?", post.Id)
	return
}
func (post *Post) update() (err error) {
	stm, err := db.Prepare("update posts set content=?,author=? where id=?")
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
