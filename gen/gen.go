package gen

import (
	"fmt"
	"net/http"
)

// 定义 gen 使用的请求处理函数，用户借此来定义路由映射的处理方法
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 用于实现 ServeHTTP 接口
type Engine struct {
	// 路由映射表，根据不同的路由映射不同的处理方法
	router map[string]HandlerFunc
}

// 构造返回 gen.Engine 实例
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
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
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

