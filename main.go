package logging

import "log"

func warning(msg string) {
	log.Printf("WARN: ", msg)
}

func info(msg string) {
	log.Printf("INFO: ", msg)
}

func error(msg string) {
	log.Printf("ERROR: ", msg)
}
