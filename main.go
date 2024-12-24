package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"moe.best.annai/request"
	"moe.best.annai/resolver"
)

func main() {

	e := echo.New()
	e.GET("/bunny", func(c echo.Context) error {

		query := c.QueryParam("query")
		if len(query) == 0 {
			return c.String(http.StatusOK, resolver.Lookup(request.Request{}).String())
		}

		req := request.NewRequest(c.QueryParam("query"))
		return c.Redirect(http.StatusTemporaryRedirect, resolver.Lookup(req).String())
	})
	e.Logger.Fatal(e.Start(":8080"))

}
