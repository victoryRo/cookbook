package chapter01

import (
	"bytes"
	"fmt"
	"os/exec"
)

// CallProcessExternal ...
func CallProcessExternal() {
	// ejecuta comando en la terminal desde go
	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})

	prc.Stdout = out
	// inicia el cmd y espera hasta que se ejecute
	err := prc.Run()
	if err != nil {
		panic(err)
	}

	// prc.ProcessState continer info acerca de la salida del process
	// prc.ProcessState.Success() reporta si la salida del process fue exitosa
	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println(out.String())
	}
}

// CallProcess2 ...
func CallProcess2() {
	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})

	prc.Stdout = out
	// inicia el cmd especificado pero no espera hasta que se complete
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	// espera hasta que la salida del cmd se complete
	prc.Wait()

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println(out.String())
	}
}
