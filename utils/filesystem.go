package utils

import (
	"log"
	"os"
)

func CreateFolder(path string) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directory %s: %v", path, err)
		}
	}
}
