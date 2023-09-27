package chapter01

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var writer *os.File

// CorrectClosure ... study
func CorrectClosure() {

	// El archivo se abre como un archivo de registro para escribir
	// De esta manera representamos los recursos
	var err error
	writer, err = os.OpenFile(fmt.Sprintf("test_%d.log", time.Now().Unix()), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// El código se ejecuta en una rutina de forma independiente. Entonces, en caso de que el programa sea
	// terminado desde afuera, necesitamos avisar a la gorutina a través de closeChan
	closeChan := make(chan bool)

	go func() {
		// bucle for infinito
		for {
			// por cada iteracion el bucle espera un segundo
			time.Sleep(time.Second)

			select {
			// el canal por default es false cuando enviamos true se termina la go rutina y salimos del bucle
			case <-closeChan:
				log.Println("Goroutine closing")
				return
			default:
				// mientras el canal sea falso retornara este mensaje por default en consola
				log.Println("Writing to log")
				// y escribira en el archivo el registro con la hora actual
				_, err := io.WriteString(writer, fmt.Sprintf("Logging access %s\n", time.Now().String()))
				if err != nil {
					panic(err)
				}
			}
		}
	}()

	// canal para almacenar la os.Signal
	sigChan := make(chan os.Signal, 1)
	// slice tipo os.Signal con algunas señales de respuestas especificadas
	sig := []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	// capta la señal del sistema y la envia al canal sigChan
	signal.Notify(sigChan, sig...)

	// Esto está bloqueando la lectura del codigo siguiente hasta que reciba el sigChan
	// desde la función Notify, hasta tanto la go rutina continuara su bucle
	<-sigChan

	// Después de recibir la señal todo el código detrás de la lectura del canal podría ser
	// considerado como una limpieza. SECCIÓN DE LIMPIEZA
	close(closeChan)
	releaseAllResources()
	fmt.Println("The application shut down gracefully")
}

func releaseAllResources() {
	_, err := io.WriteString(writer, "Application releasing all resources\n")
	if err != nil {
		panic(err)
	}
	writer.Close()
}
