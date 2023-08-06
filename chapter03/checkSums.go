package chapter03

import (
	"crypto"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

var content = "This is content to check"

// CheckSum ...
func CheckSum() {
	checksum := MD5(content)
	checksum2 := FileMD5("chapter03/content.txt")

	fmt.Printf("CheckSum 1: %s\n", checksum)
	fmt.Printf("CheckSum 2: %s\n", checksum2)

	if checksum == checksum2 {
		fmt.Println("Content matches!!!")
	}

}

// MD5 crea el md5
// hash para contenido dado codificado en
// cadena hexadecimal
func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

// FileMD5 crea hash md5 con codificación hexadecimal
// del contenido del archivo
func FileMD5(path string) string {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

// ShaPanic ...
func ShaPanic() {
	crypto.SHA1.New()
}
