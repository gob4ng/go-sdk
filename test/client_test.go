package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	sdkHttp "github.com/gob4ng/go-sdk/http"
	"github.com/gob4ng/go-sdk/log"
	"net/http"
	"testing"
)

const (
	KASPRO_INQUIRY_DETAIL = "0"
	SNAP_GW               = "1"
)

type RequestBodySnap struct {
	clientId  string `json:"clientId"`
	timestamp string `json:"timestamp"`
}

func TestHttpClient(t *testing.T) {

	apiContext := setApiContext()
	kasproInquiry := apiContext.ClientApi[KASPRO_INQUIRY_DETAIL]

	strResponse, err := kasproInquiry.HitClient()
	if err != nil {
		t.Error("ERROR", *err)
		return
	}

	t.Log("SUCCESS ", strResponse.ResponseBody)

}

func setApiContext() sdkHttp.ApiContext {
	header := make(map[string]string)
	header["Content-Type"] = gin.MIMEJSON
	header["Partner-Key"] = "5ed0eee4e661f9128fbcb022c2f5f93c6f4a4936b6c56e6e81509041"

	logID := "08161941858"
	zapLog := log.NewSetupLog("TESTING", "DEVELOPMENT", "log.txt")
	fmt.Println("ZAP LOG ", zapLog)
	kasproClientContext := sdkHttp.Context{
		ClientName: "kaspro",
		HttpMethod: http.MethodGet,
		URL:        "https://kaspro-partner.dab-partner.id/api/customer/inquiry?mobileNumber=08161941858",
		Header:     header,
		ZapLog:     &zapLog,
		OptionalContext: sdkHttp.OptionalContext{
			LogID:         logID,
			BaseAuth:      nil,
			RequestBody:   nil,
			HttpClient:    http.Client{},
			IsNeedMasking: false,
		},
	}

	snapGwClientContext := sdkHttp.Context{
		ClientName: "snap gw",
		HttpMethod: http.MethodPost,
		URL:        "https://snapgwdev.kaspro.id/tools/create-signature-token",
		Header:     header,
		ZapLog:     &zapLog,
		OptionalContext: sdkHttp.OptionalContext{
			LogID:    logID,
			BaseAuth: nil,
			RequestBody: RequestBodySnap{
				clientId:  "121312321",
				timestamp: "1231321",
			},
			HttpClient:    http.Client{},
			IsNeedMasking: false,
		},
	}

	mapApiName := make(map[string]sdkHttp.Context)
	mapApiName[KASPRO_INQUIRY_DETAIL] = kasproClientContext
	mapApiName[SNAP_GW] = snapGwClientContext
	apiContext := sdkHttp.ApiContext{ClientApi: mapApiName}
	return apiContext
}
