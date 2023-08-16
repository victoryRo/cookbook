package chapter04

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// RunCodePeriodically ...
func RunCodePeriodically() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	ticker := time.NewTicker(time.Second)
	stop := make(chan bool)

	go func() {
		defer func() { stop <- true }()
		for {
			select {
			case <-ticker.C:
				fmt.Println("Tick")
			case <-stop:
				fmt.Println("Goroutine closing")
				return
			}
		}
	}()

	// Bloquear hasta
	// recibir la señal
	<-c
	ticker.Stop()

	// Detener la gorutina
	stop <- true
	// Espera hasta que el
	<-stop
	fmt.Println("App stopped")
}
