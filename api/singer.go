package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"net/http"
	"net/url"
	"sort"
	"time"
)

type (
	Singer struct {
		ApiKey    string
		ApiSecret hash.Hash
	}
)

func (a *Singer) SignReq(r *http.Request) {
	u := r.URL
	query := u.Query()
	var b []byte
	if r.Body != nil {
		b, _ = io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewReader(b))
	}
	signStr := a.Sign(r.Method, u.Path, query, b)
	r.URL.RawQuery = query.Encode()
	r.Header.Add(SignKey, signStr)
	r.Header.Add(ApiKey, a.ApiKey)
}

func (a *Singer) Sign(method, path string, query url.Values, body []byte) string {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	query.Add(TsKey, ts)
	sb := &bytes.Buffer{}
	sb.WriteString(method)
	sb.WriteString(path)
	sb.WriteString("?")
	qs := make([]string, 0, len(query))
	for i := range query {
		qs = append(qs, i)
	}
	if len(qs) > 0 {
		sort.Strings(qs)
		for i, v := range qs {
			if i > 0 {
				sb.WriteString("&")
			}
			sb.WriteString(fmt.Sprintf("%s=%s", v, query.Get(v)))
		}
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
