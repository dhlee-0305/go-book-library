package router

import (
	"net/http"

	echo "github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"

	handler "handler"
)

func Router() *echo.Echo {

	e := echo.New()
	// 첫 화면
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// debug mode enable
	e.Debug = true

	// echo middleware func
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// health check
	e.GET("/healthy", func(c echo.Context) error {
		return c.String(http.StatusOK, "healthy!!")
	})

	// router list
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/list", list)
	e.POST("/read", handler.ReadList)

	//e.Logger.Fatal(e.Start(":1323")) // localhost:1323

	return e
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