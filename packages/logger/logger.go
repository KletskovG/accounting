package logger

import (
	"fmt"
	"log"
	"runtime/debug"
)

const errorPrefix = "[ERROR]: "
const infoPrefix = "[INFO]: "

func Error(messages ...any) {
	stack := debug.Stack()
	log.Fatalln(string(stack), errorPrefix, messages)
}

func Info(messages ...any) {
	log.Println(infoPrefix, fmt.Sprint(messages))
}
