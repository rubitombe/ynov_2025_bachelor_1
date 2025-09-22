package ynov2025bachelor1

import "fmt"

func Mayor() {

	var majority uint8 = 18
	if majority > 18 {
		fmt.Println("Personne mineur")
	} else {
		fmt.Println("Personne majeur")
	}
}
