package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

// scanme.webscantest.com otro sitio para escanar

// flag te permite pasar parametros directamente desde la terminal con el comando go run net/portConcurrente.go --site=scanme.webscantest.com
var site = flag.String("site", "scanme.nmap.org", "url to scan")

func main() {
	// Esta función hace que go revise si se le paso el flag site
	flag.Parse()
	var wg sync.WaitGroup
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			// Dial escanea el puerto del sitio que le pasaste como parametro
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *site, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open\n", port)

		}(i)
	}
	wg.Wait()
}
