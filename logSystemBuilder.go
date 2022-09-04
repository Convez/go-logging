package gologging

import (
	"fmt"
	"io"
	"log"
	"os"
)

type LogSystemBuilder struct {
	destinations  []io.Writer
	timeFormat    string
	showTimeStamp bool
	logLevels     []string
	showFileName  bool
	showSeverity  bool
}

func NewLogSystemBuilder() *LogSystemBuilder {
	return &LogSystemBuilder{
		destinations:  []io.Writer{os.Stdout},
		timeFormat:    "2006-01-02T15:04:05.999Z",
		showTimeStamp: true,
		logLevels:     []string{ERROR, WARN, INFO, DEBUG, TRACE},
		showFileName:  true,
		showSeverity:  true,
	}
}

// WithDestination sets the logging destination to the writer in input
// Default destination is only stdout
func (l *LogSystemBuilder) WithDestination(destination io.Writer) *LogSystemBuilder {
	l.destinations = []io.Writer{destination}
	return l
}

// WithDestinations sets the logging destination to the witer list in input
// Default destination is only stdout
func (l *LogSystemBuilder) WithDestinations(destinations []io.Writer) *LogSystemBuilder {
	l.destinations = destinations
	return l
}

// WithAdditionalDestination adds the destination in input
// to the list of log destinations.
// Logs will be sent in parallel to all destinations
// Default destination is only stdout
func (l *LogSystemBuilder) WithAdditionalDestination(destination io.Writer) *LogSystemBuilder {
	l.destinations = append(l.destinations, destination)
	return l
}

// WithAdditionalDestinations adds the destinations in input
// to the list of log destinations.
// Logs will be sent in parallel to all destinations
// Default destination is only stdout
func (l *LogSystemBuilder) WithAdditionalDestinations(destinations []io.Writer) *LogSystemBuilder {
	l.destinations = append(l.destinations, destinations...)
	return l
}

// WithTimestampFormat sets the format of the timestamp.
// Default is iso8690 format: 2006-01-02T15:04:05.999Z
func (l *LogSystemBuilder) WithTimestampFormat(format string) *LogSystemBuilder {
	l.timeFormat = format
	return l
}

// WithTimestampEnabled allows to enable or disable the addition of the timestamp to the log line
// Enabled (true) by default
func (l *LogSystemBuilder) WithTimestampEnabled(enabled bool) *LogSystemBuilder {
	l.showTimeStamp = enabled
	return l
}

// WithFileNameEnabled allows to enable or disable the addition of the file name to the log line
// Enabled (true) by default
func (l *LogSystemBuilder) WithFileNameEnabled(enabled bool) *LogSystemBuilder {
	l.showFileName = enabled
	return l
}

// WithSeverityEnabled allows to enable or disable the addition of the log severity to the log line
// Enabled (true) by default
func (l *LogSystemBuilder) WithSeverityEnabled(enabled bool) *LogSystemBuilder {
	l.showSeverity = enabled
	return l
}

// WithAdditionalLevelAbove allows to add a custom log level of a severity superior to one of the existing ones
// ex. newLevel = CUSTOM aboveSeverity = INFO will add the CUSTOM severity between INFO and WARN
// Default severities are ERROR,WARN,INFO,DEBUG,TRACE
func (l *LogSystemBuilder) WithAdditionalLevelAbove(newLevel string, aboveSeverity string) *LogSystemBuilder {
	newLogs := []string{}
	for _, log := range l.logLevels {
		if log == aboveSeverity {
			newLogs = append(newLogs, newLevel)
		}
		newLogs = append(newLogs, log)
	}
	l.logLevels = newLogs
	return l
}

// WithAdditionalLevelBelow allows to add a custom log level of a severity inferior to one of the existing ones
// ex. newLevel = CUSTOM belowSeverity = INFO will add the CUSTOM severity between INFO and DEBUG
// Default severities are ERROR,WARN,INFO,DEBUG,TRACE
func (l *LogSystemBuilder) WithAdditionalLevelBelow(newLevel string, belowSeverity string) *LogSystemBuilder {
	newLogs := []string{}
	for _, log := range l.logLevels {
		newLogs = append(newLogs, log)
		if log == belowSeverity {
			newLogs = append(newLogs, newLevel)
		}
	}
	l.logLevels = newLogs
	return l
}

func (l *LogSystemBuilder) Build() *LogSystem {
	logFlags := log.Lmsgprefix
	if l.showFileName {
		logFlags = logFlags | log.Lshortfile
	}
	writer := io.MultiWriter(l.destinations...)
	if l.showTimeStamp {
		writer = timeStampWriter{writer, l.timeFormat}
	}
	level := os.Getenv("CONVEZ_LOG_LEVEL")
	levelMap := map[string]*log.Logger{}
	for _, level := range l.logLevels {
		levelMap[level] = nil
	}
	if _, ok := levelMap[level]; !ok {
		level = "INFO"
	}
	shouldWrite := true
	for _, key := range l.logLevels {
		prefix := ""
		if l.showSeverity {
			prefix = fmt.Sprintf("%s: ", key)
		}
		if shouldWrite {
			levelMap[key] = log.New(writer, prefix, logFlags)
		} else {
			levelMap[key] = log.New(io.Discard, prefix, logFlags)
		}
		if key == level {
			shouldWrite = false
		}
	}
	return &LogSystem{levelMap: levelMap}
}
