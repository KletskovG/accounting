package logger

import (
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
	log.Println(infoPrefix, messages)
}
