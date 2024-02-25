package GoWebX

import (
	"net"
	"net/http"
)

type HandleFunc func(ctx Context)

type Server interface {
	http.Handler
	// Start 启动服务器
	// addr 监听地址 ":8080" or "localhost:8080"
	Start(addr string) error
	// AddRoute 增加路由注册功能
	// method 请求方法
	// path 路由
	// handleFunc 业务逻辑
	// 思路:核心API保持非常少,衍生API可以很多(Get,Post...)
	AddRoute(method string, path string, handleFunc HandleFunc)
}

// 确保一定实现了Server接口
var _ Server = &HttpServer{}

type HttpServer struct {
}

func (h *HttpServer) AddRoute(method string, path string, handleFunc HandleFunc) {
	//TODO implement me
	// 注册到路由树
	//panic("implement me")
}

func (h *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// 请求进来会调用这个方法
	// 首先构建context
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}

	h.serve(ctx)
}

func (h *HttpServer) Start(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return http.Serve(l, h)
}

func (h *HttpServer) serve(ctx *Context) {

}

// Get 请求
func (h *HttpServer) Get(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}

// Post 请求
func (h *HttpServer) Post(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}

// Put 请求
func (h *HttpServer) Put(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}

// Delete 请求
func (h *HttpServer) Delete(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}

// Options 请求
func (h *HttpServer) Options(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}

// Patch 请求
func (h *HttpServer) Patch(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}
