package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luizfpsoares/gerador-de-certificados-ssl/handler"
)

func main() {
	router := gin.Default()
	router.POST("/api/v1/ca", handler.PostGenCa)
	router.POST("/api/v1/server", handler.PostGenServer)
	router.POST("/api/v1/client", handler.PostGenClient)
	router.POST("/api/v1/chain", handler.PostGenChain)

	router.Run("0.0.0.0:8080")
}
