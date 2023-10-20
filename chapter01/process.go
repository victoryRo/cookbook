package chapter01

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// IsProcessExists ... study
// Recuperar información del proceso hijo
// La receta Llamar a un proceso externo describe cómo llamar al proceso hijo,
// de forma sincrónica y asincrónica. Naturalmente,
// para manejar el comportamiento del proceso es necesario saber más sobre el proceso.
// Esta receta muestra cómo obtener el PID y la información elemental
// sobre el proceso secundario una vez finalizado.
func IsProcessExists() {
	var cmd string

	// comprueba si el command "windows" o "sleep" existe en la terminal
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}

	fmt.Println(cmd)

	// El comando devuelve la estructura Cmd
	// para ejecutar el programa con los argumentos dados.
	prc := exec.Command(cmd, "2") // sleep 1
	// prc := exec.Command("ls", "-la")

	// Start inicia el comando especificado pero no espera a que se complete.
	// Si Start regresa exitosamente, se configurará el campo c.Process.
	err := prc.Start()
	if err != nil {
		log.Fatal(err)
	}

	// No se devuelve el estado del proceso
	// hasta que termine el proceso.
	// contiene informacion acerca del proceso
	fmt.Printf("Process state for running process %v\n", prc.ProcessState)

	// retorna el id del proceso
	fmt.Printf("PID of running process %d\n\n", prc.Process.Pid)
}
