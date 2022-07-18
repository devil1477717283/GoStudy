package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func handleDelete(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	post, err := retrieve(id)
	if err != nil {
		return nil
	}
	err = post.delete()
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}
func handleGet(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	post, err := retrieve(id)
	if err != nil {
		return err
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return err
}
func handlePut(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	post, err := retrieve(id)
	if err != nil {
		return err
	}
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return err
}
func handlePost(w http.ResponseWriter, r *http.Request) error {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)
	err := post.create()
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return err
}
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}
