package main

import (
	"crypto/rand"
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	bpasswd := make([]byte, 16)
	_, e := io.ReadFull(rand.Reader, bpasswd)
	if e != nil {
		fmt.Println(e)
	}

	p1 := base32.StdEncoding.EncodeToString(bpasswd)
	p2 := base64.StdEncoding.EncodeToString(bpasswd)

	encoded := make([]byte, 32)
	l := ascii85.Encode(encoded, bpasswd)

	fmt.Printf("%x\n%s\n%s\n%s\n", bpasswd, p1, p2, encoded[:l])
	fmt.Println(os.Args)
}
