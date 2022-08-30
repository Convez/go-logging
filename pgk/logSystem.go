package gologging

import (
	"io"
	"log"
)

type LogSystem struct {
	levelMap map[string]*log.Logger
}

const (
	ERROR string = "ERROR"
	WARN  string = "WARN"
	INFO  string = "INFO"
	DEBUG string = "DEBUG"
	TRACE string = "TRACE"
)

type timeStampWriter struct {
	io.Writer
	timeFormat string
}
