package barong

import (
	"github.com/openware/pkg/mngapi"
)

type BarongMngAPIV2 struct {
	cli *mngapi.Client
}

func New(rootAPIUrl, endpointPrefix, jwtIssuer, jwtPrivateKey string) *BarongMngAPIV2 {
	cli, _ := mngapi.New(rootAPIUrl, endpointPrefix, jwtIssuer, "RS256", jwtPrivateKey)

	return &BarongMngAPIV2{
		cli: cli,
	}
}
