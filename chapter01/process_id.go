package chapter01

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// GetProcessId ...
func GetProcessId() {
	pid := os.Getpid()

	fmt.Printf("Process PID: %d\n", pid)

	prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	out, err := prc.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
