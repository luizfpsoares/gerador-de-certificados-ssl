package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

func main() {
	GenCA("ca-v1", 365, "vpn.debugsystem.com.br")
}

func GenCA(caName string, expirationTime int, domain string) {
	res := exec.Command("openssl", "genrsa", "-out", caName+".key", "2048")
	_, err := res.Output()
	if err != nil {
		fmt.Println("Erro ao gerar chave para CA")
		return
	}
	fmt.Println("CA key gerada com sucesso...")

	res = exec.Command(
		"openssl", "req", "-x509", "-new", "-nodes",
		"-key", caName+".key",
		"-sha256",
		"-days", strconv.Itoa(expirationTime),
		"-out", caName+".crt",
		"-subj", "/CN="+domain,
	)
	_, err = res.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao gerar CA")
		return
	}
	fmt.Println("CA gerada com sucesso...")
}
