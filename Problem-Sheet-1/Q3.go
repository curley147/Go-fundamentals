//Go fundamentals Problemn sheet 1
//Author: Micheal Curley
//Adapted from https://data-representation.github.io/notes/go.html 
package main

import "fmt"

func main() {
	//for loop
	for i:=1; i<101; i++{
		
		//if statement
		if i % 15 == 0 {
                fmt.Println("FizzBuzz")
		} else if i % 5 == 0 {
    		fmt.Println("Buzz")
  		} else if i % 3 == 0 {
    		fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}