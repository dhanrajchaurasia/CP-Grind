package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"time"
)

// The secret key used for encryption and decryption (must be 16, 24, or 32 bytes long)
var key = []byte("abcdefghijklmnopqrstuvwx")

func Encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))

	return fmt.Sprintf("%x", cipherText), nil
}

func Decrypt(cipherTextHex string) (string, error) {
	cipherText, err := hexToBytes(cipherTextHex)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("invalid cipher text length")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

func hexToBytes(hex string) ([]byte, error) {
	bytes := make([]byte, len(hex)/2)
	_, err := fmt.Sscanf(hex, "%x", &bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func CreateToken(username string) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Subject:   username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(key)
	return token.SignedString(secretKey)
}
