package main

import "fmt"

const HelloPre = "Hello, "

func Hello(name string) string {
	return HelloPre + name
}

func main() {
	fmt.Println(Hello("Frost"))
}
