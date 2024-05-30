package main

import (
	"bufio"
	"fmt"
	"os"
)
var ressults []string
var conet string = ""
var content string

func double() string {
	fmt.Println("pleace!! write path to first file!( ◕ω◕)/ ^⇀ﻌ↼^")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var filefirstPath string = repeat()
	
	fmt.Println("Now! Write path to second file!(=◕ᆽ◕=)ฅ")
	var filesenodpath string = repeat()
		file, _ := os.Open(filefirstPath)
	
		defer file.Close()

		//Read the file line by line and print it to console
		scanner = bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			results = append(results, line)
			fmt.Println(line)
		}
		for i , _:= range results {
			cont = cont + results[i]
		}


		fisle, _ := os.Open(filesenodpath)
	
		defer fisle.Close()

		//Read the file line by line and print it to console
		scanner = bufio.NewScanner(fisle)
		for scanner.Scan() {
			line := scanner.Text()
			ressults = append(ressults, line)
			fmt.Println(line)
		}
		for i, _ := range ressults {
			conet = conet + ressults[i]
		}
		res := cont + conet
		return res
	}