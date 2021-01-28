package mgntapi

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
)

const (
	// RequestTimeout default value
	RequestTimeout = time.Duration(5 * time.Second)
)

// Client struct to define common data and function
type Client struct {
	rootAPIUrl       string
	endpointPrefix   string
	client           *http.Client
	jwtIssuer        string
	jwtSigningMethod jwtgo.SigningMethod
	jwtPrivateKey    *rsa.PrivateKey
}

// New to return ManagementAPIV2 struct
func New(rootAPIUrl string, endpointPrefix string, jwtIssuer string, jwtAlgo string, jwtPrivateKey string) (*Client, error) {
	pk, err := loadPrivateKeyFromString(jwtPrivateKey)
	if err != nil {
		return nil, err
	}

	if jwtAlgo == "" {
		jwtAlgo = "RS256"
	}

	sm := jwtgo.GetSigningMethod(jwtAlgo)
	if sm == nil {
		return nil, fmt.Errorf("Unsupported signing method %s", jwtAlgo)
	}

	if jwtIssuer == "" {
		return nil, fmt.Errorf("JWT issuer unset")
	}

	return &Client{
		rootAPIUrl:       rootAPIUrl,
		endpointPrefix:   endpointPrefix,
		client:           &http.Client{Timeout: RequestTimeout},
		jwtIssuer:        jwtIssuer,
		jwtSigningMethod: sm,
		jwtPrivateKey:    pk,
	}, nil
}

func (m *Client) Request(method string, path string, body []byte) ([]byte, error) {
	url, err := url.Parse(m.rootAPIUrl)
	url.Path = filepath.Join(url.Path, m.endpointPrefix, path)
	req, err := http.NewRequest(method, url.String(), bytes.NewBuffer(body))
	if err != nil {
		log.Fatalln("Request", "Can not create new request: "+err.Error())
		return nil, err
	}

	// TODO: Add JWS

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// TODO: Uncomment later
	// req.Header.Add("Authorization", "Bearer "+tokenString)

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

func (m *Client) generateJWT(data interface{}, validPeriod time.Duration, opts ...interface{}) (map[string]interface{}, error) {
	iat := time.Now()
	jti := RandomString(16)
	if len(opts) > 0 {
		iat = time.Unix(opts[0].(int64), 0)
	}
	if len(opts) > 1 {
		jti = opts[1].(string)
	}
	claims := jwtgo.MapClaims{
		"data": data,
		"iat":  iat.Unix(),
		"exp":  iat.Add(validPeriod).Unix(),
		"iss":  m.jwtIssuer,
		"jti":  jti,
	}

	t := jwtgo.NewWithClaims(m.jwtSigningMethod, claims)

	sstr, err := t.SigningString()
	if err != nil {
		return nil, err
	}

	hp := strings.Split(sstr, ".")
	if len(hp) != 2 {
		return nil, fmt.Errorf("Invalid segment count in sstr %d, expected 2", len(hp))
	}

	sig, err := t.Method.Sign(sstr, m.jwtPrivateKey)
	if err != nil {
		return nil, err
	}

	jwt := map[string]interface{}{
		"payload": hp[1],
		"signatures": []map[string]interface{}{
			{
				"protected": hp[0],
				"header":    map[string]string{"kid": m.jwtIssuer},
				"signature": sig,
			},
		},
	}

	return jwt, nil
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
