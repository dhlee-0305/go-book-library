package main

import (
	"fmt"
	router "router"

	_ "handler/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:1323
// @BasePath /v1
func main() {
	debug := true

	echoR := router.Router()

	fmt.Println("Start echo server.....")

	if debug {
		echoR.Logger.Fatal(echoR.Start(":1323"))
	} else {
		// https server start
		fmt.Println("https server not yet")
	}
}
