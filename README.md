# My Crypt

This is basically an Encrypt and Decrypt Libraries that can be used for text. Useful for storing passwords in config files.

```go
    mycrypt.SetKey("Thisismysecretkey")
    plainText := "This is the text that I want to encrypt"
    cypherText := mycrypt.En(plainText)
    fmt.Println( "plainText: ", plainText )
    fmt.Println( "cypherText: ", cypherText )
```
