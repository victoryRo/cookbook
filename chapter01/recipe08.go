package chapter01

import (
	"bytes"
	"fmt"
	"os/exec"
)

// CallProcessExternal ...
func CallProcessExternal() {
	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})

	prc.Stdout = out
	err := prc.Run()
	if err != nil {
		panic(err)
	}

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println(out.String())
	}
}

func CallProcess2() {
	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})

	prc.Stdout = out
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	prc.Wait()

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:")
		fmt.Println(out.String())
	}
}
