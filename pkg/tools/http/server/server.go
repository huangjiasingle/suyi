package server

import (
	"fmt"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"k8s.io/klog/v2"
)

// Server is swrap of api server
type Server struct {
	// router is the fasthttp router
	router *router.Router
	// interceptors is the the server reject interceptor list
	interceptors []*func(ctx *fasthttp.RequestCtx) error
	// prefix is the api prefix path
	prefix string
}

// New return a new server
func New(prefix string) *Server {
	return &Server{router.New(), nil, prefix}
}

// Registry registry a handler to the server router
func (s *Server) Registry(path, method string, handle fasthttp.RequestHandler) *Server {
	path = fmt.Sprintf("%v%v", s.prefix, path)
	switch method {
	case "GET":
		s.router.GET(path, handle)
	case "HEAD":
		s.router.HEAD(path, handle)
	case "POST":
		s.router.POST(path, handle)
	case "PUT":
		s.router.PUT(path, handle)
	case "DELETE":
		s.router.DELETE(path, handle)
	case "OPTIONS":
		s.router.OPTIONS(path, handle)
	case "PATCH":
		s.router.PATCH(path, handle)
	}
	return s
}

// Healthz add healthz handler
func (s *Server) Healthz(handle fasthttp.RequestHandler) {
	s.router.GET("/healthz", handle)
}

// Handler handler the http request
func (s *Server) Handler(ctx *fasthttp.RequestCtx) {
	var sub time.Duration

	defer func() { klog.Infof("%s  %v  %s", ctx.String(), ctx.Response.StatusCode(), sub) }()

	t := time.Now()
	for _, interfaceptor := range s.interceptors {
		if interfaceptor != nil {
			if (*interfaceptor)(ctx) != nil {
				sub = time.Since(t)
				return
			}
		}
	}
	s.router.Handler(ctx)
	sub = time.Since(t)
}

// ServerHTTP start http server listener
func (s *Server) ServerHTTP(address string) error {
	return fasthttp.ListenAndServe(address, s.Handler)
}

// AddInterceptor add interceptor to the server
func (s *Server) AddInterceptor(interceptor func(ctx *fasthttp.RequestCtx) error) {
	s.interceptors = append(s.interceptors, &interceptor)
}
