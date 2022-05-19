package main

import (
	"fmt"
	"go-web/initialize"
)

func main() {
	fmt.Println("Hello word")

	initialize.InitDataBase()

	engine := initialize.Routers()

	_ = engine.Run()

}
