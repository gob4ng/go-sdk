package log

import (
	"github.com/gin-gonic/gin"
	"github.com/gob4ng/go-sdk/http"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var GlobalZapLogContext *ZapLogContext

type ZapLogContext struct {
	ServiceName string
	Environment string
	RecordAble  bool
	ZapLog      zap.Logger
	TelegramBot *TelegramBot
}

type ZapTrackingContext struct {
	GinContext     *gin.Context
	ClientResponse http.Response
	ClientContext  http.Context
	LogID          string
	UnixTimestamp  int64
}

type TelegramBot struct {
	Environment string
	ServiceName string
	TelegramUrl string
	BotApiToken string
	GroupChatID string
}

const (
	SERVER_INFO_MESSAGE             = "SERVER-INFO-MESSAGE"
	SERVER_DEBUG_MESSAGE            = "SERVER-DEBUG-MESSAGE"
	SERVER_WARNING_MESSAGE          = "SERVER-WARNING-MESSAGE"
	SERVER_ERROR_MESSAGE            = "SERVER-ERROR-MESSAGE"
	SERVER_RESPONSE_DEBUG_MESSAGE   = "SERVER-RESPONSE-DEBUG-MESSAGE"
	SERVER_RESPONSE_WARNING_MESSAGE = "SERVER-RESPONSE-WARNING-MESSAGE"

	CLIENT_INFO_MESSAGE             = "CLIENT-INFO-MESSAGE"
	CLIENT_DEBUG_MESSAGE            = "CLIENT-DEBUG-MESSAGE"
	CLIENT_WARNING_MESSAGE          = "CLIENT-WARNING-MESSAGE"
	CLIENT_ERROR_MESSAGE            = "CLIENT-ERROR-MESSAGE"
	CLIENT_RESPONSE_DEBUG_MESSAGE   = "CLIENT-RESPONSE-DEBUG-MESSAGE"
	CLIENT_RESPONSE_WARNING_MESSAGE = "CLIENT-RESPONSE-WARNING-MESSAGE"

	SEVERITY_INFO    = "INFO"
	SEVERITY_WARNING = "WARNING"
	SEVERITY_DEBUG   = "DEBUG"
	SEVERITY_ERROR   = "ERROR"

	LOG_TYPE_SERVER_DELIVERY = "SERVER_DELIVERY"
	LOG_TYPE_SERVER_RESPONSE = "SERVER_RESPONSE"
	LOG_TYPE_DEFAULT         = "DEFAULT"
)

func NewSetupLog(serviceName string, mode string, logPath string) *error {

	_, err := os.OpenFile(logPath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return &err
	}

	zapConfig := zap.NewDevelopmentConfig()
	if strings.ToLower(mode) == "production" {
		zapConfig = zap.NewProductionConfig()
	}

	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapConfig.OutputPaths = []string{"stdout", logPath}

	zapLog, err := zapConfig.Build()
	if err != nil {
		return &err
	}

	GlobalZapLogContext = &ZapLogContext{ServiceName: serviceName, Environment: mode, ZapLog: *zapLog}

	return nil
}
