package lg

import (
	"github.com/kr/pretty"
	"log"
)

func Fail(msg string, err error) {
	log.Fatalf("Fail:: %s\n%v", msg, pretty.Formatter(err))
}

func Info(msg string, val interface{}) {
	if val == nil {
		log.Printf("Info:: %s\n", msg)
	}
	log.Printf("Info:: %s\n%v", msg, pretty.Formatter(val))
}

func Launch(msg string, val interface{}) {
	if val == nil {
		log.Printf("Launch Sequence:: %s\n", msg)
		return
	}
	log.Printf("Launch Sequence:: %s\n%v", msg, pretty.Formatter(val))
}

func Warn(msg string, val interface{}) {
	if val == nil {
		log.Printf("Warn:: %s\n", msg)
		return
	}
	log.Printf("Warn:: %s\n%v", msg, pretty.Formatter(val))
}
