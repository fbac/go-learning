/*
	type vehicle:
		- it has doors and color

	type truck and sedan:
		- they're vehicles
		- truck has four wheels = true
		- sedan is luxury = true

	create a truck and assign values
	create a sedan and assign values
	print values by single fields with good formatting

*/
package exudemy

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	fordRanger := truck{
		fourWheels: true,
		vehicle: vehicle{
			doors: 5,
			color: "Black",
		},
	}

	mercedesCCoupe := sedan{
		luxury: true,
		vehicle: vehicle{
			doors: 3,
			color: "Neon Blue",
		},
	}

	fmt.Printf("vehicle type: %T\n\tfourwheels: %v\n\tvehicle specs:\n\tdoors: %v\n\tcolor: %s\n", fordRanger, fordRanger.fourWheels, fordRanger.vehicle.doors, fordRanger.vehicle.color)

	fmt.Printf("vehicle type: %T\n\tluxury: %v\n\tvehicle specs:\n\tdoors: %v\n\tcolor: %s\n", mercedesCCoupe, mercedesCCoupe.luxury, mercedesCCoupe.vehicle.doors, mercedesCCoupe.vehicle.color)
}
