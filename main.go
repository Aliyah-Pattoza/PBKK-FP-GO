package main

import (
	"pbkk-fp-go/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()

	r.Run()

}
