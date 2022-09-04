package gologging_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	gologging "github.com/Convez/go-logging"
)

func TestPrintSystemBase(t *testing.T) {
	b := new(bytes.Buffer)
	builder := gologging.NewLogSystemBuilder().
		WithDestination(b)
	logSystem := builder.Build()
	infoLogger := logSystem.GetLogger(gologging.INFO)
	infoLogger.Println("Hello")
	printed := b.String()
	printParts := strings.Split(printed, " ")
	if len(printParts) != 4 {
		t.Error("Expecting 4 elements to the log: date-filename-severity-log. Got ", len(printParts), " elements")
	}
	if _, err := time.Parse("2006-01-02T15:04:05.999Z", printParts[0]); err != nil {
		t.Error("Printed time does not respect the 2006-01-02T15:04:05.999Z format: ", printParts[0])
	}
	if !strings.HasPrefix(printParts[1], "logSystem_test.go") {
		t.Error("Log does not contain file name.")
	}
	if !strings.HasPrefix(printParts[2], "INFO") {
		t.Error("Log does not contain severity.")
	}
	if !strings.HasPrefix(printParts[3], "Hello") {
		t.Error("Log was not printed correctly. Expected Hello, got: ", printParts[3])
	}
}

func TestPrintSystemNoDate(t *testing.T) {
	b := new(bytes.Buffer)
	builder := gologging.NewLogSystemBuilder().
		WithDestination(b).
		WithTimestampEnabled(false)
	logSystem := builder.Build()
	infoLogger := logSystem.GetLogger(gologging.INFO)
	infoLogger.Println("Hello")
	printed := b.String()
	printParts := strings.Split(printed, " ")
	if len(printParts) != 3 {
		t.Error("Expecting 3 elements to the log: filename-severity-log. Got ", len(printParts), " elements")
	}
	if _, err := time.Parse("2006-01-02T15:04:05.999Z", printParts[0]); err == nil {
		t.Error("Time printed, but it should be disabled: ", printParts[0])
	}
	if !strings.HasPrefix(printParts[0], "logSystem_test.go") {
		t.Error("Log does not contain file name.")
	}
	if !strings.HasPrefix(printParts[1], "INFO") {
		t.Error("Log does not contain severity.")
	}
	if !strings.HasPrefix(printParts[2], "Hello") {
		t.Error("Log was not printed correctly. Expected Hello, got: ", printParts[3])
	}
}

func TestPrintSystemNoFileName(t *testing.T) {
	b := new(bytes.Buffer)
	builder := gologging.NewLogSystemBuilder().
		WithDestination(b).
		WithFileNameEnabled(false)
	logSystem := builder.Build()
	infoLogger := logSystem.GetLogger(gologging.INFO)
	infoLogger.Println("Hello")
	printed := b.String()
	printParts := strings.Split(printed, " ")
	if len(printParts) != 3 {
		t.Error("Expecting 3 elements to the log: date-severity-log. Got ", len(printParts), " elements")
	}
	if strings.HasPrefix(printParts[1], "logSystem_test.go") {
		t.Error("Log contains file name, but it was disabled")
	}
	if _, err := time.Parse("2006-01-02T15:04:05.999Z", printParts[0]); err != nil {
		t.Error("Printed time does not respect the 2006-01-02T15:04:05.999Z format: ", printParts[0])
	}
	if !strings.HasPrefix(printParts[1], "INFO") {
		t.Error("Log does not contain severity.")
	}
	if !strings.HasPrefix(printParts[2], "Hello") {
		t.Error("Log was not printed correctly. Expected Hello, got: ", printParts[3])
	}
}

func TestPrintSystemNoSeverity(t *testing.T) {
	b := new(bytes.Buffer)
	builder := gologging.NewLogSystemBuilder().
		WithDestination(b).
		WithSeverityEnabled(false)
	logSystem := builder.Build()
	infoLogger := logSystem.GetLogger(gologging.INFO)
	infoLogger.Println("Hello")
	printed := b.String()
	printParts := strings.Split(printed, " ")
	if len(printParts) != 3 {
		t.Error("Expecting 3 elements to the log: date-filename-log. Got ", len(printParts), " elements")
	}
	if _, err := time.Parse("2006-01-02T15:04:05.999Z", printParts[0]); err != nil {
		t.Error("Printed time does not respect the 2006-01-02T15:04:05.999Z format: ", printParts[0])
	}
	if !strings.HasPrefix(printParts[1], "logSystem_test.go") {
		t.Error("Log does not contain file name")
	}
	if strings.HasPrefix(printParts[2], "INFO") {
		t.Error("Log contains severity, but it was disabled")
	}
	if !strings.HasPrefix(printParts[2], "Hello") {
		t.Error("Log was not printed correctly. Expected Hello, got: ", printParts[3])
	}
}
func TestPrintSystemNoDateNoFile(t *testing.T) {
	b := new(bytes.Buffer)
	builder := gologging.NewLogSystemBuilder().
		WithDestination(b).
		WithTimestampEnabled(false).
		WithFileNameEnabled(false)
	logSystem := builder.Build()
	infoLogger := logSystem.GetLogger(gologging.INFO)
	infoLogger.Println("Hello")
	printed := b.String()
	printParts := strings.Split(printed, " ")
	if len(printParts) != 2 {
		t.Error("Expecting 3 elements to the log: severity-log. Got ", len(printParts), " elements")
	}
	if _, err := time.Parse("2006-01-02T15:04:05.999Z", printParts[0]); err == nil {
		t.Error("Time printed but it was disabled: ", printParts[0])
	}
	if strings.HasPrefix(printParts[0], "logSystem_test.go") {
		t.Error("Filename was printed but it was disabled")
	}
	if !strings.HasPrefix(printParts[0], "INFO") {
		t.Error("Log does not contain severity")
	}
	if !strings.HasPrefix(printParts[1], "Hello") {
		t.Error("Log was not printed correctly. Expected Hello, got: ", printParts[3])
	}
}

