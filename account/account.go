package account

import (
	"github.com/luanruisong/pionex/api"
	"github.com/valyala/fasthttp"
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

	Account struct {
		s *api.Singer
		c *fasthttp.Client
	}
)

var (
	balancesInfo = &api.Api[any, *BalancesRes]{
		Method: http.MethodGet,
		Path:   "/api/v1/account/balances",
	}
)

// BalancesInfo https://pionex-doc.gitbook.io/apidocs/restful/account/get-balance
func (a *Account) BalancesInfo() (*api.Ret[*BalancesRes], error) {
	return balancesInfo.Do(nil, a.s, a.c)
}

func NewAccount(s *api.Singer, c *fasthttp.Client) *Account {
	return &Account{
		s: s,
		c: c,
	}
}
