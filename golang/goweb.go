package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// 建立借口函数和url路径的映射
	mux.HandleFunc("/hello", sayhello)
	mux.HandleFunc("/login", login)

	// db部分
	db, e := sql.Open("mysql", "root:root@tcp(localhost:3306)/gotest?charset=utf8")
	if e != nil { //如果连接出错,e将不是nil的
		fmt.Println("ERROR?")
		return
	} else {
		fmt.Println("connect mysql ok")
	}

	http.ListenAndServe(":9000", mux)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func login(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}
