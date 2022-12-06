package main

import "fmt"

func somar(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func subtrair(s []int, c chan int) {
	sub := 0
	for _, v := range s {
		sub -= v
	}

	c <- sub
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	c := make(chan int)

	go somar(numbers, c)
	go subtrair(numbers, c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
