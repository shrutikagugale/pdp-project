package logger

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota + 1
)

var (
	l        *log.Logger // PRIVATE LOGGER INSTANCE
	logLevel LogLevel    // PRIVATE LOGGING LEVEL
)

type LoggerConfig struct {
	logFile  string
	logLevel int
}

func Init(config *LoggerConfig) bool {
	//func init() {
	fileHandle := os.Stderr
	l = log.New(fileHandle, "", log.Ldate|log.Ltime|log.Lshortfile)
	//log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	l.Println("Logger initialized successfully.")
	//TODO
	return false
}

func llog(format string, v ...interface{}) {
	l.Output(3, fmt.Sprintf(format, v...))
}

func Debug(format string, v ...interface{}) {
	llog(format, v)
	//TODO
	//l.Output(3, fmt.Sprintf(format, v...))
}

func SetLogLevel(level LogLevel) {
	//TODO
}

/*
func main() {
	//logConfig := &LoggerConfig{}
	//Init(logConfig)
	Debug("my name is %s", "shrutika")
}
*/
