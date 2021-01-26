package mgntapi

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

const (
	// RequestTimeout default value
	RequestTimeout = time.Duration(5 * time.Second)
)

// ManagementAPIV2 struct to define common data and function
type ManagementAPIV2 struct {
	rootAPIUrl     string
	endpointPrefix string
	client         *http.Client
	keychain       *KeychainData
}

// New to return ManagementAPIV2 struct
func New(rootAPIUrl string, endpointPrefix string, keychain *KeychainData) *ManagementAPIV2 {
	return &ManagementAPIV2{
		rootAPIUrl:     rootAPIUrl,
		endpointPrefix: endpointPrefix,
		client:         &http.Client{Timeout: RequestTimeout},
		keychain:       keychain,
	}
}

func (m *ManagementAPIV2) Request(method string, path string, body []byte) ([]byte, error) {
	url, err := url.Parse(m.rootAPIUrl)
	url.Path = path.Join(url.Path, m.endpointPrefix, path)
	req, err := http.NewRequest(method, url.String(), bytes.NewBuffer(body))
	if err != nil {
		log.Fatalln("Request", "Can not create new request: "+err.Error())
		return nil, err
	}

	privateKey, err := loadPrivateKeyFromString(m.keychain.Value)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	tokenString, err := generateToken(body, privateKey)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tokenString)

	res, err := m.client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return body, nil
}

func generateToken(data interface{}, key *rsa.PrivateKey) (string, error) {
	claims := jwtgo.MapClaims{
		"data": data,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().UTC().Add(time.Hour).Unix(),
		"jit":  strconv.FormatInt(time.Now().Unix(), 10),
		"iss":  "opendax",
	}

	t := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claims)

	return t.SignedString(key)
}

func loadPrivateKeyFromString(str string) (*rsa.PrivateKey, error) {
	pem, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	key, err := jwtgo.ParseRSAPrivateKeyFromPEM(pem)
	if err != nil {
		return nil, err
	}

	return key, nil
}
