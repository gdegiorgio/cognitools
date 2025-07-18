package service

import (
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Auth interface {
	GenerateJWT(domain string, clientId string, clientSecret string, scope string) (string, error)
}

type authservice struct{}

func NewAuthService() Auth {
	return &authservice{}
}

func (a *authservice) GenerateJWT(domain string, clientId string, clientSecret string, scope string) (string, error) {

	c := http.DefaultClient

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", scope)
	data.Set("client_id", clientId)

	req, err := http.NewRequest("POST", "https://"+domain+"/oauth2/token", strings.NewReader(data.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientId+":"+clientSecret)))

	if err != nil {
		return "", err
	}

	res, err := c.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close() //nolint:errcheck

	if res.StatusCode != http.StatusOK {
		return "", err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}
	return string(body), nil
}
