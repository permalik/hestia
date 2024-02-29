package lg

import (
	"github.com/kr/pretty"
	"log"
)

func Fail(msg string, mode string, err error) {
	switch mode {
	case "live":
		log.Printf("Fail:: %s\n%v", msg, pretty.Formatter(err))
		break
	case "kill":
		log.Fatalf("Fail:: %s\n%v", msg, pretty.Formatter(err))
	case "panic":
		log.Panicf("Fail:: %s\n%v", msg, pretty.Formatter(err))
	default:
		log.Printf("Fail:: %s\n%v", msg, pretty.Formatter(err))
	}
}

func Info(msg string, val interface{}) {
	if val == nil {
		log.Printf("Info:: %s\n", msg)
	}
	log.Printf("Info:: %s\n%v", msg, pretty.Formatter(val))
}

func Launch(msg string, val interface{}) {
	if val == nil {
		log.Printf("Launch:: %s\n", msg)
		return
	}
	log.Printf("Launch:: %s\n%v", msg, pretty.Formatter(val))
}

func Warn(msg string, val interface{}) {
	if val == nil {
		log.Printf("Warn:: %s\n", msg)
		return
	}
	log.Printf("Warn:: %s\n%v", msg, pretty.Formatter(val))
}
