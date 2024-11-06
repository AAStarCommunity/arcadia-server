package main

import "arcadia/routers"

func main() {
	routers.SetRouters().Run(":8080")
}
