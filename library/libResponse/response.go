package libResponse

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	SuccessCode int = 0
	ErrorCode   int = -1
)

type Response struct {
	// code
	Code int `json:"code" example:"200"`
	// dataset
	Data interface{} `json:"data"`
	// information
	Msg string `json:"message"`
}

var response = new(Response)

func JsonExit(r *ghttp.Request, code int, msg string, data ...interface{}) {
	response.JsonExit(r, code, msg, data...)
}

func RJson(r *ghttp.Request, code int, msg string, data ...interface{}) {
	response.RJson(r, code, msg, data...)
}

func SusJson(isExit bool, r *ghttp.Request, msg string, data ...interface{}) {
	response.SusJson(isExit, r, msg, data...)
}

func FailJson(isExit bool, r *ghttp.Request, msg string, data ...interface{}) {
	response.FailJson(isExit, r, msg, data...)
}

func WriteTpl(r *ghttp.Request, tpl string, view *gview.View, params ...gview.Params) error {
	return response.WriteTpl(r, tpl, view, params...)
}

// Return JSON data and exit the current HTTP execution function。
func (res *Response) JsonExit(r *ghttp.Request, code int, msg string, data ...interface{}) {
	res.RJson(r, code, msg, data...)
	r.ExitAll()
}

// Encapsulates the standard return result data structure。
// Returns JSON with a fixed data structure:
// code: Status code (200: Success, 302: Redirect, same as the HTTP request status code);
// msg:  Request result information;
// data: Request result; the data structure of the returned result varies depending on the interface;
func (res *Response) RJson(r *ghttp.Request, code int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(&Response{
		Code: code,
		Msg:  msg,
		Data: responseData,
	})
}

// Return JSON successfully
func (res *Response) SusJson(isExit bool, r *ghttp.Request, msg string, data ...interface{}) {
	if isExit {
		res.JsonExit(r, SuccessCode, msg, data...)
	}
	res.RJson(r, SuccessCode, msg, data...)
}

// Return JSON on failure
func (res *Response) FailJson(isExit bool, r *ghttp.Request, msg string, data ...interface{}) {
	if isExit {
		res.JsonExit(r, ErrorCode, msg, data...)
	}
	res.RJson(r, ErrorCode, msg, data...)
}

// WriteTpl template output
func (res *Response) WriteTpl(r *ghttp.Request, tpl string, view *gview.View, params ...gview.Params) error {
	//Methods needed in the binding template
	view.BindFuncMap(g.Map{
		// Split the string based on length i
		"subStr": func(str interface{}, i int) (s string) {
			s1 := gconv.String(str)
			if gstr.LenRune(s1) > i {
				s = gstr.SubStrRune(s1, 0, i) + "..."
				return s
			}
			return s1
		},
	})
	r.Response.Write(view.Parse(r.GetCtx(), tpl, params...))
	return nil
}

func (res *Response) Redirect(r *ghttp.Request, location string, code ...int) {
	r.Response.RedirectTo(location, code...)
}
