package samples1

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// CatchSignals ...
func CatchSignals() {
	// creamos un canal de tipo signal
	sigChan := make(chan os.Signal, 1)

	// Notify captara las señales y la enviara por el canal sigChan
	sig := []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}
	signal.Notify(sigChan, sig...)

	// maneja el codigo de salida del switch
	exitChan := make(chan int)

	go func() {
		s := <-sigChan
		switch s {
		case syscall.SIGHUP:
			fmt.Println("the colling terminal has ben closed")
			exitChan <- 0
		case syscall.SIGINT:
			fmt.Println("the process has been interrupted by CTRL+C")
			exitChan <- 1
		case syscall.SIGTERM:
			fmt.Println("kill SIGTERM was executed for process")
			exitChan <- 1
		case syscall.SIGQUIT:
			fmt.Println("kill SIGQUIT was executed for process")
			exitChan <- 1
		}
	}()

	code := <-exitChan
	os.Exit(code)
}
