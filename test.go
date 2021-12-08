package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
)

func Test() {
	fmt.Printf("%d\n", Monday)
}
