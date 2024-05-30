package main

import (
	"bufio"
	"os"
)

func repeat() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	for true {
		res = scanner.Text()
		
		if res != ""{
			break
		}
	}
	return res
}