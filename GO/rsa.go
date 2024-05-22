package rsa

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

func main() {
	privKey, pubKey, err := GenerateRsaKeyPair()
	if err != nil {
		fmt.Println("Failed to generate RSA key pair: ", err)
		return
	}

	fmt.Println("Private key:")
	fmt.Println(privKey)
	fmt.Println("Public key:")
	fmt.Println(pubKey)

	data := []byte("Hello, RSA!")
	cipherData, err := RsaEncryptData(pubKey, data)
	if err != nil {
		fmt.Println("Failed to encrypt data: ", err)
		return
	}

	fmt.Println("Encrypted data:")
	fmt.Println(cipherData)

	plainData, err := RsaDecryptData(privKey, cipherData)
	if err != nil {
		fmt.Println("Failed to decrypt data: ", err)
		return
	}

	fmt.Println("Decrypted data:")
	fmt.Println(string(plainData))
}

func GenerateRsaKeyPair() (string, string, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate RSA key pair: %v", err)
	}

	privKeyBytes := x509.MarshalPKCS1PrivateKey(privKey)
	privKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privKeyBytes,
	}
	privKeyPem := pem.EncodeToMemory(privKeyBlock)

	pubKey := &privKey.PublicKey
	pubKeyBytes := x509.MarshalPKCS1PublicKey(pubKey)
	pubKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	}
	pubKeyPem := pem.EncodeToMemory(pubKeyBlock)

	return string(privKeyPem), string(pubKeyPem), nil
}

func RsaEncryptData(rsaPubKeyStr string, data []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(rsaPubKeyStr))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing the key")
	}

	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	cipherData, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, data, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt data: %v", err)
	}

	return cipherData, nil
}

func RsaDecryptData(rsaPriKeyStr string, cipherBytes []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(rsaPriKeyStr))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	plainData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, cipherBytes, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt data: %v", err)
	}

	return plainData, nil
}

