package main

import (
	"fmt"
	"log"
)

// Permite crear clase como una especie de fabrica ya uqe todos los moldes hacen parte de la interfaz IProduct y cuando es llamda la función constructora se le especifica cual es molde enconcreto que se va a utilizar para que devuelva este mismo.

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Desktop struct {
	Computer
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

func (c *Computer) getStock() int {
	return c.stock
}

type Laptop struct {
	Computer
}

func newLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

func newDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "laptop" {
		return newLaptop(), nil
	}

	if computerType == "desktop" {
		return newDesktop(), nil
	}

	return nil, fmt.Errorf("Invalid Computer Type")
}

func printNameAndStock(p IProduct) {
	fmt.Printf("Product name: %s, with stock: %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, err := GetComputerFactory("laptop")
	if err != nil {
		log.Fatal(err)
	}
	desktop, err := GetComputerFactory("desktop")
	if err != nil {
		log.Fatal(err)
	}
	printNameAndStock(laptop)
	printNameAndStock(desktop)
}
