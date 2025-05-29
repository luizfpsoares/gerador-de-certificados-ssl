package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.Run("0.0.0.0:8080")
}
