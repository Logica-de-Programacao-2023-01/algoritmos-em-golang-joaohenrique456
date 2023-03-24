package main

import "fmt"

func main() {
	arr := [6]float64{4.5, 2.7, 9.1, 3.5, 6.9, 8.2}
	var soma float64

	for _, x := range arr {
		soma += x
	}
	media := soma / float64(len(arr))
	fmt.Println("A média é de: ", media)

}
