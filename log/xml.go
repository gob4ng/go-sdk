package log

import (
	"github.com/gob4ng/go-sdk/database"
	"github.com/gob4ng/go-sdk/utils"
	"go.uber.org/zap"
	"strconv"
)

func (z ZapLogContext) XmlInfo(tracking ZapTrackingContext, message string, xmlStruct interface{}) {
	logTracking := setJsonLogTracking(z, tracking, SEVERITY_INFO, message)
	z.ZapLog.Error(message,
		zap.String("environment ", z.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("log-id", tracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("raw_data", utils.XmlToString(xmlStruct)))
}

func (z ZapLogContext) XmlDebug(tracking ZapTrackingContext, message string, xmlStruct interface{}) {
	logTracking := setJsonLogTracking(z, tracking, SEVERITY_DEBUG, message)
	z.ZapLog.Debug(message,
		zap.String("environment ", z.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("log-id", tracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("raw_data", utils.XmlToString(xmlStruct)))
}

func (z ZapLogContext) XmlWarning(tracking ZapTrackingContext, message string, xmlStruct interface{}) {
	logTracking := setJsonLogTracking(z, tracking, SEVERITY_WARNING, message)
	z.ZapLog.Warn(message,
		zap.String("environment ", z.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("log-id", tracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("raw_data", utils.XmlToString(xmlStruct)))
}

func (z ZapLogContext) XmlError(tracking ZapTrackingContext, message string, xmlStruct interface{}) {
	logTracking := setJsonLogTracking(z, tracking, SEVERITY_ERROR, message)
	z.ZapLog.Error(message,
		zap.String("environment ", z.Environment),
		zap.String("service-name", z.ServiceName),
		zap.String("log-id", tracking.LogID),
		zap.String("unix-timestamp", strconv.FormatInt(logTracking.UnixTimestamp, 10)),
		zap.String("duration", strconv.FormatInt(logTracking.Duration, 10)),
		zap.String("raw_data", utils.XmlToString(xmlStruct)))
}

func setXmlLogTracking(context ZapLogContext, tracking ZapTrackingContext, severity string, message string) database.LogTracking {

	return database.LogTracking{
		LogType:       LOG_TYPE_DEFAULT,
		Severity:      severity,
		LogID:         tracking.LogID,
		UnixTimestamp: tracking.UnixTimestamp,
		Environment:   context.Environment,
		Duration:      utils.GetUnixTimestamp() - tracking.UnixTimestamp,
		CustomMessage: message,
	}

}
