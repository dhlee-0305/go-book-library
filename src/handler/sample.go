package handler

import (
	"fmt"
	"logger"
	"net/http"

	"github.com/labstack/echo"
)

// sample function
func SamplePath(c echo.Context) error {

	// Path Variable -> dhlee
	// xxx.com:0000/users/dhlee
	id := c.Param("id")
	logger.Info("SamplePath " + id)

	return c.String(http.StatusOK, id)
}

func SampleGet(c echo.Context) error {
	// Query String -> team, dhlee
	// xxx.com:0000/users?team=obst&member=dhlee
	team := c.QueryParam("team")
	member := c.QueryParam("member")

	logger.Info("SampleGet team:" + team + ", member:" + member)

	return c.String(http.StatusOK, "team:"+team+", memhber:"+member)
}

func SampleForm(c echo.Context) error {
	// POST Form Submit
	team := c.FormValue("team")
	member := c.FormValue("member")
	logger.Info("SampleForm team:" + team + ", member:" + member)

	return c.String(http.StatusOK, "team:"+team+", memhber:"+member)

}

func SampleJson(c echo.Context) error {
	//var json map[string]string
	var json map[string]interface{} = map[string]interface{}{}

	err := c.Bind(&json)
	if err != nil {
		panic(err)
	}

	logger.Info("SampleJson team:" + fmt.Sprintf("%v", json["team"]) + ", member:" + fmt.Sprintf("%v", json["member"]) + ", dept:" + fmt.Sprintf("%v", json["dept"]))

	return c.String(http.StatusOK, fmt.Sprintf("%v", json))
}
