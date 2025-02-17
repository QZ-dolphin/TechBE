package utility

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

type GMCrypt struct {
	PublicFile  string
	PrivateFile string
}

var (
	path        = "/app/resources/"
	privateFile = "sm2private.pem" // 私钥文件
	publicFile  = "sm2public.pem"  // 公钥文件
)

var Crypt = GMCrypt{
	PublicFile:  path + publicFile,
	PrivateFile: path + privateFile,
}

// 生成公钥、私钥
func GenerateSM2Key() {
	// 生成私钥、公钥
	priKey, err := sm2.GenerateKey(rand.Reader) // rand.Reader 用于提供加密安全的随机数生成器
	if err != nil {
		fmt.Println("秘钥产生失败：", err)
		os.Exit(1)
	}
	pubKey := &priKey.PublicKey
	// 生成文件 保存私钥、公钥
	// x509 编码
	pemPrivKey, _ := x509.WritePrivateKeyToPem(priKey, nil)
	privateFile, _ := os.Create(path + privateFile)
	defer privateFile.Close()
	privateFile.Write(pemPrivKey)
	pemPublicKey, _ := x509.WritePublicKeyToPem(pubKey)
	publicFile, _ := os.Create(path + publicFile)
	defer publicFile.Close()
	publicFile.Write(pemPublicKey)
}

// 读取密钥文件
func readPemCxt(path string) ([]byte, error) {
	// 判断文件是否存在
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		Clog("密钥文件不存在")
		GenerateSM2Key()
	}
	file, err := os.Open(path)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return []byte{}, err
	}
	buf := make([]byte, fileInfo.Size())
	_, err = file.Read(buf)
	if err != nil {
		return []byte{}, err
	}
	return buf, err
}

// 加密
func (s *GMCrypt) Encrypt(data string) (string, error) {
	pub, err := readPemCxt(s.PublicFile)
	if err != nil {
		return "", err
	}
	// read public key
	publicKeyFromPem, err := x509.ReadPublicKeyFromPem(pub)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	ciphertxt, err := publicKeyFromPem.EncryptAsn1([]byte(data), rand.Reader)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertxt), nil
}

// 解密
func (s *GMCrypt) Decrypt(data string) (string, error) {
	pri, err := readPemCxt(s.PrivateFile)
	if err != nil {
		return "", err
	}
	privateKeyFromPem, err := x509.ReadPrivateKeyFromPem(pri, nil)
	if err != nil {
		return "", err
	}
	ciphertxt, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	plaintxt, err := privateKeyFromPem.DecryptAsn1(ciphertxt)
	if err != nil {
		return "", err
	}
	return string(plaintxt), nil
}
