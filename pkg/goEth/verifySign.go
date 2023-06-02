package goEth

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"lapson_go_api_sample/pkg/logger"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
)

var verifySignLogger *logrus.Entry = logger.GetLogger("verifySign")

func VerifySign(publicAddr string, signatureHash string, message string) bool {
	message = "\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message
	data := []byte(message)
	hash := crypto.Keccak256Hash(data)
	signatureHashByte, err := hexutil.Decode(signatureHash)
	if err != nil {
		verifySignLogger.Error("decode signature error: ", err)
		return false
	}
	signatureHashByte[64] -= 27 // Transform yellow paper V from 27/28 to 0/1
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signatureHashByte)
	if err != nil {
		verifySignLogger.Error("ecrecover error: ", err)
		return false
	}
	sigPubAddress := PublicKeyBytesToAddress(sigPublicKey)
	userPubAddress := common.HexToAddress(publicAddr)
	matches := sigPubAddress.Hex() == userPubAddress.Hex()
	return matches
}

func PublicKeyBytesToAddress(publicKey []byte) common.Address {
	var buf []byte

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKey[1:]) // remove EC prefix 04
	buf = hash.Sum(nil)
	address := buf[12:]

	return common.HexToAddress(hex.EncodeToString(address))
}

func generateSignedHex(pKey string, message string) string {
	privateKey, err := crypto.HexToECDSA(pKey)
	if err != nil {
		verifySignLogger.Error(err)
	}

	data := []byte(message)
	hash := crypto.Keccak256Hash(data)

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if signature[64] == 0 || signature[64] == 1 {
		signature[64] += 27
	}
	if err != nil {
		verifySignLogger.Error(err)
	}

	hexString := hexutil.Encode(signature)
	return hexString
}

func VerifySignWithPrivate(pKey string, message string) bool {
	privateKey, err := crypto.HexToECDSA(pKey)
	if err != nil {
		verifySignLogger.Error(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		verifySignLogger.Error("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte(message)
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x06b3dfaec148fb1bb2b066f10ec285e7c9bf402ab32aa78a5d38e34566810cd2

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x3e7791153a49c9f07157c1515212e5efd90968aa1d60ce4a9cb722465a46566f46067765ef480c6658e09b7db66b5e2b3074398b5de7ab8a6dccae0209d520bc01

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches) // true

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true
	return verified
}
