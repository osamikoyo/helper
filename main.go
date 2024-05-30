package main

import (
	"fmt"
	"os"

)
var path, command, doubles, res string



func main() {
	fmt.Println("♡♡(ŐωŐ人)♡♡")
	fmt.Println("Yumi Helper welcome you<3")
	for true{
		fmt.Fscan(os.Stdin, &command)
	switch command{
	default:
		fmt.Println("Oh.. I cant understand you(╥Д╥)")
		fmt.Println("But!! write 'commands' for understand my functions!(´•ᴥ•`)♥")
	case "yumi":
		fmt.Println("yumi can - ")
		fmt.Println("Writing and reading groups txt files(ᵒ˵ᗜ˵ᵒ🌸)")
		fmt.Println("Organaysing files in specilal folders(◍•ᴗ•◍)")
		fmt.Println("Yumi ever can tell you the wether!(〃 ▽ 〃)")
		fmt.Println("Also, Yumi can help you with many other ideis!!(ノ╭╮.╭╮)ノ💤")
		fmt.Println("You can find me in GitHub - @osamikoyoଘ(✿˵•́ ᴗ •̀˵)")
	case "commands":
		fmt.Println("weather\n", "writetxt\n", "readtxt\n", "doubletxt\n", "org")
	case "off":
		panic("Goodbuy!")
	case "writetxt":
		base()
	case "readtxt":
		
	reader()

	case "doubletxt":
		doubles = double()
		fmt.Println(doubles + "\n" + "Thats all!")
	case "org":
		org()
	}
	}
	


}