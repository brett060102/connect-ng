package util

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	// Debug logger for debugging output
	Debug = log.New(io.Discard, "", 0)
	// Info logger for standard info output
	Info = log.New(os.Stdout, "", 0)
	// QuietOut is used to simplify --quiet option
	QuietOut = log.New(io.Discard, "", 0)
)

// EnableDebug turns on debugging output
func EnableDebug() {
	Debug.SetOutput(os.Stderr)
}

func IsLoggerEnabled(l *log.Logger) bool {
	return l.Writer() != io.Discard
}

func GreenText(text string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", text)
}

func RedText(text string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", text)
}

func Bold(text string) string {
	return fmt.Sprintf("\x1b[1m%s\x1b[0m", text)
}

func LogStuff(data string) {
     	f, err := os.OpenFile("/tmp/log1.21.0.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, _ = fmt.Fprintf(f, "\n\n%s\n\n",data)
}

