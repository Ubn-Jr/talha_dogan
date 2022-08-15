package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Possible numbers when two dice are rolled
//bir zar atıldığında gelmesi ihtimal rakamlar

func main() {
	//random sayı vermeye yardımcı olan kısısm
	rand.Seed(time.Now().UnixNano())
	
        // zarın atıldığını söyleyen bir çıktı
	fmt.Println("Dice rolled, Numbers:")

	for {
               //0 dan 5 e kadar rasgele bir sayı verip üzerine 1 ekleyen kısım 
	       //(1 eklemesenin sebebi zarın üzeinde 0 bulunmaması)
		r1 := rand.Intn(5) + 1
		r2 := rand.Intn(5) + 1
                
		// verdiğimiz random sayının çıktısını veren kısım
		fmt.Println(r1)
		fmt.Println(r2)
		//break sonuz döngüyü durdurur
		//ve ekrana bir kere yazar
		break
	}
	//işlemin bittiğini haber veren çıktı 
	fmt.Println("Transaction finished")

}
