package main

import (
	"github.com/stovenn/go-commons/zipwritter"
	"log"
)

func main() {
	zw := zipwritter.NewArchiveManager(nil, nil)

	err := zw.Zip("archives", "file1.txt", "file2.txt")
	if err != nil {
		log.Fatal(err)
	}

}
