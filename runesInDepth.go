package main

import (
	"github.com/01-edu/z01"
)

func main() {
	//This code is to explain rune further and how we can navigate/manipulate positions, etc
	//First let's do an example of a string
	S := "Hello World!"
	//Now let's convert it into a slice/array of runes
	runes := []rune(S)
	//This is now how our runes variable look like behind the codes
	//In computer language, the counting ALWAYS begins with 0
	//Therefore this is how it looks like
	//  0    1    2    3    4    5    6    7    8    9    10   11
	//['H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!']

	//To access the beginning of the runes is to use 0 as the value of your declared variable in a loop
	//To access the end, "len("stringVariableNameHere") - 1" you're doing -1 because when we start counting we count from 0, so we don't need the actual number

	//The idea of using for loops to iterate forward or backwards a rune is to access their positions one by one
	for posOfRune := 0; posOfRune < len(runes); posOfRune++ {
		//This loop basically means that for as long as our position of rune is lower than the total length of our runes which is 11
		//We're going to do everything inside the loop over and over again, until the condition is met/no longer valid
		//For now, let's simply print the characters
		z01.PrintRune(runes[posOfRune])
	}
	z01.PrintRune('\n')

	//Now let's say that we want to access a specific part of the string/runes
	//Let's say we ONLY want to print the second half of the string
	//One thing we can do is to do a maths formula to get that number

	//This gives us an integer of 6 saved in posOfRune variable
	//If we now use that as our index inside the rune, we will immediately start on 6

	for posOfRune := len(S) / 2; posOfRune < len(runes); posOfRune++ {
		z01.PrintRune(runes[posOfRune])
	}
	z01.PrintRune('\n')

	//The same logic works for if we only want the first word
	for posOfRune := 0; posOfRune < (len(S) / 2); posOfRune++ {
		z01.PrintRune(runes[posOfRune])
	}
	z01.PrintRune('\n')

	//Another way of turning a string into runes is by using the 'for, range' loop
	//Rather than storing the runes somewhere first, we can just iterate through the string using the for, range loop
	//This is only possible to iterate forward/upwards/ascending
	//Here's an example
	//we use an underscore('_') if we don't intend to use the index/position of that character/rune
	//In this case, we will not be using it
	for _, char := range S {
		//Because each character is already a rune, we can just print out the character straight away
		z01.PrintRune(char)
		//This prints out all characters then breaks the loop
	}
	z01.PrintRune('\n')

	//What if we want to change or manipulate a specific character? Whether given or not?
	//Here's one way to do it
	for _, char := range S {
		//Let's say that we want to ignore the characters 'l'
		if char == 'l' || char == 'L' {
			//The moment our loop finds the character 'l' we don't do anything, rather we just continue the loop
			continue
		} else {
			//If it isn't 'l', we print the character
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')

	//What if we want to replace the character instead of ignoring it?
	//IT WORKS EXACTLY THE SAME AS IGNORING IT
	for _, char := range S {
		if char == 'l' {
			//This is IF the character is lowercase
			z01.PrintRune('r')
		} else if char == 'L' {
			//This is IF the character is uppercase
			z01.PrintRune('R')
		} else {
			//This is just to print our the characters normally
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')

	//One of the functions that is used relatively often with runes is 'Append'
	//This basicaly means to 'ADD' to the slice/array that we already have
	//Let's say that we already have an array of numbers
	num := []rune("01234")
	for a := 0; a < len(num); a++ {
		z01.PrintRune(num[a])
	}
	z01.PrintRune('\n')
	//Above, we have an array/slice of five numbers of runes, but what if we want to store another rune of number.
	//We simply just use the append function
	//the append function is used like this
	num = append(num, '5')
	for a := 0; a < len(num); a++ {
		z01.PrintRune(num[a])
	}
	//append("the variable you want to add something into", "what you want to add into it")
	z01.PrintRune('\n')
}


