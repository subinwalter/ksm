package aes

import (
	"github.com/subinwalter/toast/cipher"
	"github.com/subinwalter/toast/crypto"
)

func EncryptWithECB(key []byte, plainText []byte) ([]byte, error) {
	mode := cipher.NewECBMode()
	cipher, err := crypto.NewAESWith(key, mode)
	if err != nil {
		return nil, err
	}

	ciphertext := cipher.Encrypt(plainText)
	return ciphertext, nil
}
