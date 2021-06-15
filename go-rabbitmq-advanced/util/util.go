package util

import (
	"log"
	"os"
	"strings"
)

// FailOnError function for logging error messages to the console.
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// BodyFrom function for ensuring given input on start is not empty or has at least 2 letters
func BodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}