package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"moe.best.annai/resolver"
	"moe.best.annai/session"
)

func main() {

	e := echo.New()
	e.GET("/bunny", func(c echo.Context) error {

		query := c.QueryParam("query")
		if len(query) == 0 {
			return c.String(http.StatusOK, resolver.Lookup(session.Session{}).String())
		}

		req := session.NewSession(c.QueryParam("query"), c.Request().Header)

		return c.Redirect(http.StatusTemporaryRedirect, resolver.Lookup(req).String())
	})
	e.Logger.Fatal(e.Start(":8080"))

}
