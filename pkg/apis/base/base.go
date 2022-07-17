package base

import (
	"encoding/json"

	"github.com/valyala/fasthttp"

	"github.com/huangjiasingle/suyi/cmd/apiserver/app/options"
)

type Handler struct {
	*options.ServerRunOptions
}

type Response struct {
	Code     int         `json:"code"`
	Reason   string      `json:"reason,omitempty"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Continue string      `json:"continue,omitempty"`
}

func (resp *Response) bytes() []byte {
	if resp == nil {
		resp = &Response{
			Code: fasthttp.StatusOK,
		}
	}
	bytes, _ := json.Marshal(resp)
	return bytes
}

func (h *Handler) WriteResponse(ctx *fasthttp.RequestCtx, resp *Response) {
	Cross(ctx)
	ctx.SetBody(resp.bytes())
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(resp.Code)
}

func Cross(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowCredentials, "true")
	ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowHeaders, "backstagetoken, content-type, Authorization, Content-Length, X-CSRF-Token, Token,session, x-aip-username")
	ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowMethods, "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	ctx.Response.Header.Set(fasthttp.HeaderAccessControlAllowOrigin, "*")
	ctx.Response.Header.Set(fasthttp.HeaderAccessControlExposeHeaders, "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
}
