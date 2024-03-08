package handler

import (
	"net/http"
	"util"

	"github.com/labstack/echo"
)

// sample function
func GetUser2(c echo.Context) error {

	// Path Variable -> dhlee
	// xxx.com:0000/users/dhlee
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func Show2(c echo.Context) error {
	// Query String -> team, dhlee
	// xxx.com:0000/users?team=obst&member=dhlee
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", memhber:"+member)
}

func List2(c echo.Context) error {
	// POST Form Submit
	team := c.FormValue("team")
	member := c.FormValue("member")
	util.LoadFile("data/books.csv")
	return c.String(http.StatusOK, "team:"+team+", memhber:"+member)

}
