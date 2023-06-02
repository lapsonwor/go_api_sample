package goEth

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestVerifiedEncryptWithPublicKey(t *testing.T) {
	message := `{"wallet":"Ox0704ec7713e0d9b0b3225ac994d37b4362f932e6","mark":238}`
	data := []byte(message)
	pubKey := PubToPublicKey(PublicKey)
	encryptedMessageByte := EncryptWithPublicKey(data, pubKey)
	encryptedMessage := base64.StdEncoding.EncodeToString(encryptedMessageByte)

	fmt.Println("encrypted:", encryptedMessage)

	walletPrivateKey := "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
	signedHex := generateSignedHex(walletPrivateKey, encryptedMessage)
	fmt.Println("signed Hex:", signedHex)

	publicAddr := "0x96216849c49358B10257cb55b28eA603c874b05E"
	match := VerifySign(publicAddr, signedHex, encryptedMessage)
	fmt.Println("match", match)

	rsaPrivateKey := PemToPrivateKey(PrivateKey)
	encryptedMessageByteDecoded, _ := base64.StdEncoding.DecodeString(encryptedMessage)
	messageByte := DecryptWithPrivateKey(encryptedMessageByteDecoded, rsaPrivateKey)
	plaintext := string(messageByte[:])
	fmt.Println("plain text:", plaintext)
}
