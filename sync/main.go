package main

import (
	"fmt"
	"sync"
)

var balance int = 100

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	// Bloque subroutinas entrantes para que esperen a que te termines de editar la variable que comparten(balance)
	lock.Lock()
	b := balance
	balance = b + amount
	// Desbloquea esas variables que comparten las GoRoutines
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	// Este es solo de lectura, quiere decir que pueden haber N lectores leyendo la varable sin ser detenidos en la linea 21(no bloque Goroutines entrantes si ya que solo van a leer).
	// En el otro caso (lock.Lock()) este bloquea las proximas GoRoutines hasta que tremine de escribir (llege al lock.Unlock()) ya que solo puede haber 1 escritor a la vez y no haya condiciones de carrera
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b
}

func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Deposit((i * 100), &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance(&lock))
}
