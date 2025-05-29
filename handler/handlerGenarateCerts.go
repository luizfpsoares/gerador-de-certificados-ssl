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
	caRes, errCa := service.GenCA(newCa)
	if errCa != nil {
		c.JSON(500, gin.H{"error": errCa})
	}

	c.JSON(201, gin.H{"data": caRes})

}

func PostGenServer(c *gin.Context) {
	var newServer model.GenServerCert

	err := c.BindJSON(&newServer)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	newServer.Dir = service.GenDir()
	serverRes, errServer := service.GenServerCert(newServer)
	if errServer != nil {
		c.JSON(500, gin.H{"error": errServer})
		return
	}

	c.JSON(201, gin.H{"data": serverRes})
}

func PostGenClient(c *gin.Context) {
	var newClient model.GenClientCert

	err := c.BindJSON(&newClient)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
	}

	newClient.Dir = service.GenDir()
	clientRes, errClient := service.GenClientCert(newClient)
	if errClient != nil {
		c.JSON(500, gin.H{"error": errClient})
		return
	}
	c.JSON(201, gin.H{"data": clientRes})
}

func PostGenChain(c *gin.Context) {
	var newChain model.GenChain

	err := c.BindJSON(&newChain)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newChain.Dir = service.GenDir()

	newCa := model.GenCA{
		Name:           newChain.CaName,
		ExpirationTime: newChain.CaExpirationTime,
		Domain:         newChain.Domain,
		Dir:            newChain.Dir,
	}

	resCa, errCa := service.GenCA(newCa)
	if errCa != nil {
		c.JSON(500, gin.H{"error": errCa})
		return
	}

	newServer := model.GenServerCert{
		Name:           newChain.ServerName,
		Domain:         newChain.Domain,
		CaContent:      resCa.Ca,
		CaKeyContent:   resCa.CaKey,
		ExpirationTime: newChain.ServerExpirationTime,
		Dir:            newChain.Dir,
	}
	resServer, errServer := service.GenServerCert(newServer)
	if errServer != nil {
		c.JSON(500, gin.H{"error": errServer})
		return
	}

	newClient := model.GenClientCert{
		Name:           newChain.ClientName,
		Domain:         newChain.Domain,
		CaContent:      resCa.Ca,
		CaKeyContent:   resCa.CaKey,
		ExpirationTime: newChain.ClientExpirationTime,
		Dir:            newChain.Dir,
	}

	resClient, errClient := service.GenClientCert(newClient)
	if errClient != nil {
		c.JSON(500, gin.H{"error": errClient})
		return
	}

	resChain := model.ResChain{
		Ca:        resCa.Ca,
		CaKey:     resCa.CaKey,
		ServerCsr: resServer.ServerCsr,
		Server:    resServer.Server,
		ServerKey: resServer.ServerKey,
		ClientCsr: resClient.ClientCsr,
		Client:    resClient.Client,
		ClientKey: resClient.ClientKey,
	}

	c.JSON(201, gin.H{"data": resChain})

}
