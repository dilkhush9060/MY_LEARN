package main

import "fmt"


func main() {
 Age := 18
//  if Age >= 18 {
//   fmt.Println("Adult")
//  } else {
//    fmt.Println("Child")
//  }

 if Age >= 18 {
  fmt.Println("Adult")
 } else if Age >= 13 {
  fmt.Println("Teen")
 } else {
  fmt.Println("Child")
 }

}
