package main

import (
    "github.com/apexskier/crypto_padding"
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "io/ioutil"
)

var (
    padding crypto_padding.PKCS7
    key = []byte("wombatwombatwomb") // don't worry, this won't be in production
    blocksize = len(key)
)

func Encrypt(filename string) (encrypted []byte, err error) {
    var aesCipher cipher.Block
    aesCipher, err = aes.NewCipher(key)
    if err != nil { return }

    var data []byte
    data, err = ioutil.ReadFile(filename)
    if err != nil {
        return encrypted, err
    }

    var tempenc = make([]byte, blocksize)
    padded, err := padding.Pad(data, blocksize)
    if err != nil { return encrypted, err }
    for i := 0; i < len(padded) / blocksize; i++ {
        aesCipher.Encrypt(tempenc, padded[i * blocksize:i*blocksize + blocksize])
        encrypted = append(encrypted, tempenc...)
    }
    return encrypted, nil
}

func Decrypt(data []byte) (decrypted []byte, err error) {
    var tempdec = make([]byte, blocksize)
    var aesCipher cipher.Block
    aesCipher, err = aes.NewCipher(key)
    if err != nil { return decrypted, err }
    for i := 0; i < len(data) / blocksize; i++ {
        aesCipher.Decrypt(tempdec, data[i * blocksize:i*blocksize + blocksize])
        decrypted = append(decrypted, tempdec...)
    }
    unpadded, err := padding.Unpad(decrypted, blocksize)
    return unpadded, err
}

func main() {
    enc, err := Encrypt("test.txt")
    if err != nil {
        fmt.Println(err)
    }
    dec, err := Decrypt(enc)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(string(dec));
    }
}
