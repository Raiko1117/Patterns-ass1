package main

import "fmt"

func strategyExample() {
	calculator := new(Calculator)
	calculator.SetStrategy(&AddOperation{})
	result := calculator.Calculate(7, 8)
	fmt.Println(result)

	calculator.SetStrategy(&SubstractOperation{})
	result = calculator.Calculate(7, 8)
	fmt.Println(result)
}
func observerExample() {
	ticketAvailability := &TicketAvailability{}

	traveler1 := &Traveler{name: "Raiymbek"}
	traveler2 := &Traveler{name: "Arman"}

	ticketAvailability.Attach(traveler1)
	ticketAvailability.Attach(traveler2)

	ticketAvailability.SetAvailability(true)  // Билеты доступны
	ticketAvailability.SetAvailability(false) // Билеты распроданы
}
func main() {
	strategyExample()
	//observerExample()
}
