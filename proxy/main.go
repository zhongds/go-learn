package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	http.HandleFunc("/", commponProxy) //设置访问的路由

	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	fmt.Println("=============")

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func commponProxy(w http.ResponseWriter, r *http.Request) {
	rpURL, err := url.Parse("http://proxy.com")
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(rpURL)

	proxy.ServeHTTP(w, r)
}

// 重定向的Proxy
func getRedirectProxy(baseUrl string, redirectPath string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rpURL, err := url.Parse(baseUrl)
		if err != nil {
			log.Fatal(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(rpURL)
		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = rpURL.Scheme
			req.URL.Host = rpURL.Host
			req.Host = rpURL.Host
			req.URL.Path = redirectPath
		}

		proxy.ServeHTTP(w, r)
	}
}
