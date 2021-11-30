package main

import "github.com/01-edu/z01"

func main() {
	//So the idea of this code is to countdown
	//Let's assume that they want us to countdown from 999
	//This should work somewhat similar to PrintComb incase this isn't exactly how to do it

	//Let's worry about the first 9
	//Note that we're doing the numbers as a RUNE type of variable
	//This is because we need it to be a rune when we print them out
	for digit1 := '9'; digit1 >= '0'; digit1-- {
		//Now let's do the second 9
		for digit2 := '9'; digit2 >= '0'; digit2-- {
			//Now the third 9
			for digit3 := '9'; digit3 >= '0'; digit3-- {
				//Now that we have created our loop to countdown let's print them out
				//Let's print out the first digit
				z01.PrintRune(digit1)
				//We want it to be right next to each other so let's print the other two digits without doing anything else
				z01.PrintRune(digit2)
				z01.PrintRune(digit3)
				//Now that we are print the 3 numbers at the same time, let's print out a new line each time to see the separation
				z01.PrintRune('\n')

				//Additional configuration can be added here such as more if else statements
				//to add a comma after each numbers and a space and only print a new line once we reach '000'
			}
		}
	}
}

//This code have the exact same logic as PrintComb and PrintCombRev
