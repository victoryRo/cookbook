package chapter01

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// GetSignal ...
func GetSignal() {
	// Crea un canal donde se recibe se envia la señal.
	// El Notify no se bloqueará cuando la señal
	// se envía y el canal no está listo.
	// Entonces es mejor crea un canal almacenado en búfer.
	sChan := make(chan os.Signal, 1)

	// Notify capturará de señales dadas y envío
	// el valor os.Signal a través del sChan.
	// Si no se especifica ninguna señal en el
	// argumento, todas las señales coinciden.
	sig := []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}
	signal.Notify(sChan, sig...)

	// Crear canal para esperar hasta que la señal es manejada.
	exitChan := make(chan int)

	go func() {
		s := <-sChan
		switch s {
		case syscall.SIGHUP:
			fmt.Println("The calling terminal has been closed")
			exitChan <- 0
		case syscall.SIGINT:
			fmt.Println("The process has been interrupted by CTRL+C")
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
