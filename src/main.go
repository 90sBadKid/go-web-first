package main

import "go-web/initialize"

func main() {
	//fmt.Println("Hello word")
	//token, _ := jwt.GenerateToken()
	//println(token)
	initialize.InitDataBase()

	engine := initialize.Routers()

	_ = engine.Run()

}
