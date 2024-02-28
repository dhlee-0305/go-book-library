package read

import (
	"net/http"

	"github.com/labstack/echo"
)

func ReadList(c echo.Context) error {
	team := c.FormValue("team")
	member := c.FormValue("member")
	return c.String(http.StatusOK, "team:"+team+", memhber:"+member)

}
