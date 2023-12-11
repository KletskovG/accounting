package logger

import (
	"fmt"
	"log"
	"runtime/debug"
)

const errorPrefix = "[ERROR]: "
const infoPrefix = "[INFO]: "

// Error logs stack and arguments to the console and exit program with Fatalln
func Error(messages ...any) {
	stack := debug.Stack()
	log.Fatalln(string(stack), errorPrefix, messages)
}

// Info logs prints arguments to the console with prefix
func Info(messages ...any) {
	log.Println(infoPrefix, fmt.Sprint(messages...))
}
