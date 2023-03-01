package log

import (
	"fmt"
	"github.com/gob4ng/go-sdk/database"
	"github.com/gob4ng/go-sdk/utils"
	"go.uber.org/zap"
	"strconv"
)

func (z ZapLogContext) ClientInfo(tracking ZapTrackingContext) {
	logTracking := setClientLogTracking(z, tracking, SEVERITY_INFO, nil)
	z.ZapLog.Info(CLIENT_INFO_MESSAGE,
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

func (z ZapLogContext) ClientDebug(tracking ZapTrackingContext) {
	logTracking := setClientLogTracking(z, tracking, SEVERITY_DEBUG, nil)
	z.ZapLog.Debug(CLIENT_DEBUG_MESSAGE,
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

func (z ZapLogContext) ClientWarning(tracking ZapTrackingContext) {
	logTracking := setClientLogTracking(z, tracking, SEVERITY_WARNING, nil)
	z.ZapLog.Warn(CLIENT_WARNING_MESSAGE,
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

func (z ZapLogContext) ClientError(tracking ZapTrackingContext) {
	logTracking := setClientLogTracking(z, tracking, SEVERITY_ERROR, nil)
	z.ZapLog.Error(CLIENT_ERROR_MESSAGE,
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

func (z ZapLogContext) ClientResponseDebug(tracking ZapTrackingContext, responseBody string) {
	logTracking := setClientLogTracking(z, tracking, SEVERITY_DEBUG, &responseBody)
	z.ZapLog.Debug(CLIENT_RESPONSE_DEBUG_MESSAGE,
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

func (z ZapLogContext) ClientResponseWarning(tracking ZapTrackingContext, responseBody string) {
	logTracking := setClientLogTracking(z, tracking, SEVERITY_ERROR, &responseBody)
	z.ZapLog.Error(CLIENT_RESPONSE_WARNING_MESSAGE,
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

func setClientLogTracking(context ZapLogContext, tracking ZapTrackingContext, severity string, responseBody *string) database.LogTracking {

	logTracking := database.LogTracking{
		LogType:        LOG_TYPE_SERVER_DELIVERY,
		Severity:       severity,
		LogID:          tracking.LogID,
		UnixTimestamp:  tracking.UnixTimestamp,
		Environment:    context.Environment,
		Duration:       utils.GetUnixTimestamp() - tracking.UnixTimestamp,
		ClientUrl:      tracking.ClientContext.URL,
		ClientMethod:   tracking.ClientContext.HttpMethod,
		ClientHeader:   fmt.Sprintf("%s", tracking.ClientContext.Header),
		ClientRequest:  fmt.Sprintf("%s", tracking.ClientContext.OptionalContext.RequestBody),
		ClientHttpCode: tracking.ClientResponse.HttpCode,
		ClientResponse: tracking.ClientResponse.ResponseBody,
	}

	if responseBody != nil {
		logTracking.ServerResponse = *responseBody
		logTracking.LogType = LOG_TYPE_SERVER_RESPONSE
	}

	return logTracking

}
