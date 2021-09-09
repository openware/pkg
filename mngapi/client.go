package mngapi

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	jwtgo "github.com/golang-jwt/jwt"
)

const (
	// RequestTimeout default value to 30 seconds
	RequestTimeout = time.Duration(30 * time.Second)

	// JWTExpireDuration default value to 1 hour
	JWTExpireDuration = time.Hour

	//JWTAlgorithm default value to RS256
	JWTAlgorithm = "RS256"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// DefaultClient interface
type DefaultClient interface {
	Request(method string, path string, body interface{}) ([]byte, *APIError)
}

// Client instance
type Client struct {
	URL              string
	jwtIssuer        string
	jwtSigningMethod jwtgo.SigningMethod
	jwtPrivateKey    *rsa.PrivateKey
	httpClient       HTTPClient
}

// APIError response from management API
type APIError struct {
	StatusCode int      `json:"code"`
	Error      string   `json:"error,omitempty"`
	Errors     []string `json:"errors,omitempty"`
}

// New to return Client struct
func New(URL string, jwtIssuer string, jwtAlgo string, jwtPrivateKey string) (*Client, error) {
	pk, err := loadPrivateKeyFromString(jwtPrivateKey)
	if err != nil {
		return nil, err
	}

	if jwtAlgo == "" {
		jwtAlgo = JWTAlgorithm
	}

	sm := jwtgo.GetSigningMethod(jwtAlgo)
	if sm == nil {
		return nil, fmt.Errorf("Unsupported signing method %s", jwtAlgo)
	}

	if jwtIssuer == "" {
		return nil, fmt.Errorf("JWT issuer unset")
	}

	return &Client{
		httpClient:       &http.Client{Timeout: RequestTimeout},
		URL:              URL,
		jwtIssuer:        jwtIssuer,
		jwtSigningMethod: sm,
		jwtPrivateKey:    pk,
	}, nil
}

// Request to call HTTP request
func (m *Client) Request(method string, path string, body interface{}) ([]byte, *APIError) {
	// Check for allowed HTTP methods
	if !allowedHTTPMethods(method) {
		return nil, &APIError{StatusCode: 500, Error: "HTTP method is not allowed, accept only POST and PUT"}
	}

	url, err := url.Parse(m.URL)
	url.Path = filepath.Join(url.Path, path)

	// TODO: Add to support JWT with multiple signatures
	// Generate JWT
	jwt, err := m.generateJWT(convertToStringInterface(body), JWTExpireDuration)
	if err != nil {
		return nil, &APIError{StatusCode: 500, Error: err.Error()}
	}

	// Convert jwt to json string
	jwtstr, err := json.Marshal(jwt)
	if err != nil {
		return nil, &APIError{StatusCode: 500, Error: err.Error()}
	}

	// Create new HTTP request
	req, err := http.NewRequest(method, url.String(), bytes.NewBuffer(jwtstr))
	if err != nil {
		return nil, &APIError{StatusCode: 500, Error: err.Error()}
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// Call HTTP request
	res, err := m.httpClient.Do(req)
	if err != nil {
		return nil, &APIError{StatusCode: 500, Error: err.Error()}
	}

	defer res.Body.Close()

	// Convert response body to []byte
	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, &APIError{StatusCode: 500, Error: err.Error()}
	}

	// Check for API error
	if !(res.StatusCode == 200 || res.StatusCode == 201) {
		apiError := APIError{
			StatusCode: res.StatusCode,
			Error:      res.Status,
		}

		_ = json.Unmarshal(resbody, &apiError)

		return nil, &apiError
	}

	return resbody, nil
}

func (m *Client) generateJWT(data map[string]interface{}, validPeriod time.Duration, opts ...interface{}) (map[string]interface{}, error) {
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

func allowedHTTPMethods(method string) bool {
	if len(method) == 0 {
		return false
	}

	var methods = []string{http.MethodPost, http.MethodPut}

	for _, v := range methods {
		if v == method {
			return true
		}
	}

	return false
}

func convertToStringInterface(input interface{}) map[string]interface{} {
	var mapinterface map[string]interface{}
	str, _ := json.Marshal(input)
	json.Unmarshal(str, &mapinterface)

	return mapinterface
}
