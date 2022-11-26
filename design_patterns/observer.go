package main

import "fmt"

// Este patron se utiliza para no estar constantemente preguntando si ha ocurrido un evento si no que en caso de que se haya dado 

// Evento al que se observa (En un caso real Topic devberia tener un nombre mas aclarativo com disponibilidad)
type Topic interface {
	register(observer Observer)
	broadcast()
}

// Observador
type Observer interface {
	getId() string
	updateValue(string)
}

// Item -> No disponible
// Item -> Disponible -> Avise -> Hay Item

type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item: %s is available\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) register(o Observer) {
	i.observers = append(i.observers, o)
}

type EmailClient struct {
	id string
}

func (ec *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %s available from client %s\n", value, ec.id)
}

func (ec *EmailClient) getId() string {
	return ec.id
}

func main() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "12ab",
	}
	secondObserver := &EmailClient{
		id: "34dc",
	}
	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.UpdateAvailability()
}
