package main

import "assignment-project/routers"

var PORT = ":8080"

func main() {
	routers.StartServer().Run(PORT)
}
