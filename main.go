package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	dir := genDir()
	ca, caKey := GenCA("ca-v1", 365, "vpn.debugsystem.com.br", dir)
	genServerCert("server", "vpn.debugsystem.com.br", ca, caKey, 365, dir)
	genClientCert("luiz.fernando", "vpn.debugsystem.com.br", ca, caKey, 365, dir)

}

func genDir() string {
	prefix := "tmp-"

	dir, err := os.MkdirTemp("", prefix)
	if err != nil {
		fmt.Println("Erro ao criar diretório", err)
		return ""
	}
	fmt.Println("Diretório temporario criado: ", dir)
	return dir
}

func GenCA(caName string, expirationTime int, domain string, dir string) (string, string) {
	res := exec.Command("openssl", "genrsa", "-out", dir+"/"+caName+".key", "2048")
	_, err := res.Output()
	if err != nil {
		fmt.Println("Erro ao gerar chave para CA")
		return "", ""
	}
	fmt.Println("CA key gerada com sucesso...")

	res = exec.Command(
		"openssl", "req", "-x509", "-new", "-nodes",
		"-key", dir+"/"+caName+".key",
		"-sha256",
		"-days", strconv.Itoa(expirationTime),
		"-out", dir+"/"+caName+".crt",
		"-subj", "/CN="+domain,
	)
	_, err = res.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao gerar CA")
		return "", ""
	}
	fmt.Println("CA gerada com sucesso...")
	ca, _ := exec.Command("cat", dir+"/"+caName+".crt").Output()
	caKey, _ := exec.Command("cat", dir+"/"+caName+".key").Output()
	return string(ca), string(caKey)

}

func genServerCert(crtName string, domain string, ca string, caKey string, expirationTime int, dir string) (string, string, string) {
	caFile, err := os.Create("ca.crt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo da CA")
		return "", "", ""
	}
	defer caFile.Close()
	_, err = caFile.WriteString(ca)
	if err != nil {
		fmt.Println("Erro ao escrever conteúdo da CA")
		return "", "", ""
	}
	fmt.Println("Ca importada com sucesso!")

	caKeyFile, err := os.Create("ca.key")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo de chave da CA")
		return "", "", ""
	}
	defer caKeyFile.Close()
	_, err = caKeyFile.WriteString(caKey)
	if err != nil {
		fmt.Println("Erro ao escrever conteúdo da chave da CA")
		return "", "", ""
	}
	fmt.Println("Ca key importada com sucesso!")

	res := exec.Command("openssl", "genrsa", "-out", dir+"/"+crtName+".key", "2048")
	_, err = res.Output()
	if err != nil {
		fmt.Println("Erro ao gerar chave para CA")
		return "", "", ""
	}
	fmt.Println("CA key gerada com sucesso...")

	res = exec.Command(
		"openssl", "req", "-new", "-key",
		dir+"/"+crtName+".key",
		"-out", dir+"/"+crtName+".csr",
		"-subj", "/CN="+domain,
	)
	_, err = res.Output()
	if err != nil {
		fmt.Println("Erro ao Gerar CSR do Servidor")
		return "", "", ""
	}
	fmt.Println("CSR Gerado com sucesso!")

	res = exec.Command(
		"openssl", "x509", "-req", "-in",
		dir+"/"+crtName+".csr",
		"-CA", "ca.crt",
		"-CAkey", "ca.key",
		"-CAcreateserial",
		"-out", dir+"/"+crtName+".crt",
		"-days", strconv.Itoa(expirationTime),
		"-sha256",
	)
	_, err = res.Output()
	if err != nil {
		fmt.Println("Falha ao gerar certificado do servidor")
		return "", "", ""
	}
	fmt.Println("Certificado de servidor gerado com sucesso!")

	serverCsr, _ := exec.Command("cat", dir+"/"+crtName+".csr").Output()
	server, _ := exec.Command("cat", dir+"/"+crtName+".crt").Output()
	serverKey, _ := exec.Command("cat", dir+"/"+crtName+".key").Output()

	_, err = exec.Command("rm", "-rf", "ca.crt").Output()
	if err != nil {
		fmt.Println("Erro ao deletar CA temporaria: ", err)
	}
	_, err = exec.Command("rm", "-rf", "ca.key").Output()
	if err != nil {
		fmt.Println("Erro ao deletar CA Key temporaria: ", err)
	}

	return string(serverCsr), string(server), string(serverKey)
}

func genClientCert(crtName string, domain string, ca string, caKey string, expirationTime int, dir string) (string, string, string) {
	caFile, err := os.Create("ca.crt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo da CA")
		return "", "", ""
	}
	defer caFile.Close()
	_, err = caFile.WriteString(ca)
	if err != nil {
		fmt.Println("Erro ao escrever conteúdo da CA")
		return "", "", ""
	}
	fmt.Println("Ca importada com sucesso!")

	caKeyFile, err := os.Create("ca.key")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo de chave da CA")
		return "", "", ""
	}
	defer caKeyFile.Close()
	_, err = caKeyFile.WriteString(caKey)
	if err != nil {
		fmt.Println("Erro ao escrever conteúdo da chave da CA")
		return "", "", ""
	}
	fmt.Println("Ca key importada com sucesso!")

	res := exec.Command("openssl", "genrsa", "-out", dir+"/"+crtName+".key", "2048")
	_, err = res.Output()
	if err != nil {
		fmt.Println("Erro ao gerar chave para CA")
		return "", "", ""
	}
	fmt.Println("CA key gerada com sucesso...")

	res = exec.Command(
		"openssl", "req", "-new", "-key",
		dir+"/"+crtName+".key",
		"-out", dir+"/"+crtName+".csr",
		"-subj", "/CN="+domain,
	)
	_, err = res.Output()
	if err != nil {
		fmt.Println("Erro ao Gerar CSR do Cliente")
		return "", "", ""
	}
	fmt.Println("CSR Gerado com sucesso!")

	res = exec.Command(
		"openssl", "x509", "-req", "-in",
		dir+"/"+crtName+".csr",
		"-CA", "ca.crt",
		"-CAkey", "ca.key",
		"-CAcreateserial",
		"-out", dir+"/"+crtName+".crt",
		"-days", strconv.Itoa(expirationTime),
		"-sha256",
	)
	_, err = res.Output()
	if err != nil {
		fmt.Println("Falha ao gerar certificado do Cliente")
		return "", "", ""
	}
	fmt.Println("Certificado de Cliente gerado com sucesso!")

	clientCsr, _ := exec.Command("cat", dir+"/"+crtName+".csr").Output()
	client, _ := exec.Command("cat", dir+"/"+crtName+".crt").Output()
	clientKey, _ := exec.Command("cat", dir+"/"+crtName+".key").Output()

	_, err = exec.Command("rm", "-rf", "ca.crt").Output()
	if err != nil {
		fmt.Println("Erro ao deletar CA temporaria: ", err)
	}
	_, err = exec.Command("rm", "-rf", "ca.key").Output()
	if err != nil {
		fmt.Println("Erro ao deletar CA Key temporaria: ", err)
	}

	return string(clientCsr), string(client), string(clientKey)
}
