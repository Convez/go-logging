package gologging

import (
	"fmt"
	"io"
	"time"
)

type timeStampWriter struct {
	writer     io.Writer
	timeFormat string
}

func (t timeStampWriter) Write(bytes []byte) (int, error) {
	return t.writer.Write([]byte(fmt.Sprintf("%s %s", time.Now().UTC().Format(t.timeFormat), string(bytes))))
}
