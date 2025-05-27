package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	ca, caKey := GenCA("ca-v1", 365, "vpn.debugsystem.com.br")
	genServerCert("server", "vpn.debugsystem.com.br", ca, caKey, 365)
	genClientCert("luiz.fernando", "vpn.debugsystem.com.br", ca, caKey, 365)

}

func GenCA(caName string, expirationTime int, domain string) (string, string) {
	res := exec.Command("openssl", "genrsa", "-out", caName+".key", "2048")
	_, err := res.Output()
	if err != nil {
		fmt.Println("Erro ao gerar chave para CA")
		return "", ""
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
		return "", ""
	}
	fmt.Println("CA gerada com sucesso...")
	ca, _ := exec.Command("cat", caName+".crt").Output()
	caKey, _ := exec.Command("cat", caName+".key").Output()
	return string(ca), string(caKey)

}

func genServerCert(crtName string, domain string, ca string, caKey string, expirationTime int) {
	caFile, err := os.Create("ca.crt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo da CA")
		return
	}
	defer caFile.Close()
	_, err = caFile.WriteString(ca)
	if err != nil {
		fmt.Println("Erro ao escrever conteúdo da CA")
		return
	}
	fmt.Println("Ca importada com sucesso!")

	caKeyFile, err := os.Create("ca.key")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo de chave da CA")
		return
	}
	defer caKeyFile.Close()
	_, err = caKeyFile.WriteString(caKey)
	if err != nil {
		fmt.Println("Erro ao escrever conteúdo da chave da CA")
		return
	}
	fmt.Println("Ca key importada com sucesso!")

	res := exec.Command("openssl", "genrsa", "-out", crtName+".key", "2048")
	_, err = res.Output()
	if err != nil {
		fmt.Println("Erro ao gerar chave para CA")
		return
	}
	fmt.Println("CA key gerada com sucesso...")

	res = exec.Command(
		"openssl", "req", "-new", "-key",
		crtName+".key",
		"-out", crtName+".csr",
		"-subj", "/CN="+domain,
	)
	_, err = res.Output()
	if err != nil {
		fmt.Println("Erro ao Gerar CSR do Servidor")
		return
	}
	fmt.Println("CSR Gerado com sucesso!")

	res = exec.Command(
		"openssl", "x509", "-req", "-in",
		crtName+".csr",
		"-CA", "ca.crt",
		"-CAkey", "ca.key",
		"-CAcreateserial",
		"-out", crtName+".crt",
		"-days", strconv.Itoa(expirationTime),
		"-sha256",
	)
	_, err = res.Output()
	if err != nil {
		fmt.Println("Falha ao gerar certificado do servidor")
		return
	}
	fmt.Println("Certificado de servidor gerado com sucesso!")
	exec.Command("rm", "-rf", "ca.crt")
	exec.Command("rm", "-rf", "ca.key")
}

func genClientCert(crtName string, domain string, ca string, caKey string, expirationTime int) {
	caFile, err := os.Create("ca.crt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo da CA")
		return
	}
	defer caFile.Close()
	_, err = caFile.WriteString(ca)
	if err != nil {
		fmt.Println("Erro ao escrever conteúdo da CA")
		return
	}
	fmt.Println("Ca importada com sucesso!")

	caKeyFile, err := os.Create("ca.key")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo de chave da CA")
		return
	}
	defer caKeyFile.Close()
	_, err = caKeyFile.WriteString(caKey)
	if err != nil {
		fmt.Println("Erro ao escrever conteúdo da chave da CA")
		return
	}
	fmt.Println("Ca key importada com sucesso!")

	res := exec.Command("openssl", "genrsa", "-out", crtName+".key", "2048")
	_, err = res.Output()
	if err != nil {
		fmt.Println("Erro ao gerar chave para CA")
		return
	}
	fmt.Println("CA key gerada com sucesso...")

	res = exec.Command(
		"openssl", "req", "-new", "-key",
		crtName+".key",
		"-out", crtName+".csr",
		"-subj", "/CN="+domain,
	)
	_, err = res.Output()
	if err != nil {
		fmt.Println("Erro ao Gerar CSR do Cliente")
		return
	}
	fmt.Println("CSR Gerado com sucesso!")

	res = exec.Command(
		"openssl", "x509", "-req", "-in",
		crtName+".csr",
		"-CA", "ca.crt",
		"-CAkey", "ca.key",
		"-CAcreateserial",
		"-out", crtName+".crt",
		"-days", strconv.Itoa(expirationTime),
		"-sha256",
	)
	_, err = res.Output()
	if err != nil {
		fmt.Println("Falha ao gerar certificado do Cliente")
		return
	}
	fmt.Println("Certificado de Cliente gerado com sucesso!")
	exec.Command("rm", "-rf", "ca.crt")
	exec.Command("rm", "-rf", "ca.key")
}
