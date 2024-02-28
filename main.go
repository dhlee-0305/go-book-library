package main

import (
	"fmt"
	router "router"
)

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
