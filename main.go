package main

import "github.com/similadayo/imageRecog/routes"

func main() {
	//importing the routes and server created in our routes folder/package.
	routes.CreateRouter()
	routes.InitializeRoutes()
	routes.ServerStart()
}
