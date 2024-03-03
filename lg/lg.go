package lg

import (
	"github.com/kr/pretty"
	"log"
)

// TODO: refactor with "terse" option
func Fail(msg string, mode string, v interface{}) {
	switch mode {
	case "live":
		log.Printf("Fail:: %s\n%v", msg, pretty.Formatter(v))
		break
	case "kill":
		log.Fatalf("Fail:: %s\n%v", msg, pretty.Formatter(v))
	case "panic":
		log.Panicf("Fail:: %s\n%v", msg, pretty.Formatter(v))
	default:
		log.Printf("Fail:: %s\n%v", msg, pretty.Formatter(v))
	}
}

func Info(msg string, terse bool, val interface{}) {
	if terse {
		log.Printf("Info:: %s\n", msg)
		return
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
