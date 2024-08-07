package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

type (
	Api[Req, Res any] struct {
		Method          string
		Path            string
		PublicInterface bool
	}
	Ret[T any] struct {
		Result    bool   `json:"result"`
		Timestamp int64  `json:"timestamp"`
		Code      string `json:"code"`
		Message   string `json:"message"`
		Data      T      `json:"data"`
	}
)

func (a *Api[Req, Res]) Do(param Req, sign *Singer, client *fasthttp.Client) (*Ret[Res], error) {
	if client == nil {
		return nil, errors.New("http client is invalid")
	}
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	a.preParseRequest(req, param)
	if !a.PublicInterface {
		if sign == nil {
			return nil, errors.New("signer is invalid")
		}
		sign.SignReq(req)
	}
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := client.Do(req, resp); err != nil {
		return nil, err
	}
	body, err := resp.BodyUncompressed()
	if err != nil {
		return nil, err
	}
	if code := resp.StatusCode(); code >= 200 && code < 300 {
		ret := new(Ret[Res])
		if err = jsoniter.Unmarshal(body, ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
	return nil, errors.New(string(body))
}
