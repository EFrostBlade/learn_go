package main

import "fmt"

const HelloPre = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return HelloPre + name
}

func main() {
	fmt.Println(Hello("Frost"))
}
