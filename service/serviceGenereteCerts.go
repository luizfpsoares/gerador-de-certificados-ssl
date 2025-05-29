package service

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/luizfpsoares/gerador-de-certificados-ssl/model"
)

func GenDir() string {
	prefix := "tmp-"

	dir, err := os.MkdirTemp("", prefix)
	if err != nil {
		fmt.Println("Erro ao criar diretório", err)
		return ""
	}
	fmt.Println("Diretório temporario criado: ", dir)
	return dir
}

func GenCA(genCA model.GenCA) (model.ResCa, error) {
	res := exec.Command("openssl", "genrsa", "-out", genCA.Dir+"/"+genCA.Name+".key", "2048")
	_, err := res.Output()
	if err != nil {
		fmt.Println("Erro ao gerar chave para CA: ", err)
		return model.ResCa{}, err
	}
	fmt.Println("CA key gerada com sucesso...")

	res = exec.Command(
		"openssl", "req", "-x509", "-new", "-nodes",
		"-key", genCA.Dir+"/"+genCA.Name+".key",
		"-sha256",
		"-days", strconv.Itoa(genCA.ExpirationTime),
		"-out", genCA.Dir+"/"+genCA.Name+".crt",
		"-subj", "/CN="+genCA.Domain,
	)
	_, err = res.CombinedOutput()
	if err != nil {
		fmt.Println("Erro ao gerar CA", err)
		return model.ResCa{}, err
	}
	fmt.Println("CA gerada com sucesso...")

	caContent, _ := exec.Command("cat", genCA.Dir+"/"+genCA.Name+".crt").Output()
	caKeyContent, _ := exec.Command("cat", genCA.Dir+"/"+genCA.Name+".key").Output()

	resCa := model.ResCa{
		Ca:    string(caContent),
		CaKey: string(caKeyContent),
	}

	_, err = exec.Command("rm", "-rf", "ca.srl").Output()
	if err != nil {
		fmt.Println("Erro ao deletar lixo: ", err)
	}

	return resCa, nil

}

func GenServerCert(genServerCert model.GenServerCert) (model.ResServer, error) {
	caFile, err := os.Create("ca.crt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo da CA")
		return model.ResServer{}, err
	}
	defer caFile.Close()
	_, err = caFile.WriteString(genServerCert.CaContent)
	if err != nil {
		fmt.Println("Erro ao escrever conteúdo da CA")
		return model.ResServer{}, err
	}
	fmt.Println("Ca importada com sucesso!")

	caKeyFile, err := os.Create("ca.key")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo de chave da CA")
		return model.ResServer{}, err
	}
	defer caKeyFile.Close()
	_, err = caKeyFile.WriteString(genServerCert.CaKeyContent)
	if err != nil {
		fmt.Println("Erro ao escrever conteúdo da chave da CA")
		return model.ResServer{}, err
	}
	fmt.Println("Ca key importada com sucesso!")

	res := exec.Command("openssl", "genrsa", "-out", genServerCert.Dir+"/"+genServerCert.Name+".key", "2048")
	_, err = res.Output()
	if err != nil {
		fmt.Println("Erro ao gerar chave para CA")
		return model.ResServer{}, err
	}
	fmt.Println("CA key gerada com sucesso...")

	res = exec.Command(
		"openssl", "req", "-new", "-key",
		genServerCert.Dir+"/"+genServerCert.Name+".key",
		"-out", genServerCert.Dir+"/"+genServerCert.Name+".csr",
		"-subj", "/CN="+genServerCert.Domain,
	)
	_, err = res.Output()
	if err != nil {
		fmt.Println("Erro ao Gerar CSR do Servidor")
		return model.ResServer{}, err
	}
	fmt.Println("CSR Gerado com sucesso!")

	res = exec.Command(
		"openssl", "x509", "-req", "-in",
		genServerCert.Dir+"/"+genServerCert.Name+".csr",
		"-CA", "ca.crt",
		"-CAkey", "ca.key",
		"-CAcreateserial",
		"-out", genServerCert.Dir+"/"+genServerCert.Name+".crt",
		"-days", strconv.Itoa(genServerCert.ExpirationTime),
		"-sha256",
	)
	_, err = res.Output()
	if err != nil {
		fmt.Println("Falha ao gerar certificado do servidor")
		return model.ResServer{}, err
	}
	fmt.Println("Certificado de servidor gerado com sucesso!")

	serverCsr, _ := exec.Command("cat", genServerCert.Dir+"/"+genServerCert.Name+".csr").Output()
	server, _ := exec.Command("cat", genServerCert.Dir+"/"+genServerCert.Name+".crt").Output()
	serverKey, _ := exec.Command("cat", genServerCert.Dir+"/"+genServerCert.Name+".key").Output()

	_, err = exec.Command("rm", "-rf", "ca.crt").Output()
	if err != nil {
		fmt.Println("Erro ao deletar CA temporaria: ", err)
	}
	_, err = exec.Command("rm", "-rf", "ca.key").Output()
	if err != nil {
		fmt.Println("Erro ao deletar CA Key temporaria: ", err)
	}

	resServer := model.ResServer{
		Server:    string(server),
		ServerKey: string(serverKey),
		ServerCsr: string(serverCsr),
	}

	return resServer, nil
}

