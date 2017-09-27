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
//declare constant
const DELTA = 0.0000001

//newton's function
func sqRoot(x float64, z float64) float64{
    //function using newtons method to calculate square root and returns a float64
    step := func() float64 {
    	return z - (z*z - x) / (2 * z)
    }
    //for loop iterates with each guess until the difference in the current guess and last is < 0.0000001
    for zz := step(); math.Abs(zz - z) > DELTA
    {
		//uses result as next guess and repeats
        z = zz
        //call newton equation on zz
        zz = step()
        //prints current guess to console
        fmt.Print("\nCurrent guess: \n", z)
    }
    return z
}
//main function calling Newtons method and Math method to compare
func main() {
    //declare variables
    var num, guess float64
    //ask user for inputs
    fmt.Print("Please enter the number you want the square root of: ")
    fmt.Scanf("%f ", &num)
    fmt.Print("Please enter you're first guess at the square root(i.e starting point in Newton's method): ")
    fmt.Scanf("%f ", &guess)
    //prints answer using Newtons method and method from math package to compare
	fmt.Println("\nNewtons method: ", sqRoot(num, guess))
	fmt.Println("\nMath.Sqrt method: ", math.Sqrt(num))
}