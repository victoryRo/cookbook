package chapter01

import (
	"fmt"
	"runtime"
)

const info = `
    Application %s starting.
    The binary was build by GO: %s
`

// GetInfo ...
func GetInfo() {
	fmt.Printf(info, "Example", runtime.Version())
}
