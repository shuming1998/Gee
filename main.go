package main

import (
	"fmt"
	"net/http"
	"gee"
)


func main() {
	// 创建 gee 实例
	r := gee.New()

	// 添加路由
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	// 启动 web 服务
	r.Run(":9999")
}