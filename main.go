/*
 * @Author: cuishuming@baidu.com
 * @Date: 2023-07-12 22:23:51
 * @LastEditors: cuishuming@baidu.com
 * @LastEditTime: 2023-07-14 00:39:26
 * @FilePath: \Gee\main.go
 * @Description:
 * Copyright (c) 2023 by ${cuishuming@baidu.com}, All Rights Reserved.
 */
package main

import (
	"gen"
	"net/http"
)

func main() {
	// 创建 gen 实例
	r := gen.New()
	// 添加路由
	r.GET("/", func(c *gen.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gen</h1>")
	})
	r.GET("/hello", func(c *gen.Context) {
		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gen.Context) {
		c.JSON(http.StatusOK, gen.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	// 启动 web 服务
	r.Run(":9999")
}
