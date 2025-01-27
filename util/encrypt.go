package util

import (
	"encoding/base64"
	"math/rand"

	"github.com/chanyipiaomiao/hltool"
)

const key = "Zl0/R0cj4B3sWvRCQvP2Y01i1P2w0zi2"

var GoAes = hltool.NewGoAES([]byte(key))

// AesEncrypt 加密
func AesEncrypt(origData string) (string, error) {
	// 加密数据
	encrypt, err := GoAes.Encrypt([]byte(origData))
	if err != nil {
		Log().Error("encrypt err: %v", err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypt), err
}

// AesDecrypt 解密
func AesDecrypt(ciphertext string) (string, error) {
	// 解密数据
	decrypt, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		Log().Error("decrypt err: %v", err)
		return "", err
	}
	origData, err := GoAes.Decrypt(decrypt)
	if err != nil {
		Log().Error("decrypt err: %v", err)
		return "", err
	}
	return string(origData), err
}

// GenerateAesKey 生成AES密钥
func GenerateAesKey() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	Log().Info("key: %s, len: %d", string(key), len(key))
	return key
}
