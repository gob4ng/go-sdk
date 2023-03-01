package test

import (
	"github.com/gob4ng/go-sdk/log"
	"testing"
)

func TestZapLog(t *testing.T) {
	serviceName := "test-service"
	mode := "development"
	logPath := "log.txt"
	zapLogSetup := log.NewSetupLog(serviceName, mode, logPath)
	zapLogSetup.Warning("TEST", "ASKDJADLKDSSD")
}
