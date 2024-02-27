package main

import (
	"net/http"

	"read"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	// 첫 화면
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/list", list)
	e.POST("/read", read.ReadList)

	e.Logger.Fatal(e.Start(":1323")) // localhost:1323
}

func getUser(c echo.Context) error {

	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", memhber:"+member)
}

func list(c echo.Context) error {
	team := c.FormValue("team")
	member := c.FormValue("member")
	return c.String(http.StatusOK, "team:"+team+", memhber:"+member)

}
