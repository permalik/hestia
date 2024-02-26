package lg

import "log"

func Fail(msg string, err error) {
	log.Printf("Fail:: %s\n%v", msg, err)
}

func Info(msg string, val interface{}) {
	log.Printf("Info:: %s\n%v", msg, val)
}
