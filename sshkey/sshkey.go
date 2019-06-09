package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"golang.org/x/crypto/ssh"
	"strings"
)

const (
	SSH_KEY_LENGTH = 2048
)

func GenerateSshKey() (string, string, error) {

	privateKey, err := rsa.GenerateKey(rand.Reader, SSH_KEY_LENGTH)
	if err != nil {
		return "", "", errors.New("rsa key generate failure")
	}

	//private key
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	bufferPri := &bytes.Buffer{}
	err = pem.Encode(bufferPri, block)
	if err != nil {
		return "", "", errors.New("rsa key pem encoding failure")
	}

	//public key
	publicKey := &privateKey.PublicKey
	pub, err := ssh.NewPublicKey(publicKey)
	if err != nil {
		return "", "", errors.New("ssh key encoding failure")
	}
	stringPub := string(ssh.MarshalAuthorizedKey(pub))
	stringPub = strings.Trim(stringPub, "\n")
	stringPub += " Generated-by-Jvirt\n"

	return bufferPri.String(), stringPub, nil
}
