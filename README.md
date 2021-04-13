# My Crypt

This is basically an Encrypt and Decrypt Libraries that can be used for text. Useful for storing passwords in config files.

```go
    mycrypt.SetKey("Thisismysecretkey")
    plainText := "This is the text that I want to encrypt"
    cypherText := mycrypt.En(plainText)
    fmt.Println( "plainText: ", plainText )
    fmt.Println( "cypherText: ", cypherText )
```

## Install

Sometimes `go get` puts the result in strange places. On muy system it goes into `~/go/bin/myencrypt` and `~/go/bin/mydecrypt`

```bash
go get github.com/raxisau/mycrypt/cmd...
```

## Usage

```bash
~/go/bin/myencrypt -h
```

```text
Usage of ~/go/bin/myencrypt:
  -k string
        This is the encryption key. The passkey will be expanded or trimmed to 32 characters (default "MyEncryptionKey")
  -t string
        This is your plain text (default "The Quick Brown fox Jumped over the lazy dog")
```

```bash
~/go/bin/myencrypt -k=ThisIsMyKey "-t=The Quick Brown fox 012345"
```

```text
The Quick Brown fox 012345  =>  :e:AU3poHENWWYBsR5pSdrA/SuJb5uqzvOehBoN0vZ4RtoEA/KjLadCYL8A9kqVkyungsKhCg==
```

```bash
~/go/bin/mydecrypt -h
```

```text
Usage of ~/go/bin/mydecrypt:
  -k string
        This is the encryption key. The passkey will be expanded or trimmed to 32 characters (default "MyEncryptionKey")
  -t string
        This is the ciphertext that you want to decrypt (default "This is plain text but it will decrypt as normal")
```

```bash
~/go/bin/mydecrypt -k=ThisIsMyKey "-t=:e:AU3poHENWWYBsR5pSdrA/SuJb5uqzvOehBoN0vZ4RtoEA/KjLadCYL8A9kqVkyungsKhCg=="
```

```text
:e:AU3poHENWWYBsR5pSdrA/SuJb5uqzvOehBoN0vZ4RtoEA/KjLadCYL8A9kqVkyungsKhCg==  =>  The Quick Brown fox 012345
```
