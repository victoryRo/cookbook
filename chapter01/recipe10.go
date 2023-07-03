package chapter01

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// ReadOutput ...
func ReadOutput() {
	var cmd string

	if runtime.GOOS == "windows" {
		cmd = "dir"
	} else {
		cmd = "ls"
	}

	// genera el cmd de la terminal
	proc := exec.Command(cmd)

	// ejecuta el comando
	// retorna el resultado en bytes
	byt, err := proc.Output()
	if err != nil {
		panic(err)
	}

	// transforma los bytes a string()
	fmt.Println(string(byt))
}

// ReadOutputTwo ...
func ReadOutputTwo() {
	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "dir"
	} else {
		cmd = "ls"
	}

	// prepara el cmd
	proc := exec.Command(cmd)

	// prepara el buffer para leer los datos existentes
	buf := bytes.NewBuffer([]byte{})

	// escribe la salida en el buf de bytes
	// El búfer que implementa
	// La interfaz io.Writer se asigna a
	// Salida estándar del proceso
	proc.Stdout = buf

	// para evitar condiciones de carrera
	// ejecuta el comando y espera a que se complete
	err := proc.Run()
	if err != nil {
		panic(err)
	}

	// El proceso escribe la salida en
	// el buffer y usamos los bytes
	// para imprimir la salida.
	fmt.Println(string(buf.Bytes()))
}

// ReadRead ...
func ReadRead() {
	cmd := "ping"
	timeout := 2 * time.Second

	// La herramienta de línea de comando
	// "ping" se ejecuta para
	// 2 segundos
	ctx, _ := context.WithTimeout(context.TODO(), timeout)
	// prepara el comando con un contexto de tiempo
	proc := exec.CommandContext(ctx, cmd, "example.com")

	// Se obtiene la salida del proceso
	// en forma de io.ReadCloser. el subyacente
	// implementación usa os.Pipe
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	if err := proc.Start(); err != nil {
		log.Fatal(err)
	}

	// Para una lectura más cómoda
	// se utiliza bufio.Scanner.
	// La llamada de lectura está bloqueando.
	s := bufio.NewScanner(stdout)
	for s.Scan() {
		fmt.Println(s.Text())
	}
}

// Scanner ...
func Scanner() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}

// ExecuteProgram ...
func ExecuteProgram() {
	cmd := []string{"go", "run", "sample.go"}

	// Se obtiene la entrada del proceso
	// en forma de io.WriteCloser. el subyacente
	// implementación usa os.Pipe
	proc := exec.Command(cmd[0], cmd[1], cmd[2])
	stdin, _ := proc.StdinPipe()
	defer stdin.Close()

	// Para fines de depuración observamos la
	// salida del proceso ejecutado
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	go func() {
		s := bufio.NewScanner(stdout)
		for s.Scan() {
			fmt.Println("Program says:" + s.Text())
		}
	}()

	proc.Start()

	// Ahora las siguientes lineas
	// se escriben al niño
	// procesar la entrada estándar
	fmt.Println("writing input")
	io.WriteString(stdin, "hello\n")
	io.WriteString(stdin, "Golang\n")
	io.WriteString(stdin, "is awesome\n")

	time.Sleep(time.Second * 2)

	proc.Process.Kill()
}
