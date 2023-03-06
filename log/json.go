package log

import (
	"github.com/gob4ng/go-sdk/sql"
	"github.com/gob4ng/go-sdk/utils"
	"go.uber.org/zap"
	"strconv"
)

func (z ZapLogContext) JsonInfo(tracking ZapTrackingContext, message string, jsonStruct interface{}) {
	logTracking := setJsonLogTracking(z, tracking, SEVERITY_INFO, message)
	z.ZapLog.Info(message,
		zap.String("environment ", z.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("log-id", logTracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("raw-data", utils.JsonToString(jsonStruct)))
}

func (z ZapLogContext) JsonDebug(tracking ZapTrackingContext, message string, jsonStruct interface{}) {
	logTracking := setJsonLogTracking(z, tracking, SEVERITY_DEBUG, message)
	z.ZapLog.Debug(message,
		zap.String("environment ", z.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("log-id", tracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("raw-data", utils.JsonToString(jsonStruct)))
}

func (z ZapLogContext) JsonWarning(tracking ZapTrackingContext, message string, jsonStruct interface{}) {
	logTracking := setJsonLogTracking(z, tracking, SEVERITY_WARNING, message)
	z.ZapLog.Warn(message,
		zap.String("log-id", tracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("raw_data", utils.JsonToString(jsonStruct)))
}

func (z ZapLogContext) JsonError(tracking ZapTrackingContext, message string, jsonStruct interface{}) {
	logTracking := setJsonLogTracking(z, tracking, SEVERITY_ERROR, message)
	z.ZapLog.Error(message,
		zap.String("environment ", z.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("log-id", tracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("raw_data", utils.JsonToString(jsonStruct)))
}

func setJsonLogTracking(context ZapLogContext, tracking ZapTrackingContext, severity string, message string) sql.LogTracking {

	return sql.LogTracking{
		LogType:       LOG_TYPE_DEFAULT,
		Severity:      severity,
		LogID:         tracking.LogID,
		UnixTimestamp: tracking.UnixTimestamp,
		Environment:   context.Environment,
		Duration:      utils.GetUnixTimestamp() - tracking.UnixTimestamp,
		CustomMessage: message,
	}

}
