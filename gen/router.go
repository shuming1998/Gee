/*
 * @Author: cuishuming@baidu.com
 * @Date: 2023-07-13 16:23:30
 * @LastEditors: cuishuming@baidu.com
 * @LastEditTime: 2023-07-13 16:38:35
 * @FilePath: /baidu/Gen/gen/router.go
 * @Description:
 * Copyright (c) 2023 by ${cuishuming@baidu.com}, All Rights Reserved.
 */
package gen

import (
	"log"
	"net/http"
)


type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}