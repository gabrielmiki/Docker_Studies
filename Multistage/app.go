package main
import (
    "fmt"
)

func main() {
  fmt.Println("What is your name:? ")
  var name string
  fmt.Scanln(&name)
  fmt.Printf("Hi, %s! I am Go language! ", name)
}
