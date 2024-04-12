package utils

import (
	"log"
	stos_messages "logger/stos-messages"
)

func PanicOnError(err error, message string) {
	if err != nil {
		log.Panicf("%s: %s", message, err)
	}
}

func EventToLoggerLevel(event stos_messages.LogType) string {
	switch event {
	case stos_messages.LogType_ERROR:
	        return "ERROR"
	case stos_messages.LogType_INFO:
	        return "INFO"
	case stos_messages.LogType_WARNING:
	        return "WARNING"
	}
	return "INFO"
}
