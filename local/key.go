package local

import (
	"crontools/utils"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

func generatePrivateKey(bitSize int) (privateKey *rsa.PrivateKey, err error) {
	if privateKey, err = rsa.GenerateKey(rand.Reader, bitSize); err != nil {
		return
	}
	if err = privateKey.Validate(); err != nil {
		return
	}
	return
}

func generatePublicKey(privateKey *rsa.PublicKey) (pubKeyBytes []byte, err error) {
	var publicRsaKey ssh.PublicKey
	if publicRsaKey, err = ssh.NewPublicKey(privateKey); err != nil {
		return
	} else {
		pubKeyBytes = ssh.MarshalAuthorizedKey(publicRsaKey)
	}
	return
}

func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) (privatePEM []byte) {
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	privatePEM = pem.EncodeToMemory(&privBlock)
	return
}

func HaveSSHkey() bool {
	if utils.Exists(PrivateKeyFile) && utils.Exists(PublicKeyFile) {
		return true
	} else {
		return false
	}
}

func GenSSHkey() (err error) {
	if HaveSSHkey() == true {
		return
	}

	if utils.Exists(SSHdir) == false {
		if err = os.MkdirAll(SSHdir, 0700); err != nil {
			return
		}
	}

	var privateKey *rsa.PrivateKey
	if privateKey, err = generatePrivateKey(2048); err != nil {
		return
	}
	privateKeyBytes := encodePrivateKeyToPEM(privateKey)

	var publicKeyBytes []byte
	if publicKeyBytes, err = generatePublicKey(&privateKey.PublicKey); err != nil {
		return
	}

	if err = ioutil.WriteFile(PrivateKeyFile, privateKeyBytes, 0600); err != nil {
		return
	}
	if err = ioutil.WriteFile(PublicKeyFile, publicKeyBytes, 0600); err != nil {
		return
	}

	return
}

func GetPublicKey() (publicKey []byte, err error) {
	if publicKey, err = ioutil.ReadFile(PublicKeyFile); err != nil {
		return
	}
	return
}

func GetPrivateKey() (privateKey []byte, err error) {
	if privateKey, err = ioutil.ReadFile(PrivateKeyFile); err != nil {
		return
	}
	return
}

func GetSSHAuth() (auth []ssh.AuthMethod, err error) {
	var privateKey []byte
	if privateKey, err = GetPrivateKey(); err != nil {
		return
	}
	var signer ssh.Signer
	if signer, err = ssh.ParsePrivateKey(privateKey); err != nil {
		return
	}
	auth = append(auth, ssh.PublicKeys(signer))
	return
}
