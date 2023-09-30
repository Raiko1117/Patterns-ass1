package main

import "fmt"

type Observer interface {
	Update(string)
}
type TicketAvailability struct {
	observers []Observer
	available bool
}

func (t *TicketAvailability) Attach(observer Observer) {
	t.observers = append(t.observers, observer)
}
func (t *TicketAvailability) SetAvailability(available bool) {
	t.available = available
	t.NotifyObservers()
}
func (t *TicketAvailability) NotifyObservers() {
	message := "Билеты доступны!"
	if !t.available {
		message = "Билеты распроданы."
	}
	for _, observer := range t.observers {
		observer.Update(message)
	}
}

type Traveler struct {
	name string
}

func (t *Traveler) Update(message string) {
	fmt.Printf("Путешественник %s: %s\n", t.name, message)
}
