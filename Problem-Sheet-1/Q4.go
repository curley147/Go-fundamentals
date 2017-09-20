//Go fundamentals Problemn sheet 1
//Author: Micheal Curley
//Adapted from http://www.golangprograms.com/basic-programs/go-program-to-find-factorial-of-a-number.html 
//Program to find Factorial of number
package main
import "fmt"
 
/* Variable Declaration */
var factVal int = 1 
                       
var i int = 1
var n int
 
/*     function declaration        */
func factorial(n int) int {   
    if(n < 0){
        fmt.Print("Factorial of negative number doesn't exist.")    
    }else{        
        for i:=1; i<=n; i++ {
            factVal *= int(i)  
        }
         
    }    
    return factVal  /* return from function*/
}
 
func main(){    
    fmt.Print("Enter a positive integer between 0 - 50 : ")
    fmt.Scan(&n)   
    fmt.Print("Factorial is: ",factorial(n))
}