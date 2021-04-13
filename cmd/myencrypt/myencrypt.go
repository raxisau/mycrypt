package main

import (
	"flag"
	"fmt"

	"github.com/raxisau/mycrypt"
)

var (
	encryptionKey = flag.String("k", "MyEncryptionKey", "This is the encryption key. The passkey will be expanded or trimmed to 32 characters")
	plainText     = flag.String("t", "The Quick Brown fox Jumped over the lazy dog", "This is your plain text")
)

func main() {
	flag.Parse()
	mycrypt.SetKey(*encryptionKey)
	ciphertext := mycrypt.En(*plainText)
	fmt.Println(*plainText, " => ", ciphertext)
}
