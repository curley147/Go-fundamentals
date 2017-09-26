//Go fundamentals Problemn sheet 1
//Author: Micheal Curley
//Adapted fromhttps://github.com/golang/example/blob/master/stringutil/reverse.go
//Program to reverse a string
package main
import "fmt"
// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	//declare array of runes
	r := []rune(s)
	//for going from 0 to end off array and swapping letters two at a time
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//main function 
func main() {
	//declare variable for input string
	var str string
	//ask user for string
	fmt.Println("Please enter string: ")
	//scans user input into console
	fmt.Scanf("%s ", &str)
	fmt.Println()
	//outputs result
	fmt.Println("Reverse string: ", Reverse(str))
}