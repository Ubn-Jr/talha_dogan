package main

import "fmt"

func div(a float32, b float32) float32 {

	var sonuc float32

	sonuc = a / b
	return sonuc
}

func main() {

	var x float32
	x = 22.0
	var y float32
	y = 5.0
	var z = div(x, y)

	if z == 3.142857 {
		fmt.Printf("evet sayınız  bir pi sayısı")

	} else {

		fmt.Printf("sayınız pi sayısı ile eşit değil!")
	}

}
