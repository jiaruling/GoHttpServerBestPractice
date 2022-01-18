package global

import "net/http"

/*
   功能说明: 自定义多路复用器
   参考: https://darjun.github.io/2021/07/13/in-post/godailylib/nethttp/
   创建人: 贾汝凌
   创建时间: 2022/1/18 15:27
*/

type Middleware func(handler http.Handler) http.Handler

type MyMux struct {
	*http.ServeMux
	middleware []Middleware
}

func NewMyMux() *MyMux {
	return &MyMux{
		ServeMux: http.NewServeMux(),
	}
}

func (m *MyMux) Use(middleware ...Middleware) {
	m.middleware = append(m.middleware, middleware...)
}

//func (m *MyMux) Handle(pattern string, handler http.Handler) {
//	handler = ApplyMiddleware(handler, m.middleware...)
//	m.ServeMux.Handle(pattern, handler)
//}
//
//func (m *MyMux) HandleFunc(pattern string, handler http.HandlerFunc) {
//	newHandler := ApplyMiddleware(handler, m.middleware...)
//	m.ServeMux.Handle(pattern, newHandler)
//}

func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	h, _ := m.Handler(r)
	// 保证后添加的中间件能生效
	h = ApplyMiddleware(h, m.middleware...)
	h.ServeHTTP(w, r)
}

func ApplyMiddleware(handler http.Handler, middleware ...Middleware) http.Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	return handler
}
