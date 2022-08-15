package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Possible numbers when two dice are rolled

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Dice rolled, Numbers:")

	for {

		r1 := rand.Intn(5) + 1
		r2 := rand.Intn(5) + 1

		fmt.Println(r1)
		fmt.Println(r2)
		break
	}
	fmt.Println("Transaction finished")

}
