package barong

import (
	"github.com/openware/pkg/mngapi"
)

type BarongMngAPIV2 struct {
	cli *mngapi.Client
}

func New(cli *mngapi.Client) *BarongMngAPIV2 {
	return &BarongMngAPIV2{
		cli: cli,
	}
}
