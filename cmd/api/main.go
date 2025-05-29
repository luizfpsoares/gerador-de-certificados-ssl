package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luizfpsoares/gerador-de-certificados-ssl/handler"
)

func main() {
	router := gin.Default()
	router.POST("/api/v1/ca", handler.PostGenCa)

	router.Run("0.0.0.0:8080")
}
