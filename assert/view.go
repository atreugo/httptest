package assert

import (
	"testing"

	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

func View(t *testing.T, req *fasthttp.Request, fnView atreugo.View, assertFn func(resp *fasthttp.Response)) {
	s := atreugo.New(atreugo.Config{LogLevel: "fatal"})
	s.Path(string(req.Header.Method()), string(req.URI().PathOriginal()), fnView)

	Path(t, req, s, assertFn)
}
