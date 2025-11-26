package account

import (
	"net/http"

	"github.com/luanruisong/pionex/api"
	"github.com/valyala/fasthttp"
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
		balancesInfo *api.Api[any, *BalancesRes]
	}
)

// BalancesInfo https://pionex-doc.gitbook.io/apidocs/restful/account/get-balance
func (a *Account) BalancesInfo() (*api.Ret[*BalancesRes], error) {
	return a.balancesInfo.Do(nil)
}

func NewAccount(s *api.Singer, c *fasthttp.Client) *Account {
	return &Account{
		balancesInfo: api.NewApi[any, *BalancesRes](
			http.MethodGet,
			"/api/v1/account/balances",
			api.WithSigner(s),
			api.WithClient(c),
		),
	}
}
