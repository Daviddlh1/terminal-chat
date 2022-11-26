package main

import "fmt"

// Se utiliza para que una clase satisfaga una interface y al mismo tiempo permita customizacion o polimorfismo para cada caso en concreto por medio de un adaptador.

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using bankaccount %d\n", bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	bank := &BankPaymentAdapter{
		BankPayment: &BankPayment{},
		bankAccount: 5,
	}
	ProcessPayment(bank)
}
