package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/luizfpsoares/gerador-de-certificados-ssl/model"
	"github.com/luizfpsoares/gerador-de-certificados-ssl/service"
)

func PostGenCa(c *gin.Context) {
	var newCa model.GenCA

	err := c.BindJSON(&newCa)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newCa.Dir = service.GenDir()
	caRes, _ := service.GenCA(newCa)

	c.JSON(201, gin.H{"data": caRes})

}
