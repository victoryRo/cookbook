package chapter04

import (
	"fmt"
	"sync"
	"time"
)

// WaitingAmountOfTime ...
// Esta receta le mostrará cómo ejecutar el código con un retraso
func WaitingAmountOfTime() {
	// lanzamos un new timer
	t := time.NewTimer(3 * time.Second)

	fmt.Printf("Start waiting at %v\n", time.Now().Format(time.UnixDate))
	<-t.C // el canal ejecuta el codigo con un retraso de 3 segundos
	fmt.Printf("Code execute at %v\n", time.Now().Format(time.UnixDate))

	// usamos el WaitGroup para esperar a que se cumpla el siguiente timer
	wg := &sync.WaitGroup{}
	wg.Add(1)
	fmt.Printf("Start waiting for AfterFunc at %v\n", time.Now().Format(time.UnixDate))
	time.AfterFunc(3*time.Second, func() {
		fmt.Printf("Code executed for AfterFunc at %v\n", time.Now().Format(time.UnixDate))
		wg.Done()
	})

	wg.Wait()

	fmt.Printf("Waiting on time.After at %v\n", time.Now().Format(time.UnixDate))
	// devuelve un canal que entrega el tick despues de un periodo de tiempo
	<-time.After(3 * time.Second)
	fmt.Printf("Code resumed at %v\n", time.Now().Format(time.UnixDate))
}
