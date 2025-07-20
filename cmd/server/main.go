package main

import (
	"github.com/3xhvy/go-backend/internal/router"
)

func main() {
	router := router.NewServer()
	router.Run(":8002") // listen and serve on 0.0.0.0:8080
}
