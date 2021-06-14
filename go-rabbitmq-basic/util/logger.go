package util

import "log"

// FailOnError function for logging error messages to the console.
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
