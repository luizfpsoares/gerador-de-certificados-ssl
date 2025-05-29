package model

type GenCA struct {
	Name           string `json:"name"`
	ExpirationTime int    `json:"expiration_time"`
	Domain         string `json:"domain"`
	Dir            string `json:"dir"`
}

type GenServerCert struct {
	Name           string `json:"name"`
	Domain         string `json:"domain"`
	CaContent      string `json:"ca_content"`
	CaKeyContent   string `json:"ca_key_content"`
	ExpirationTime int    `json:"expiration_time"`
	Dir            string `json:"dir"`
}

type GenClientCert struct {
	Name           string `json:"name"`
	Domain         string `json:"domain"`
	CaContent      string `json:"ca_content"`
	CaKeyContent   string `json:"ca_key_content"`
	ExpirationTime int    `json:"expiration_time"`
	Dir            string `json:"dir"`
}

type GenChain struct {
	CaName               string `json:"ca_name"`
	ServerName           string `json:"server_name"`
	ClientName           string `json:"client_name"`
	Domain               string `json:"domain"`
	CaExpirationTime     int    `json:"ca_expiration_time"`
	ServerExpirationTime int    `json:"server_expiration_time"`
	ClientExpirationTime int    `json:"client_expiration_time"`
	Dir                  string `json:"dir"`
}

type ResCa struct {
	Ca    string `json:"ca"`
	CaKey string `json:"ca_key"`
}

type ResServer struct {
	Server    string `json:"server"`
	ServerKey string `json:"server_key"`
	ServerCsr string `json:"server_csr"`
}

type ResClient struct {
	Client    string `json:"client"`
	ClientKey string `json:"client_key"`
	ClientCsr string `json:"client_csr"`
}

type ResChain struct {
	Ca        string `json:"ca"`
	CaKey     string `json:"ca_key"`
	ServerCsr string `json:"server_csr"`
	Server    string `json:"server"`
	ServerKey string `json:"server_key"`
	ClientCsr string `json:"client_csr"`
	Client    string `json:"client"`
	ClientKey string `json:"client_key"`
}
