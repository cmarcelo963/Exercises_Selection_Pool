package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	//1st arg is "Hello World!"
	//2nd arg is "l"
	//3rd arg is "r"
	runeOfArg := []rune(args[2])
	runeOfArg2 := []rune(args[3])
	for _, char := range args[1] {
		if char == runeOfArg[0] {
			z01.PrintRune(runeOfArg2[0])
		} else {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
