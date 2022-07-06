/*
	Create a type person with
	- first name
	- last name
	- favorite ice cream flavors

	Create two person, print the values with loops
*/
package exudemy

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {

	pOni := person{firstName: "Oni", lastName: "Aranda", favIceCream: []string{"all of them!"}}
	pCoco := person{firstName: "Coco", lastName: "Aranda", favIceCream: []string{"vanilla", "rum"}}
	pKim := person{firstName: "Kim", lastName: "Aranda", favIceCream: []string{"turrón", "mojito"}}

	fmt.Println(pOni.firstName, pOni.lastName)
	for i := 0; i < len(pOni.favIceCream); i++ {
		fmt.Println("\t", pOni.favIceCream[i])
	}
	fmt.Println(pCoco.firstName, pCoco.lastName)
	for i := 0; i < len(pCoco.favIceCream); i++ {
		fmt.Println("\t", pCoco.favIceCream[i])
	}
	fmt.Println(pKim.firstName, pKim.lastName)
	for i := 0; i < len(pKim.favIceCream); i++ {
		fmt.Println("\t", pKim.favIceCream[i])
	}
}
