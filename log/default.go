package log

import (
	"github.com/gob4ng/go-sdk/sql"
	"github.com/gob4ng/go-sdk/utils"
	"go.uber.org/zap"
	"strconv"
)

func (z ZapLogContext) Info(tracking ZapTrackingContext, message string) {
	logTracking := setDefaultLogTracking(z, tracking, SEVERITY_INFO, message)
	z.ZapLog.Info(logTracking.CustomMessage,
		zap.String("environment ", logTracking.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("logID", logTracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)))
}

func (z ZapLogContext) Debug(tracking ZapTrackingContext, message string) {
	logTracking := setDefaultLogTracking(z, tracking, SEVERITY_DEBUG, message)
	z.ZapLog.Debug(logTracking.CustomMessage,
		zap.String("environment ", logTracking.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("logID", logTracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)))
}

func (z ZapLogContext) Warning(tracking ZapTrackingContext, message string) {
	logTracking := setDefaultLogTracking(z, tracking, SEVERITY_WARNING, message)
	z.ZapLog.Warn(logTracking.CustomMessage,
		zap.String("environment ", logTracking.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("logID", logTracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)))
}

func (z ZapLogContext) Error(tracking ZapTrackingContext, message string) {
	logTracking := setDefaultLogTracking(z, tracking, SEVERITY_ERROR, message)
	z.ZapLog.Error(logTracking.CustomMessage,
		zap.String("environment ", logTracking.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("logID", logTracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)))
}

func setDefaultLogTracking(context ZapLogContext, tracking ZapTrackingContext, severity string, message string) sql.LogTracking {

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
