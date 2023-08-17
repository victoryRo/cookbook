package chapter04

import (
	"encoding/json"
	"fmt"
	"time"
)

// SerializingTimeAndDate ...
func SerializingTimeAndDate() {
	// establece la zona horaria
	eur, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}
	// creamos una fecha con la zona horaria
	t := time.Date(2017, 11, 20, 11, 20, 10, 0, eur)

	// muestra el tipo y el valor del obj Date
	fmt.Printf("type %T and value %v\n", t, t)

	// json.Marshaler interface
	// transforma el obj Date en una slice de bytes
	// en formato de fecha RFC3339
	b, err := t.MarshalJSON()
	if err != nil {
		panic(err)
	}
	// muestra el tipo y el valor del obj Date
	// despues de pasar por (MarshalJSON) FormatoJson
	fmt.Printf("type %T and value %v\n", b, b)

	// muestra el slice de bytes como un string 'Serializado'
	fmt.Println("Serialized as RFC 3339:", string(b))
	t2 := time.Time{}
	// recibe un slice de bytes 'el obj deserializado'
	t2.UnmarshalJSON(b)
	fmt.Println("Deserialized from RFC 3339:", t2)
	fmt.Println("-----------------------------------------------------------------")

	// Serializar como época
	epoch := t.Unix()
	fmt.Println("Serialized epoch:", epoch)

	// Deserializar época
	jsonStr := fmt.Sprintf("{ \"created\":%d }", epoch)
	data := struct {
		Created int64 `json:"created"`
	}{}

	json.Unmarshal([]byte(jsonStr), &data)
	deserialized := time.Unix(data.Created, 0)
	fmt.Println("Deserialized from epoch:", deserialized)
}
