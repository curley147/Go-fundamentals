//Go fundamentals Problemn sheet 1
//Author: Micheal Curley
//http://www.golangpro.com/2016/01/check-string-palindrome-go.html
//Program to check if given string is a palindrome
package main

import (
 "fmt"
 "strings"
)

func main() {

// Variable Declaration
 var ip string
 //ask user for word
 fmt.Println("Enter string:")
 fmt.Scanf("%s\n", &ip)
 //convert to lowercase
 ip = strings.ToLower(ip)
 //call function to see if given word is a palinfrome
 fmt.Println(isP(ip))
}
//Function to test if the string entered is a Palindrome
func isP(s string) string {
//variable declaration for middle letter
 mid := len(s) / 2
 //variable declaration for last letter
 last := len(s) - 1
 //iterates to middle letter
 for i := 0; i < mid; i++ {
	 //checks if first and last letter are the same
  if s[i] != s[last-i] {
   return "It's not a Palindrome."
  }
 }
 return "It's a Palindrome"
}