package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	CollectRoutes(r)
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
