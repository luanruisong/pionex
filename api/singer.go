package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/valyala/fasthttp"
	"hash"
	"sort"
	"strings"
	"time"
)

type (
	Singer struct {
		ApiKey    string
		ApiSecret hash.Hash
	}
)

func (a *Singer) SignReq(r *fasthttp.Request) {
	u := r.URI()
	method := string(r.Header.Method())
	path := string(u.Path())
	signStr := a.Sign(method, path, u.QueryArgs(), r.Body())
	r.Header.Add(SignKey, signStr)
	r.Header.Add(ApiKey, a.ApiKey)
}

func (a *Singer) Sign(method, path string, queryArg *fasthttp.Args, body []byte) string {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	queryArg.Add(TsKey, ts)
	sb := &bytes.Buffer{}
	sb.WriteString(method)
	sb.WriteString(path)
	sb.WriteString("?")
	qs := make([]string, 0, queryArg.Len())
	queryArg.VisitAll(func(key, value []byte) {
		qs = append(qs, fmt.Sprintf("%s=%s", string(key), string(value)))
	})
	if len(qs) > 0 {
		sort.Strings(qs)
		sb.WriteString(strings.Join(qs, "&"))
	}
	if body != nil {
		sb.Write(body)
	}
	return a.hash(sb.Bytes())
}

func (a *Singer) hash(message []byte) string {
	defer a.ApiSecret.Reset()
	a.ApiSecret.Write(message)
	return hex.EncodeToString(a.ApiSecret.Sum(nil))
}

func NewSigner(apiKey, apiSecret string) *Singer {
	return &Singer{
		ApiKey:    apiKey,
		ApiSecret: hmac.New(sha256.New, []byte(apiSecret)),
	}
}
