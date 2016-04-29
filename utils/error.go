package utils

import (
	"log"
	"os"
)

func CheckError(err error) {
	if err != nil {
		L.Critical("Fatal error: %s", err.Error())
		log.Fatal(err)
		os.Exit(1)
	}
}
