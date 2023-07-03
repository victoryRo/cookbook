package samples1

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

// ProcessID ... obtenemos el id del proceso y ejecutamos un cmd desde la terminal
// para recuperar mas informacion del proceso en cuestion
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

// ExternalProcess creamos una carpeta desde una function go
// usando un proceso secundario ejecutando un cmd en la terminal
func ExternalProcess() {
	// generamos el comando para el proceso
	process := exec.Command("mkdir", "thanks")
	// buffer de bytes para almacenar la salida
	out := bytes.NewBuffer([]byte{})

	process.Stdout = out
	// lanzamos el proceso 'el CMD'
	err := process.Run()
	if err != nil {
		panic(err)
	}

	// test si el proceso fue exitoso
	if process.ProcessState.Success() {
		fmt.Println("We created a process directory")
		fmt.Println(out.String())
	}
}

// ChildProcess obtiene informacion del proceso secundario (child)
func ChildProcess() {
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}

	proc := exec.Command(cmd, "1")
	err := proc.Start()
	if err != nil {
		log.Fatal(err)
	}

	// La función Wait()
	// esperar hasta que finalice el proceso.
	err = proc.Wait()
	if err != nil {
		log.Fatal(err)
	}

	// Después de que termine el proceso
	// el *os.ProcessState contiene
	// información sencilla
	// sobre la ejecución del proceso
	fmt.Printf("PID: %d\n", proc.ProcessState.Pid())

	fmt.Printf("Process took %dms\n", proc.ProcessState.SystemTime()/time.Microsecond)

	fmt.Printf("Exited successfully : %t\n", proc.ProcessState.Success())
}
