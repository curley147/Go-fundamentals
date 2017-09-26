//Go fundamentals Problemn sheet 1
//Author: Micheal Curley
//Adapted from https://gist.github.com/abesto/3476594
//A program to calculate the square root of a number using Newton’s method
package main
//import packages
import (
	"fmt"
	"math"
)
//declare constants
const DELTA = 0.0000001
const INITIAL_Z = 100.0
//newton's function
func Sqrt(x float64) (z float64) {
    z = INITIAL_Z
    //function using newtons method to calculate square root and returns a float64
    step := func() float64 {
    	return z - (z*z - x) / (2 * z)
    }
    //for loop until absolute value > 0.0000001
    for zz := step(); math.Abs(zz - z) > DELTA
    {
		//repeats calculation until difference is insignificant
    	z = zz
	zz = step()
    }
    return
}
//main function calling Newtons method and Math method to compare
func main() {
	fmt.Println("Newtons method: ", Sqrt(500))
	fmt.Println("Math.Sqrt method: ", math.Sqrt(500))
}