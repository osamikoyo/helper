package main

import (
	"bufio"
	"os"

    "io/ioutil"
    "log"
)
var fls []string
func repeat() string {
	
	for true {
		scanner := bufio.NewScanner(os.Stdin)
	    scanner.Scan()
		res = scanner.Text()
		
		if res != ""{
			break
		}
	}
	return res
}
func slicefiles(filepad string) []string {
	files, err := ioutil.ReadDir(filepad)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fls = append(fls, file.Name())
    }
	return fls
}
