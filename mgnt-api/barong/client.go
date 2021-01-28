package barong

import mgntapi "github.com/openware/pkg/mgnt-api"

type BarongMngAPIV2 struct {
	cli *mgntapi.Client
}

func New(cli *mgntapi.Client) *BarongMngAPIV2 {
	return &BarongMngAPIV2{
		cli: cli,
	}
}
