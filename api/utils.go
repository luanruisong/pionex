package api

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"net/http"
	"reflect"
)

const (
	TsKey   = "timestamp"
	SignKey = "PIONEX-SIGNATURE"
	ApiKey  = "PIONEX-KEY"
	Host    = "https://api.pionex.com"
)

func (a *Api[Req, Res]) preParseRequest(req *fasthttp.Request, data any) {
	req.Header.SetMethod(a.Method)
	req.SetRequestURI(Host)
	u := req.URI()
	u.SetPath(a.Path)
	switch a.Method {
	case http.MethodPost, http.MethodPut:
		b, _ := jsoniter.Marshal(data)
		req.SetBody(b)
		return
	}
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return
		}
		v = v.Elem()
	case reflect.Invalid:
		return
	}
	if !v.IsZero() {
		arg := u.QueryArgs()
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			fv := fmt.Sprintf("%v", v.Field(i).Interface())
			if tag := field.Tag.Get("form"); len(tag) > 0 {
				arg.Add(tag, fv)
				continue
			}
			arg.Add(field.Name, fv)
		}
	}
}
