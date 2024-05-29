package main

import (
	"bufio"
	"fmt"
	"os"

	
)
var cont string = ""
var results []string
func reader() {
//Ask user to paste the file path
fmt.Println("Write path to this file! pleace!!(=♡ᆽ♡=)ฅ")


for true{
    scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
var filePath string = scanner.Text()
if filePath != ""{
    file, err := os.Open(filePath)
if err != nil {
	fmt.Println("Error opening the file:", err)
	return
}
defer file.Close()

//Read the file line by line and print it to console
scanner = bufio.NewScanner(file)
for scanner.Scan() {
	line := scanner.Text()
    results = append(results, line)
	fmt.Println(line)
}
for i, _ := range(results){
    cont = cont + results[i]
}

if err := scanner.Err(); err != nil {
	fmt.Println("Error reading the file:", err)
	return
}
break
}
}



}