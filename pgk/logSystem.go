package gologging

import (
	"log"
)

type LogSystem struct {
	levelMap map[string]*log.Logger
}

func (l LogSystem) GetLogger(level string) *log.Logger {
	return l.levelMap[level]
}
