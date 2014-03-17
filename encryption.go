package main

// We should consider moving to a gcm encryption mode once cryptojs supports it.

import (
    "github.com/apexskier/crypto_padding"
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "io/ioutil"
    "encoding/base64"
)

var (
    padding crypto_padding.PKCS7
    blockSize = 16
)

// Encrypts the contents of a file and returns the encrypted byte slice.
// AES encryption with PCKS#7 padding.
func Encrypt(filename, key string) (encrypted []byte, err error) {
    var (
        aesCipher cipher.Block
        data []byte
        tempenc = make([]byte, blockSize)
    )
    aesCipher, err = aes.NewCipher([]byte(key))
    if err != nil { return }

    data, err = ioutil.ReadFile(filename)
    if err != nil {
        return encrypted, err
    }

    padded, err := padding.Pad(data, blockSize)
    if err != nil { return encrypted, err }
    for i := 0; i < len(padded) / blockSize; i++ {
        aesCipher.Encrypt(tempenc, padded[i * blockSize:i*blockSize + blockSize])
        encrypted = append(encrypted, tempenc...)
    }
    return encrypted, nil
}

// Decrypts a byte slice and returns the unencrypted byte slice.
func Decrypt(data []byte, keys string) (decrypted []byte, err error) {
    var (
        tempdec = make([]byte, blockSize)
        block cipher.Block
        iv []byte
        key []byte
    )
    if err != nil { return decrypted, err }

    b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=")

    iv = make([]byte, blockSize)
    b64.Encode(iv, []byte("wombatwombat"))
    fmt.Println("iv", string(iv))

    key = make([]byte, base64.StdEncoding.EncodedLen(len(keys)))
    b64.Encode(key, []byte(keys))
    fmt.Println("key", string(key))

    block, err = aes.NewCipher(key)
    if err != nil { return decrypted, err }

    mode := cipher.NewCBCDecrypter(block, []byte(iv))

    fmt.Println(data)

    for i := 0; i < len(data) / blockSize; i++ {
        mode.CryptBlocks(tempdec, data[i * blockSize:i*blockSize + blockSize])
        fmt.Println(tempdec)
        decrypted = append(decrypted, tempdec...)
    }
    fmt.Println(decrypted)
    unpadded, err := padding.Unpad(decrypted, blockSize)
    return unpadded, err
}

func main() {
    data, err := ioutil.ReadFile("../wombat-server/files/apexskier/test.txt")
    if err != nil { panic(err.Error()) }
    dec, err := Decrypt(data, "0b1bc60da50f9220")
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(string(dec));
    }
}
