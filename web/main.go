package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func setCORS(w http.ResponseWriter, r *http.Request) bool {
	// set CROS
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Max-Age", "1000")

	w.Header().Set("Access-Control-Allow-Headers", `Content-Type,X-Requested-With,Expires,Cache-Control,Origin,Current-Path,Accept,Pragma`)
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return false
	}
	return true
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	if !setCORS(w, r) {
		return
	}

	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由

	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	fmt.Println("=============")

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
