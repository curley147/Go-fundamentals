//Adapted from: https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go
//Program to merge two lists of integers and sort them
package main
//importing sort package aswell as fmt
import (
	"fmt"
	"sort"
)

func main() {

	var a [5]int
  	//ask user for list of values
  	fmt.Println("Please enter 5 numbers for 1st array")
  
  	//populate array with values given
  	for i:=0; i<5;i++{
	  //ask user for array values
	fmt.Println("Please enter number: ")
	//scan values into slice
	fmt.Scanf("%d ",&a[i])
  	}
	var b [5]int
  	//ask user for list of values
  	fmt.Println("Please enter 5 numbers for 2nd array")
  
  	//populate array with values given
  	for i:=0; i<5;i++{
	  //ask user for array values
	fmt.Println("Please enter number: ")
	fmt.Scanf("%d ",&b[i])
  	}

	//declaring capacity of slices and copying from arrays
	x, y := a[:5], b[:5]
	fmt.Printf("x: %v, y: %v\n", x, y)

	x = append(x, y...)
	fmt.Printf("x: %v\n", x)
	//sorting slices
	sort.Ints(x)
	//prints sorted list
	fmt.Printf("Sorted list: %v", x)
}