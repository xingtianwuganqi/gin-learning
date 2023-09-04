package main

import (
	"fmt"
	"gin-test/routers"
)

func main() {
	fmt.Println("")
	r := routers.RegisterRouter()

	r.Run()
}
