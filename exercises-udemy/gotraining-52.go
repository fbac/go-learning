/*
	Create a type person with
	- first name
	- last name
	- favorite ice cream flavors

	Create two person, print the values with loops
*/
package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {

	pOni := person{firstName: "Oni", lastName: "Aranda", favIceCream: []string{"all of them"}}
	pCoco := person{firstName: "Coco", lastName: "Aranda", favIceCream: []string{"vanilla", "rum"}}
	pKim := person{firstName: "Kim", lastName: "Aranda", favIceCream: []string{"turrón", "mojito"}}

	fmt.Println(pOni)
	fmt.Println(pCoco)
	fmt.Println(pKim)

	people := make(map[string]person)
	people[pOni.firstName] = pOni
	people[pCoco.firstName] = pCoco
	people[pKim.firstName] = pKim

	for _, value := range people {
		fmt.Println(value)
	}
}
