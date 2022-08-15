package main

import "fmt"

func fishCount() int {
	//you can post the number of fish here
	//           ||
	//           \/
	fishCount := 10
	return fishCount
}

func getNewFishCount() {
	if fishCount() < 10 {
		//if you have less than ten fish this code will work
		fmt.Println("You have a les than 10 fish, buy some more")
	} else if fishCount() == 10 {
		//if you have a exactly ten fish this code will work
		fmt.Println("You have a 10 fish")
	} else if fishCount() > 10 {
		//if you have more than ten fish this code will work
		getNewFish()

	}
}

func getNewFish() {
	fmt.Println("You have more than ten fish, Enlarge the aquarium ")
	//return getNewFish

}

//number of fish in the aquarium
func main() {
	fishCount()
	getNewFishCount()
	deneme()
}
func deneme() {

	fishes := [3]string{"Lepistes ", "Moli ", "Japon "}

	fmt.Println(fishes)

}
