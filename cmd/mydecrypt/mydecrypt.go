package main

import (
	"flag"
	"fmt"

	"github.com/raxisau/mycrypt"
)

var (
	decryptionKey = flag.String("k", "MyEncryptionKey", "This is the encryption key. The passkey will be expanded or trimmed to 32 characters")
	ciphertext    = flag.String("t", "This is plain text but it will decrypt as normal", "This is the ciphertext that you want to decrypt")
)

func main() {
	flag.Parse()
	mycrypt.SetKey(*decryptionKey)
	plainText := mycrypt.De(*ciphertext)
	fmt.Println(*ciphertext, " => ", plainText)
}
