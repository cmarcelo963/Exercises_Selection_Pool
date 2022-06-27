package main

import "github.com/01-edu/z01"

func main() {
	//your first variable will contain a string, this would usually be
	//'s' IF you are writing a function
	s := "Enter Text Here or Digits in a Form of String Like 12345"
	//Anything that tells you to print something, try to figure out first what is your string
	//In this case it is 's'
	//If the string isn't given like 'Print Alphabet or Numbers'
	//You basically have to declare the variable and input the string there
	//Just like the current 's' variable we have declared

	//Now we need to turn/convert our string into a slice/array of runes
	runeOfS := []rune(s)
	//This basically makes an array/slice of runes containing all characters of our string
	//This includes spaces

	//This is to find out the length of our slice of runes
	lenOfRuneOfS := len(runeOfS)
	//This can also be done directly to the string variable like below
	lenOfS := len(s)

	//To print our string/runes in z01 format
	for a := 0; a < lenOfS; a++ {
		//This Prints out EACH character/rune one by one in a loop until all characters are printed out
		z01.PrintRune(runeOfS[a])
		//This Prints out the characters/runes from beginning to end
	}
	//Don't forget the new line assuming this is the end of the code
	//In this code, let's add a new line just to make the output clearer
	z01.PrintRune('\n')

	for b := lenOfRuneOfS - 1; b != 0; b-- {
		z01.PrintRune(runeOfS[b])
		//This Prints out the characters/runes from end to beginning(backwards)
	}

	//Remember that if you need to go forward you can do '++'
	//If going backwards, do '--'

	//This code explains how to print any STRING that is given OR if you already know what you should print.

	//This is the end of the code so let's do the new line

	z01.PrintRune('\n')
}
