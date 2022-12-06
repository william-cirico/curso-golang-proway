package main

import (
	"fmt"
	"time"
)

func mostraNumeros() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 300)
	}
}

func mostraLetras() {
	for i := 'a'; i < 'k'; i++ {
		fmt.Printf("%c", i)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	go mostraNumeros()
	go mostraLetras()
	time.Sleep(time.Second * 5)
}
