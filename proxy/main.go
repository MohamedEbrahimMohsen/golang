package main

import (
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// the server that we want to redirect to.
	url, _ := url.Parse("http://localhost:9090")
	proxy := httputil.NewSingleHostReverseProxy(url)

	// "/" means all the routes, for specific route such as any request has "books" change it to "/books"
	prefix := "/"
	routerGroup := e.Group(prefix)
	routerGroup.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			req := context.Request()
			res := context.Response().Writer

			req.Host = url.Host
			req.URL.Host = url.Host
			req.URL.Scheme = url.Scheme

			path := req.URL.Path
			req.URL.Path = strings.TrimLeft(path, prefix)

			proxy.ServeHTTP(res, req)
			return nil
		}
	})

	// allow all CORS.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Start(":8080")
}
