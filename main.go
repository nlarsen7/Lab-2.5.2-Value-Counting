//Nicholas Larsen
//March 5, 2020
//creates a random array of 50 that determines how many odd and even numbers there are and how many of each number there is
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  var array [50]int
  var hist [20]int
  //creates blank array
  sumEven:=0
  sumOdd:=0
  zero:=0
  one:=1
  rand.Seed(time.Now().UnixNano())

  for i:=0; i<50; i++{
    computer:= (rand.Intn(20)+1)
    array[i]=computer
    //random assignment of numbers to the array
  } 
 for _, num:= range array {
   if num%2==0 {
    sumEven+=num%2+1
    //sum of the even numbers
   }
   if num%2!=0{
    sumOdd+=num%2
    //sum of the odd numbers
   }
 }
 fmt.Println("there are",sumEven, "even numbers and there are", sumOdd,"odd numbers")
 //prints out the amount of even and odd numbers

}

