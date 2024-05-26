package account

import (
	"github.com/luanruisong/pionex/api"
	"net/http"
)

type (
	Balance struct {
		Coin   string `json:"coin"`
		Free   string `json:"free"`
		Frozen string `json:"frozen"`
	}
	BalancesRes struct {
		Balances []Balance `json:"balances"`
	}
)

var (
	balancesInfo = &api.Api[any, *BalancesRes]{
		Method: http.MethodGet,
		Path:   "/api/v1/account/balances",
	}
)

// BalancesInfo https://pionex-doc.gitbook.io/apidocs/restful/account/get-balance
func BalancesInfo() (*api.Ret[*BalancesRes], error) {
	return balancesInfo.Do(nil)
}