func TestPrintSystemNoDateNoSeverity(t *testing.T) {
	b := new(bytes.Buffer)
	builder := gologging.NewLogSystemBuilder().
		WithDestination(b).
		WithTimestampEnabled(false).
		WithSeverityEnabled(false)
	logSystem := builder.Build()
	infoLogger := logSystem.GetLogger(gologging.INFO)
	infoLogger.Println("Hello")
	printed := b.String()
	printParts := strings.Split(printed, " ")
	if len(printParts) != 2 {
		t.Error("Expecting 3 elements to the log: filename-log. Got ", len(printParts), " elements")
	}
	if _, err := time.Parse("2006-01-02T15:04:05.999Z", printParts[0]); err == nil {
		t.Error("Time printed but it was disabled: ", printParts[0])
	}
	if !strings.HasPrefix(printParts[0], "logSystem_test.go") {
		t.Error("Filename was not printed")
	}
	if strings.HasPrefix(printParts[1], "INFO") {
		t.Error("Severity was printed, but it was disabled")
	}
	if !strings.HasPrefix(printParts[1], "Hello") {
		t.Error("Log was not printed correctly. Expected Hello, got: ", printParts[3])
	}
}

func TestPrintSystemNoFileNoSeverity(t *testing.T) {
	b := new(bytes.Buffer)
	builder := gologging.NewLogSystemBuilder().
		WithDestination(b).
		WithFileNameEnabled(false).
		WithSeverityEnabled(false)
	logSystem := builder.Build()
	infoLogger := logSystem.GetLogger(gologging.INFO)
	infoLogger.Println("Hello")
	printed := b.String()
	printParts := strings.Split(printed, " ")
	if len(printParts) != 2 {
		t.Error("Expecting 3 elements to the log: date-log. Got ", len(printParts), " elements")
	}
	if _, err := time.Parse("2006-01-02T15:04:05.999Z", printParts[0]); err != nil {
		t.Error("Time was not printed using the: 2006-01-02T15:04:05.999Z format", printParts[0])
	}
	if strings.HasPrefix(printParts[1], "logSystem_test.go") {
		t.Error("Filename was printed, but it was disabled")
	}
	if strings.HasPrefix(printParts[1], "INFO") {
		t.Error("Severity was printed, but it was disabled")
	}
	if !strings.HasPrefix(printParts[1], "Hello") {
		t.Error("Log was not printed correctly. Expected Hello, got: ", printParts[3])
	}
}
func TestCustomTimestampFormat(t *testing.T) {
	b := new(bytes.Buffer)
	builder := gologging.NewLogSystemBuilder().
		WithDestination(b).
		WithTimestampFormat("yyyy-MM-dd*HH:mm:ss")
	logSystem := builder.Build()
	infoLogger := logSystem.GetLogger(gologging.INFO)
	infoLogger.Println("Hello")
	printed := b.String()
	printParts := strings.Split(printed, " ")
	if _, err := time.Parse("yyyy-MM-dd*HH:mm:ss", printParts[0]); err != nil {
		t.Error("Printed time doesn't respect yyyy-MM-dd*HH:mm:ss format: ", printParts[0])
	}
}
func TestMultipleDestinations(t *testing.T) {
	b := new(bytes.Buffer)
	c := new(bytes.Buffer)
	builder := gologging.NewLogSystemBuilder().
		WithDestination(b).
		WithAdditionalDestination(c)
	logSystem := builder.Build()
	infoLogger := logSystem.GetLogger(gologging.INFO)
	infoLogger.Println("Hello")
	if b.String() != c.String() {
		t.Error("The log was propagated differently to the different buffers:")
		t.Log(b.String())
		t.Log(c.String())
	}
}

func TestAdditionalLogLevels(t *testing.T) {
	b := new(bytes.Buffer)
	builder := gologging.NewLogSystemBuilder().
		WithDestination(b).
		WithAdditionalLevelAbove("SHOULDSHOW", "INFO").
		WithAdditionalLevelBelow("SHOULDNOTSHOW", "INFO")
	logSystem := builder.Build()
	// Default min severity is INFO. This severity should be shown
	// as it's a severity above INFO
	shouldLogger := logSystem.GetLogger("SHOULDSHOW")
	shouldLogger.Print("Hello1")
	if !strings.Contains(b.String(), "Hello1") {
		t.Error("The string was not printed on the custom logger")
	}
	// Default min severity is INFO. This severity should not be shown
	// as it's a severity below INFO
	shouldntLogger := logSystem.GetLogger("SHOULDNOTSHOW")
	shouldntLogger.Print("Hello2")
	if strings.Contains(b.String(), "Hello2") {
		t.Error("The string was printed on the custom logger")
	}
}
