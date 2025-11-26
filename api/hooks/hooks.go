package hooks

import (
	"github.com/valyala/fasthttp"
)

type (
	Hook  func(req *fasthttp.Request, res *fasthttp.Response)
	Hooks struct {
		h []Hook
	}
)

func (hooks *Hooks) Add(hook Hook) {
	hooks.h = append(hooks.h, hook)
}

func (hooks *Hooks) Hook(req *fasthttp.Request, res *fasthttp.Response) {
	if hooks == nil {
		return
	}
	for _, v := range hooks.h {
		v(req, res)
	}
}

func CurlHook(f func(string)) Hook {
	return func(req *fasthttp.Request, res *fasthttp.Response) {
		f(curlHook(req))
	}
}

func RespBodyHook(f func(string)) Hook {
	return func(req *fasthttp.Request, res *fasthttp.Response) {
		f(string(res.Body()))
	}
}
