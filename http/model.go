package http

import (
	"github.com/gob4ng/go-sdk/log"
	"net/http"
)

type ApiContext struct {
	ClientApi map[string]Context
}

type Context struct {
	ClientName      string
	HttpMethod      string
	URL             string
	Header          map[string]string
	Debug           bool
	ZapLog          *log.ZapLogContext
	OptionalContext *OptionalContext
}

type OptionalContext struct {
	LogID            string
	UnixTimestamp    int64
	QueryParam       *map[string]string
	BaseAuth         *map[string]string
	FormData         *map[string]map[string]string
	RequestBody      interface{}
	HttpClient       http.Client
	IsNeedMasking    bool
	HiddenLogContext *HiddenLogContext
}

type HiddenLogContext struct {
	ClientName   bool
	HttpMethod   bool
	URL          bool
	Header       bool
	RequestBody  bool
	ResponseBody bool
}

type Response struct {
	HttpCode     int
	ResponseBody string
}
