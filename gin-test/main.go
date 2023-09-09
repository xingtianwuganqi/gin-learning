package main

import (
	"fmt"
	"gin-test/db"
	"gin-test/routers"
)

func main() {
	fmt.Println("")

	db.Init()
	r := routers.RegisterRouter()

	r.Run()
}
