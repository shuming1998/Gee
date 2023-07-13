/*
 * @Author: cuishuming@baidu.com
 * @Date: 2023-07-13 15:31:52
 * @LastEditors: cuishuming@baidu.com
 * @LastEditTime: 2023-07-13 16:59:28
 * @FilePath: /baidu/Gen/gen/gen.go
 * @Description:
 * Copyright (c) 2023 by ${cuishuming@baidu.com}, All Rights Reserved.
 */
package gen

import (
	"net/http"
)

// 定义 gen 使用的请求处理函数，用户借此来定义路由映射的处理方法
type HandlerFunc func(*Context)

// Engine 用于实现 ServeHTTP 接口
type Engine struct {
	// 路由映射表，根据不同的路由映射不同的处理方法
	router *router
}

// 构造返回 gen.Engine 实例
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// 用于增加 GET 请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// 用于增加 POST 请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 封装 ListenAndServe，用于启动 http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// ResponseWriter 可以构造针对该请求的响应
// Request 包含了该 HTTP 请求的所有的信息，比如请求地址、Header和Body等信息
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
