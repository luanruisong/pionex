package api

import (
	"errors"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/luanruisong/pionex/api/hooks"
	"github.com/valyala/fasthttp"
)

type (
	api struct {
		Method          string
		Path            string
		PublicInterface bool
		sign            *Singer
		client          *fasthttp.Client
		before, after   *hooks.Hooks
	}
	Api[Req, Res any] struct {
		api *api
	}
	Ret[T any] struct {
		Result    bool   `json:"result"`
		Timestamp int64  `json:"timestamp"`
		Code      string `json:"code"`
		Message   string `json:"message"`
		Data      T      `json:"data"`
	}

	ApiOpts func(api *api)
)

func WithSigner(signer *Singer) ApiOpts {
	return func(api *api) {
		api.sign = signer
		api.PublicInterface = false
	}
}

func WithClient(client *fasthttp.Client) ApiOpts {
	return func(api *api) {
		api.client = client
	}
}

func WithBeforeHook(before hooks.Hook) ApiOpts {
	return func(api *api) {
		api.HookBefore(before)
	}
}

func WithAfterHook(after hooks.Hook) ApiOpts {
	return func(api *api) {
		api.HookAfter(after)
	}
}

func NewApi[req, res any](method, path string, opts ...ApiOpts) *Api[req, res] {
	a := &api{
		Method:          method,
		Path:            path,
		PublicInterface: true,
		client:          &fasthttp.Client{},
	}
	for _, v := range opts {
		v(a)
	}
	return &Api[req, res]{
		api: a,
	}
}

func (a *Api[Req, Res]) Do(param Req) (*Ret[Res], error) {
	ret := new(Ret[Res])
	if err := a.api.Do(param, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (a *api) Do(param, res any) error {
	if a.client == nil {
		return errors.New("http client is invalid")
	}
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	preParseRequest(req, param)
	if !a.PublicInterface && a.sign == nil {
		return errors.New("signer is invalid")
	}
	a.sign.SignReq(req)
	a.before.Hook(req, nil)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	if err := a.client.Do(req, resp); err != nil {
		return err
	}
	body, err := resp.BodyUncompressed()
	if err != nil {
		return err
	}
	a.after.Hook(req, resp)
	if code := resp.StatusCode(); code >= 200 && code < 300 {
		if err = jsoniter.Unmarshal(body, res); err != nil {
			return err
		}
		return nil
	}
	return errors.New(fmt.Sprintf("http response status:%d, error: %s", resp.StatusCode(), string(body)))
}

func (a *api) HookBefore(hook hooks.Hook) {
	if a.before == nil {
		a.before = new(hooks.Hooks)
	}
	a.before.Add(hook)
}

func (a *api) HookAfter(hook hooks.Hook) {
	if a.after == nil {
		a.after = new(hooks.Hooks)
	}
	a.after.Add(hook)
}
