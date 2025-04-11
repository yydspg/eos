package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"github.com/eos/tools/errs"
)


func MD5Encrypt(s string,salt ...string) string {
	t := md5.New()
	t.Write([]byte(s))
	if len(salt) > 0 {
		t.Write([]byte(salt[0]))
	}
	ans := t.Sum(nil)
	return hex.EncodeToString(ans)
}

func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errs.WrapMsg(err, "NewCipher failed", "key", key)
	}
	blockSize := block.BlockSize()
	encryptBytes := pkcs7Padding(data, blockSize)
	crypted := make([]byte, len(encryptBytes))
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

// AesDecrypt decrypts the data with the key using AES encryption.
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errs.WrapMsg(err, "NewCipher failed", "key", key)
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	crypted := make([]byte, len(data))
	blockMode.CryptBlocks(crypted, data)
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

func pkcs7Padding(data []byte,blockSize int) []byte {
	padding := blockSize  - len(data) % blockSize
	text := bytes.Repeat([]byte{byte(padding)},padding)
	return append(data,text...)
}

func pkcs7UnPadding(data []byte) ([]byte,error) {
	length := len(data)
	if length == 0 {
		return nil,errs.New("data is nil")
	}
	ans  := int(data[length - 1])
	return data[:(length - ans)],nil
}