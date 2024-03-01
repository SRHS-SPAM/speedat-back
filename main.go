package main

import "speedat-back/controllers"

func main() {
	controllers.NewController(":9000")
}
