package main

import (
	"github.com/luizfpsoares/gerador-de-certificados-ssl/service"
)

func main() {
	dir := service.GenDir()
	ca, caKey := service.GenCA("ca-v1", 365, "vpn.debugsystem.com.br", dir)
	service.GenServerCert("server", "vpn.debugsystem.com.br", ca, caKey, 365, dir)
	service.GenClientCert("luiz.fernando", "vpn.debugsystem.com.br", ca, caKey, 365, dir)

}
