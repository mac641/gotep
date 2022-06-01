package logger_test

import (
	"reflect"
	"testing"

	"github.com/mac641/gotep/src/lib/logger"
)

func TestGetLogger(t *testing.T) {
	logger1 := logger.GetLogger()
	logger2 := logger.GetLogger()

	if !reflect.DeepEqual(logger1, logger2) {
		t.Errorf("Expected logger to exist only once but got different instances returned")
	}
}

func TestReport(t *testing.T) {
	codeSuccess := logger.GetLogger().Report(1, 1)
	if codeSuccess != 0 {
		t.Errorf("Expected report to return success code (0) but got %d", codeSuccess)
	}

	codeFailure := logger.GetLogger().Report(0, 2)
	if codeFailure == 0 {
		t.Errorf("Expected report to return error code but got success code")
	}
}
