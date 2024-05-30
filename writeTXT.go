package main 

import(
	"os"
	"fmt"

)
var fileath string = "C:/Users/Turbo Boost/Documents/Texts"
var namefile string
var text string


func base(){


	
	fmt.Println("Write name of file, which you want!<3")
	fmt.Fscan(os.Stdin, &namefile)
	fmt.Println(namefile)
	fmt.Println("Now, Tell me text to write it in file!(๑✪ᆺ✪๑)")
	fmt.Fscan(os.Stdin, &text)
	
	
	fmt.Println(fileath + namefile)
	write(text, namefile, fileath)
}
func write(text string, namefile string, pathfile string){
	file, err := os.Create(pathfile +"/" + namefile + ".txt")
	if err != nil {
		fmt.Println("Emh... Error, can you fix him pleace?!(ㅅ´ ˘ `)♡", err)
	return
	}
	defer file.Close()


// Write a string to the file
	_, err = file.WriteString(text)
	if err != nil {
			fmt.Println("Emh... Error, can you fix him pleace?!(ㅅ´ ˘ `)♡", err)
		return
	}

	fmt.Println("Yumi do it!!!! *°~ LoVE You ~°*")
}
