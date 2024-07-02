package api

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

const (
	TsKey   = "timestamp"
	SignKey = "PIONEX-SIGNATURE"
	ApiKey  = "PIONEX-KEY"
	Host    = "https://api.pionex.com"
)

var (
	client *http.Client
)

func WithHttpClient(c *http.Client) {
	client = c
}

func ParseRequestData(method string, data any) (url.Values, io.ReadCloser) {
	switch method {
	case http.MethodPost, http.MethodDelete:
		b, _ := jsoniter.Marshal(data)
		return url.Values{}, io.NopCloser(bytes.NewReader(b))
	}
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return url.Values{}, nil
		}
		v = v.Elem()
	case reflect.Invalid:
		return url.Values{}, nil
	}
	if v.IsZero() {
		return url.Values{}, nil
	}
	var u = url.Values{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fv := fmt.Sprintf("%v", v.Field(i).Interface())
		if tag := field.Tag.Get("form"); len(tag) > 0 {
			u.Add(tag, fv)
			continue
		}
		u.Add(field.Name, fv)
	}
	return u, nil
}
