package main

import (
	"github.com/insaneadinesia/go-mockemok/config"
	"github.com/insaneadinesia/go-mockemok/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	// load config
	e := echo.New()
	conf := config.Load()

	for _, mock := range conf.Mock {
		mock := mock
		group := e.Group(mock.Group)

		for _, request := range mock.Request {

			group.Use(setRequestMiddleware(request))

			switch request.Method {
			case "GET":
				group.GET(request.Path, handler.Get)
				break
			case "POST":
				group.POST(request.Path, handler.Post)
				break
			case "PUT":
				group.PUT(request.Path, handler.Put)
				break
			case "PATCH":
				group.PATCH(request.Path, handler.Patch)
				break
			case "DELETE":
				group.DELETE(request.Path, handler.Delete)
				break
			}
		}
	}

	e.Logger.Fatal(e.Start(conf.GetPort()))
}

func setRequestMiddleware(request config.MockGroupRequest) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("Request", request)
			return next(c)
		}
	}
}
