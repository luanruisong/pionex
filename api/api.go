package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"net/url"
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

func (a *Api[Req, Res]) createRequest(query url.Values, body io.ReadCloser) (*http.Request, error) {
	u, err := url.Parse(Host)
	if err != nil {
		return nil, err
	}
	u.RawQuery = query.Encode()
	u.Path = a.Path
	req, err := http.NewRequest(a.Method, u.String(), body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (a *Api[Req, Res]) Do(param Req) (*Ret[Res], error) {
	if client == nil {
		return nil, errors.New("http client is invalid")
	}
	query, body := ParseRequestData(a.Method, param)
	req, err := a.createRequest(query, body)
	if err != nil {
		return nil, err
	}
	if !a.PublicInterface {
		if sign == nil {
			return nil, errors.New("signer is invalid")
		}
		sign.SignReq(req)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if code := res.StatusCode; code >= 200 && code <= 300 {
		ret := new(Ret[Res])
		if err := jsoniter.Unmarshal(b, ret); err != nil {
			return nil, err
		}
		return ret, nil
	}
	return nil, errors.New(string(b))
}
