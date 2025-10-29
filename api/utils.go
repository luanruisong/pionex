package api

import (
	"net/http"

	"github.com/go-playground/form/v4"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

const (
	TsKey   = "timestamp"
	SignKey = "PIONEX-SIGNATURE"
	ApiKey  = "PIONEX-KEY"
)

var (
	Host        = "https://api.pionex.com"
	formEncoder *form.Encoder
)

func init() {
	formEncoder = form.NewEncoder()
}

func (a *Api[Req, Res]) preParseRequest(req *fasthttp.Request, data any) {
	req.Header.SetMethod(a.Method)
	req.SetRequestURI(Host)
	u := req.URI()
	u.SetPath(a.Path)
	switch a.Method {
	case http.MethodGet:
		q, _ := formEncoder.Encode(data)
		if len(q) > 0 {
			u.SetQueryString(q.Encode())
		}
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		b, _ := jsoniter.Marshal(data)
		req.SetBody(b)
	}
}
