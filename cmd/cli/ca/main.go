package main

import (
	"fmt"

	"github.com/luizfpsoares/gerador-de-certificados-ssl/model"
	"github.com/luizfpsoares/gerador-de-certificados-ssl/service"
)

func main() {

	genCa := model.GenCA{
		Name:           "ca-v2",
		ExpirationTime: 365,
		Domain:         "vpn.example.com",
		Dir:            service.GenDir(),
	}

	caData, _ := service.GenCA(genCa)

	fmt.Println(caData)

}
