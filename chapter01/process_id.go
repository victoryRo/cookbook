package chapter01

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// GetProcessID ... study
// Esta receta te muestra cómo usar el os paquete para obtener un PID del programa ejecutado,
// y utilizarlo con la utilidad del sistema operativo para obtener más información.
func GetProcessID() {
	// retorna el id del proceso llamado
	pid := os.Getpid()

	fmt.Printf("Process PID: %d\n", pid)

	// retorna la estructura del cmd para ejecutar
	prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	// corre el commando y retorna el stdout
	out, err := prc.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
