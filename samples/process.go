package samples

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// ProcessID ...
func ProcessID() {
	// id del proceso en ejecucion
	pid := os.Getpid()

	fmt.Printf("Process id Nº: %d\n", pid)

	// preparamos el comando a ejecutar
	process := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	// ejecuta el comando y retorna el stdout
	out, err := process.Output()
	if err != nil {
		panic(err)
	}

	// hacemos casting []bytes out a string
	fmt.Println(string(out))
}
