package gee

import (
	"net/http"
)

// HandlerFunc defines the request handler used by gee
// 定义了路由映射处理方法，用户可自定义使用
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
// 放置路由表将路由与handler映射
type Engine struct {
	router *router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{newRouter()}
}

// 被GET和POST封装的底层func，用请求方法与路由生成key映射handler
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
// 将路由和处理方法注册到映射表router
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
// 封装用engine启动 Web 服务，启动listen
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 执行注册的方法，否则返回404
// 此接口接管所有http请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}