package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  var array [50]int
  sumEven:=0
  sumOdd:=0
  rand.Seed(time.Now().UnixNano())

  for i:=0; i<50; i++{
    computer:= (rand.Intn(20)+1)
    array[i]=computer
  } 
 for _, num:= range array {
   if num%2==0 {
    sumEven+=num%2+1
   }
   if num%2!=0{
    sumOdd+=num%2
   }
 }
 fmt.Println("there are",sumeven, "even numbers and there are", sumodd,"odd numbers")

 
}