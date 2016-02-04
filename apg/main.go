package main

import (
	"fmt"
	"crypto/rand"
	"io"
	"encoding/base64"
	"encoding/base32"
	"encoding/ascii85"
	"encoding/hex"
	"os"
)


func main() {
	bpasswd := make([]byte,16)
	_,e := io.ReadFull(rand.Reader,bpasswd)
	if(e!=nil){
		fmt.Println(e)
	}

	p0 := hex.EncodeToString(bpasswd)
	p1 := base32.StdEncoding.EncodeToString(bpasswd)
	p2 := base64.StdEncoding.EncodeToString(bpasswd)
	
	encoded := make([]byte,32)
	l := ascii85.Encode(encoded,bpasswd)
	
	fmt.Printf("%s\n%s\n%s\n%s\n", p0,p1,p2,encoded[:l] )
	fmt.Println(os.Args)
}