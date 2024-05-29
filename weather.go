package main

import (
"fmt"
"net/http"
"io/ioutil"
)

func weather(city string) {

url := fmt.Sprintf("https://wttr.in/%s?format=3", city)

resp, err := http.Get(url)
if err != nil {
	fmt.Println("Error while fetching weather data:", err)
	return
}
defer resp.Body.Close()

body, err := ioutil.ReadAll(resp.Body)
if err != nil {
	fmt.Println("Error while reading response body:", err)
	return
}

fmt.Println(string(body))
}