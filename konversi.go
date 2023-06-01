package main

import "fmt"

func main() {

	var nilai32 = 128
	var nilai42 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai42)
	fmt.Println(nilai8)
}