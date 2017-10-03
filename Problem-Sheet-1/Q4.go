//Go fundamentals Problemn sheet 1
//Author: Micheal Curley
//Adapted from https://github.com/chirag04/golang-project_euler/blob/master/p020.go
//Program to find the sum of the factorial of number
package main
//importing fmt package
import "fmt"
 
//sum of factorial function
func sumFact(num int) int {
    //initialise variables
    sum := 0;
    //digit array to hold individual values
    digits := [200] int{};
    digits[0] = 1;
    //for loop from 2 up to value user inputs
    for i := 2; i <= num; i++ {
        //for loop ranging the length of digit array
    	for j := 0; j < len(digits); j++ {
            //multiplying value by previous value to get factorial
            digits[j] *= i;
            //
    		if j > 0 && digits[j - 1] > 9 {
    			digits[j] += int(digits[j - 1] / 10);
    			digits[j - 1] %= 10;
    		}
    	}
    }
    //for loop to add up individual digits
    for i := 0; i < len(digits); i++ {
    	sum += digits[i];
    }
    return sum;
}

func main(){
    var n int

    fmt.Print("Enter an positive integer: \n")
    fmt.Scanf("%d ",&n)

    fmt.Print("Sum of factorial: \n", sumFact(n))
}
