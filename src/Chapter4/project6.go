package main

import "fmt"

func main() {
	var conver float64 = 36.5
	fmt.Println(ftToMeter(conver))

	fmt.Println(ftoC(conver))

}

func ftToMeter(feet float64) float64 {
	return feet * .3048
}

func ftoC(fara float64) float64 {
	fara = fara - 32
	return fara * .55555
}