func GenClientCert(genClientCert model.GenClientCert) (model.ResClient, error) {
	caFile, err := os.Create("ca.crt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo da CA")
		return model.ResClient{}, err
	}
	defer caFile.Close()
	_, err = caFile.WriteString(genClientCert.CaContent)
	if err != nil {
		fmt.Println("Erro ao escrever conteúdo da CA")
		return model.ResClient{}, err
	}
	fmt.Println("Ca importada com sucesso!")

	caKeyFile, err := os.Create("ca.key")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo de chave da CA")
		return model.ResClient{}, err
	}
	defer caKeyFile.Close()
	_, err = caKeyFile.WriteString(genClientCert.CaKeyContent)
	if err != nil {
		fmt.Println("Erro ao escrever conteúdo da chave da CA")
		return model.ResClient{}, err
	}
	fmt.Println("Ca key importada com sucesso!")

	res := exec.Command("openssl", "genrsa", "-out", genClientCert.Dir+"/"+genClientCert.Name+".key", "2048")
	_, err = res.Output()
	if err != nil {
		fmt.Println("Erro ao gerar chave para CA")
		return model.ResClient{}, err
	}
	fmt.Println("CA key gerada com sucesso...")

	res = exec.Command(
		"openssl", "req", "-new", "-key",
		genClientCert.Dir+"/"+genClientCert.Name+".key",
		"-out", genClientCert.Dir+"/"+genClientCert.Name+".csr",
		"-subj", "/CN="+genClientCert.Domain,
	)
	_, err = res.Output()
	if err != nil {
		fmt.Println("Erro ao Gerar CSR do Cliente")
		return model.ResClient{}, err
	}
	fmt.Println("CSR Gerado com sucesso!")

	res = exec.Command(
		"openssl", "x509", "-req", "-in",
		genClientCert.Dir+"/"+genClientCert.Name+".csr",
		"-CA", "ca.crt",
		"-CAkey", "ca.key",
		"-CAcreateserial",
		"-out", genClientCert.Dir+"/"+genClientCert.Name+".crt",
		"-days", strconv.Itoa(genClientCert.ExpirationTime),
		"-sha256",
	)
	_, err = res.Output()
	if err != nil {
		fmt.Println("Falha ao gerar certificado do Cliente")
		return model.ResClient{}, err
	}
	fmt.Println("Certificado de Cliente gerado com sucesso!")

	clientCsr, _ := exec.Command("cat", genClientCert.Dir+"/"+genClientCert.Name+".csr").Output()
	client, _ := exec.Command("cat", genClientCert.Dir+"/"+genClientCert.Name+".crt").Output()
	clientKey, _ := exec.Command("cat", genClientCert.Dir+"/"+genClientCert.Name+".key").Output()

	_, err = exec.Command("rm", "-rf", "ca.crt").Output()
	if err != nil {
		fmt.Println("Erro ao deletar CA temporaria: ", err)
	}
	_, err = exec.Command("rm", "-rf", "ca.key").Output()
	if err != nil {
		fmt.Println("Erro ao deletar CA Key temporaria: ", err)
	}

	resClient := model.ResClient{
		Client:    string(client),
		ClientKey: string(clientKey),
		ClientCsr: string(clientCsr),
	}

	return resClient, nil
}
