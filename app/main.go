package main

import "app/controller"

func main() {

	r := controller.ServerSetup()

	r.Run()
}
