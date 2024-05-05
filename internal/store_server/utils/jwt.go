package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/19fachri/store-app/internal/store_server/config"
)

func Encrypt(key, text []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("Encrypt: invalid key passed to encrypt data. Error: %v", err)
		return "", fmt.Errorf("invalid key passed to encrypt data. Error: %v", err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Printf("Encrypt: error while encrypting.Error: %v", err)
		return "", fmt.Errorf("error while encrypting. Error: %v", err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(key []byte, b64 string) (string, error) {
	text, _ := base64.StdEncoding.DecodeString(b64)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("Decrypt: unable to create new cipher, error : %v", err)
		return "", fmt.Errorf("unable to create new cipher, error : %v", err)
	}
	if len(text) < aes.BlockSize {
		log.Printf("Decrypt: invalid string to decrypt")
		return "", fmt.Errorf("invalid string to decrypt")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return string(text), nil
}

func DecryptToken(token string) ([]string, error) {
	key := []byte(config.Get().Auth.SecretKey)
	decryptedAuthToken, err := Decrypt(key, token)
	if err != nil {
		return nil, fmt.Errorf("unable to decrypt token. Error %v", err)
	}
	return strings.Split(decryptedAuthToken, config.Delimiter), nil
}

func GenerateToken(uid string) (time.Time, string, error) {
	key := []byte(config.Get().Auth.SecretKey)
	expiry := time.Now().Add(time.Hour * config.Get().Auth.TokenExpiry)
	defaultKey := []string{uid, expiry.Format(time.RFC3339)}
	authTokenString := strings.Join(defaultKey, config.Delimiter)
	token, err := Encrypt(key, []byte(authTokenString))
	return expiry, token, err
}

func GetUIDFromToken(token string) (string, error) {
	splitToken, err := DecryptToken(token)
	if err != nil {
		return "", fmt.Errorf("invalid token")
	}

	expiry, err := time.Parse(time.RFC3339, splitToken[1])
	if err != nil {
		return "", fmt.Errorf("invalid expiry token")
	}

	if time.Now().After(expiry) {
		return "", fmt.Errorf("token expired")
	}

	return splitToken[0], nil
}
