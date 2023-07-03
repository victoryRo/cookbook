package chapter01

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// IsProcessExists ...
func IsProcessExists() {
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}

	prc := exec.Command(cmd, "1")
	err := prc.Start()
	if err != nil {
		log.Fatal(err)
	}

	// No se devuelve el estado del proceso
	// hasta que termine el proceso.
	fmt.Printf("Process state for running process %v\n", prc.ProcessState)

	fmt.Printf("PID of running process %d\n\n", prc.Process.Pid)
}
