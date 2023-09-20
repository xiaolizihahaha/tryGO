package main

import (
	"fmt"
	"test1/router"
)

func main() {
	router := router.StartRouter()

	err := router.Run("0.0.0.0:3000")
	fmt.Println(err)
}
