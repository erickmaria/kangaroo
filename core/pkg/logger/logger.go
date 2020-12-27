package logger

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

type Level string

var (
	INFO  Level = "INFO"
	WARN  Level = "WARN"
	ERROR Level = "ERROR"
)

type Color string

var (
	colorReset  Color = "\033[0m"
	colorRed    Color = "\033[31m"
	colorGreen  Color = "\033[32m"
	colorYellow Color = "\033[33m"
	colorBlue   Color = "\033[34m"
	colorPurple Color = "\033[35m"
	colorCyan   Color = "\033[36m"
	colorWhite  Color = "\033[37m"
)

type Stacktace struct {
	Caller string
	Method string
	Line   int
	Stack  string
}

type Formatter struct {
	Time      string
	Module    string
	Level     Level
	Message   string
	Stacktace Stacktace
}

func (frmttr *Formatter) String() string {

	data, err := json.Marshal(frmttr)
	if err != nil {
		fmt.Println("error to marshal log to json.")
	}

	return fmt.Sprintf("time=\"%s\" [%s] level=\"%s\" Message=\"%s\" %s", frmttr.Time, frmttr.Module, frmttr.Level, frmttr.Message, data)
}

func Log(level Level, message string) {

	// get caller runtime information
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	// gorotine stack
	buf := make([]byte, 1<<16)
	stackSize := runtime.Stack(buf, true)

	st := Stacktace{
		Caller: frame.File,
		Method: frame.Function,
		Line:   frame.Line,
		Stack:  string(buf[0:stackSize]),
	}

	frmttr := Formatter{
		Time:    time.Now().Format("2006-01-02 15:04:05.000Z"),
		Module:  "Test",
		Level:   level,
		Message: message,
	}

	switch level {
	case INFO:
		fmt.Print(string(colorGreen))
	case WARN:
		fmt.Print(string(colorYellow))
	case ERROR:
		frmttr.Stacktace = st
		fmt.Print(string(colorRed))
	default:
	}

	fmt.Println(frmttr.String())
	fmt.Print(string(colorReset))

}
