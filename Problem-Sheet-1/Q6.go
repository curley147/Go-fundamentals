//Go fundamentals Problemn sheet 1
//Author: Micheal Curley
//Adapted from https://stackoverflow.com/questions/34259800/is-there-a-built-in-min-function-for-a-slice-of-int-arguments-or-a-variable-numb
//Program to find min and max number of list
package main

import "fmt"

//function to produce min and max values
func MinMax(array [10]int) (int, int) {
	//setting min and max to first index of array
	var max int = array[0]
	var min int = array[0]
	//iterating over array to find min and max values
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
	}
    return min, max
}

func main() {
  var x [10]int
  //ask user for list of values
  fmt.Println("Please enter  10 numbers")
  
  //populate array with values given
  for i:=0; i<10;i++{
	  //ask user for array values
	fmt.Println("Please enter number: ")
	fmt.Scanf("%d ",&x[i])
  }
  
  fmt.Println(MinMax(x))
  
}