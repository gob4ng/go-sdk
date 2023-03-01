package log

import (
	"fmt"
	"github.com/gob4ng/go-sdk/database"
	"github.com/gob4ng/go-sdk/utils"
	"go.uber.org/zap"
	"strconv"
)

func (z ZapLogContext) ServerInfo(tracking ZapTrackingContext) {
	logTracking := setServerLogTracking(z, tracking, SEVERITY_INFO, nil, nil)
	z.ZapLog.Info(SERVER_INFO_MESSAGE,
		zap.String("environment ", logTracking.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("logID", logTracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("server-url", logTracking.ServerUrl),
		zap.String("server-method", logTracking.ServerMethod),
		zap.String("server-header", logTracking.ServerHeader),
		zap.String("server-request", logTracking.ServerRequest))
}

func (z ZapLogContext) ServerDebug(tracking ZapTrackingContext) {
	logTracking := setServerLogTracking(z, tracking, SEVERITY_DEBUG, nil, nil)
	z.ZapLog.Debug(SERVER_DEBUG_MESSAGE,
		zap.String("environment ", logTracking.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("logID", logTracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("server-url", logTracking.ServerUrl),
		zap.String("server-method", logTracking.ServerMethod),
		zap.String("server-header", logTracking.ServerHeader),
		zap.String("server-request", logTracking.ServerRequest))
}

func (z ZapLogContext) ServerWarning(tracking ZapTrackingContext) {
	logTracking := setServerLogTracking(z, tracking, SEVERITY_WARNING, nil, nil)
	z.ZapLog.Warn(SERVER_WARNING_MESSAGE,
		zap.String("environment ", logTracking.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("logID", logTracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("server-url", logTracking.ServerUrl),
		zap.String("server-method", logTracking.ServerMethod),
		zap.String("server-header", logTracking.ServerHeader),
		zap.String("server-request", logTracking.ServerRequest))
}

func (z ZapLogContext) ServerError(tracking ZapTrackingContext) {
	logTracking := setServerLogTracking(z, tracking, SEVERITY_ERROR, nil, nil)
	z.ZapLog.Error(SERVER_ERROR_MESSAGE,
		zap.String("environment ", logTracking.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("logID", logTracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("server-url", logTracking.ServerUrl),
		zap.String("server-method", logTracking.ServerMethod),
		zap.String("server-header", logTracking.ServerHeader),
		zap.String("server-request", logTracking.ServerRequest))
}

func (z ZapLogContext) ServerResponseDebug(tracking ZapTrackingContext, httpCode int, responseBody string) {
	logTracking := setServerLogTracking(z, tracking, SEVERITY_DEBUG, &httpCode, &responseBody)
	z.ZapLog.Debug(SERVER_RESPONSE_DEBUG_MESSAGE,
		zap.String("environment ", logTracking.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("logID", logTracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("server-url", logTracking.ServerUrl),
		zap.String("server-method", logTracking.ServerMethod),
		zap.String("server-header", logTracking.ServerHeader),
		zap.String("server-request", logTracking.ServerRequest),
		zap.String("server-response", logTracking.ServerResponse))
}

func (z ZapLogContext) ServerResponseWarning(tracking ZapTrackingContext, httpCode int, responseBody string) {
	logTracking := setServerLogTracking(z, tracking, SEVERITY_ERROR, &httpCode, &responseBody)
	z.ZapLog.Error(SERVER_RESPONSE_WARNING_MESSAGE,
		zap.String("environment ", logTracking.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("logID", logTracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("server-url", logTracking.ServerUrl),
		zap.String("server-method", logTracking.ServerMethod),
		zap.String("server-header", logTracking.ServerHeader),
		zap.String("server-request", logTracking.ServerRequest),
		zap.String("server-response", logTracking.ServerResponse))
}

func setServerLogTracking(context ZapLogContext, tracking ZapTrackingContext, severity string, httpCode *int, responseBody *string) database.LogTracking {

	logTracking := database.LogTracking{
		LogType:       LOG_TYPE_SERVER_DELIVERY,
		Severity:      severity,
		LogID:         tracking.LogID,
		UnixTimestamp: tracking.UnixTimestamp,
		Environment:   context.Environment,
		Duration:      utils.GetUnixTimestamp() - tracking.UnixTimestamp,
		ServerUrl:     tracking.GinContext.Request.Host + tracking.GinContext.Request.URL.Path,
		ServerMethod:  tracking.GinContext.Request.Method,
		ServerHeader:  fmt.Sprintf("%s", tracking.GinContext.Request.Header),
		ServerRequest: utils.GetRawData(tracking.GinContext.Request.Body),
	}

	if responseBody != nil {
		logTracking.ServerHttpCode = *httpCode
		logTracking.ServerResponse = *responseBody
		logTracking.LogType = LOG_TYPE_SERVER_RESPONSE
	}

	return logTracking

}
