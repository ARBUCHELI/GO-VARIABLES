package main

import (
  "fmt"
)

func main() {
  var name string
  fmt.Println("What is your name?")
  fmt.Scan(&name)
  fmt.Println("Hello", name)

  var age int
  fmt.Println("What is your age?")
  fmt.Scan(&age)
  //fmt.Printf("Your name is %s, and your age is %d", name, age)
  newMessage := fmt.Sprintf("Your name is %s, and your age is %d", name, age)
  fmt.Println(newMessage)
}