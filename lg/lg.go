package lg

import "log"

func Fail(msg string, err error) {
	log.Printf("Fail:: %s\n%v", msg, err)
}

func Info(msg string, val interface{}) {
	if val == nil {
		log.Printf("Info:: %s\n", msg)
	}
	log.Printf("Info:: %s\n%v", msg, val)
}

func Launch(msg string, val interface{}) {
	if val == nil {
		log.Printf("Launch Sequence:: %s\n", msg)
	}
	log.Printf("Launch Sequence:: %s\n%v", msg, val)
}
