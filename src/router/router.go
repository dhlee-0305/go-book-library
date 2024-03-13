package router

import (
	"net/http"

	echo "github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"

	handler "handler"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func Router() *echo.Echo {

	e := echo.New()

	// Swaggo
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// 첫 화면
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// debug mode enable
	e.Debug = true

	// echo middleware func - after route middleware
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

	//============== ROUTER LIST ==================
	// 샘플
	e.GET("/samplePath/:id", handler.SamplePath)
	e.GET("/sampleGet", handler.SampleGet)
	e.POST("/sampleForm", handler.SampleForm)
	e.POST("/sampleJson", handler.SampleJson)

	// 검색 관련 그룹
	searchGroup := e.Group("/search")
	{
		searchGroup.GET("/:bookName", handler.Search)
		searchGroup.GET("/readn", handler.SearchReadnBook)
		searchGroup.GET("/csv/:bookName", handler.SearchCSV)
		searchGroup.GET("/unRead/:userName", handler.SearchBookToRead)
	}

	// 도서 상태 변경 그룹
	holdGroup := e.Group("/hold")
	{
		holdGroup.POST("/discard/reg", handler.DiscardBookReg)
		holdGroup.POST("/sell/reg", handler.SellBookReg)
		holdGroup.POST("/donate/reg", handler.DonateBookReg)
	}

	// 독서 내역 관리 그룹
	readGroup := e.Group("/read")
	{
		readGroup.POST("/reg", handler.ReadBookReg)
	}

	// 데이터 적재 그룹
	loaderGroup := e.Group("/loader")
	{
		loaderGroup.GET("/books", handler.LoadBookFromCSV)
		loaderGroup.GET("/readHist", handler.LoadReadHistoryFromCSV)
		loaderGroup.GET("/discard", handler.LoadDiscardHistoryFromCSV)
	}

	return e
}
