package chapter01

import (
	"encoding/json"
	"fmt"
	"os"
)

type client struct {
	consulIP   string
	connString string
}

func (c *client) String() string {
	return fmt.Sprintf("ConsulIP: %s , Connection String: %s", c.consulIP, c.connString)
}

var defaultClient = client{
	consulIP:   "localhost:9000",
	connString: "postgres://localhost:5432",
}

// configFunc funciona como un tipo a utilizar
// como opciones de funcion
type configFunc func(opt *client)

// fromFile func devuelve configFunc
// tipo. Entonces de esta manera podría leer la configuración.
// del json.
func fromFile(path string) configFunc {
	return func(opt *client) {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}(f)
		decoder := json.NewDecoder(f)

		fop := struct {
			ConsulIP string `json:"consul_ip"`
		}{}
		err = decoder.Decode(&fop)
		if err != nil {
			panic(err)
		}
		opt.consulIP = fop.ConsulIP
	}
}

// fromEnv lee la configuración
// de las variables de entorno
// y las combina con los existentes.
func fromEnv() configFunc {
	return func(opt *client) {
		connStr, exist := os.LookupEnv("CONN_DB")
		if exist {
			opt.connString = connStr
		}
	}
}

func newClient(opts ...configFunc) *client {
	client := defaultClient
	for _, fuc := range opts {
		fuc(&client)
	}
	return &client
}

func Execute() {
	client := newClient(fromFile("config.json"), fromEnv())
	fmt.Println(client.String())
}
