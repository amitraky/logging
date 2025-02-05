package logging

import "log"

func Warning(msg string) {
	log.Printf("WARN: %v", msg)
}

func Info(msg string) {
	log.Printf("INFO: %v", msg)
}

func Error(msg string) {
	log.Printf("ERROR: %v", msg)
}
