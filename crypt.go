package mycrypt

// This encryption allows to store passwords in the repo that are encrypted
// This is an additional layer of security that is unnecessary because you can secure access to the repos
// however you can use the utility program to encrypt passwords.
// Passwords can be plaintext is you prefer. If the string does not start with encryption meta tag
// then it assumes that it is plaintext
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"
)

const (
	meta = ":e:"
)

var (
	// Length of string is 32 characters but this is taken care of in SetKey
	//               1111111111222222222233
	//     01234567890123456789012345678901
	key = "YouNeedToCallSetKeyToChangeThis1"
)

// SetKey name says it all
func SetKey(myKey string) {
	for {
		if len(myKey) >= 32 {
			break
		}
		myKey = myKey + myKey
	}
	key = myKey[0:32]
}

// En Encrypts the passed plain text and adds the meta at the beginning
func En(plainText string) string {
	if plainText == "" {
		return plainText
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}
	b := base64.StdEncoding.EncodeToString([]byte(plainText))
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return ""
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return meta + base64.StdEncoding.EncodeToString(ciphertext)
}

// De Decrypts the cypherText and bypases if it does not start with meta
func De(cypherText string) string {

	// This means that if there is no meta at the beginning then it is not encrypted
	// So this means can ignore
	if strings.HasPrefix(cypherText, meta) {
		cypherText = cypherText[len(meta):]
	} else {
		return cypherText
	}

	text, err := base64.StdEncoding.DecodeString(cypherText)
	if err != nil {
		return ""
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}
	if len(text) < aes.BlockSize {
		return ""
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return ""
	}
	return string(data)
}
