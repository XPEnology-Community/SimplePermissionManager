package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

const rootPubKey = `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`

type PackageKey = string
type UserKey = string

type Package struct {
	UID     int    `json:"uid"`
	GID     int    `json:"gid"`
	User    string `json:"user"`
	Package string `json:"package"`
	Enabled bool   `json:"enabled"`
}

type User struct {
	UID     int    `json:"uid"`
	GID     int    `json:"gid"`
	User    string `json:"user"`
	Enabled bool   `json:"enabled"`
}

type Permission struct {
	Packages map[PackageKey]Package `json:"packages"`
	Users    map[UserKey]User       `json:"users"`
}

type Signature struct {
	Version    int         `json:"version"`
	Signature  string      `json:"signature"`
	PublicKeys []PublicKey `json:"publicKeys"`
}

type PublicKey struct {
	PublicKey string `json:"publicKey"`
	Signature string `json:"signature"`
}

func loadPermissionConfig() (*Permission, error) {
	fileContent, err := os.ReadFile("/var/packages/SimplePermissionManager/etc/permissions.json")
	if err != nil {
		return nil, err
	}

	var permissions Permission
	err = json.Unmarshal(fileContent, &permissions)
	if err != nil {
		return nil, err
	}

	return &permissions, nil
}

func loadSignature(path string) (*Signature, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var signature Signature
	err = json.Unmarshal(fileContent, &signature)
	if err != nil {
		return nil, err
	}

	return &signature, nil
}

func isPermissionOK(uid int) bool {
	permissions, err := loadPermissionConfig()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "load config failed: %s\n", err)
		return false
	}

	for _, p := range permissions.Packages {
		if p.UID == uid {
			return p.Enabled
		}
	}

	for _, u := range permissions.Users {
		if u.UID == uid {
			return u.Enabled
		}
	}
	return false
}

func isSignatureMatched(pubKey []byte, data []byte, signature []byte) error {
	pgpSignature := crypto.NewPGPSignature(signature)

	publicKeyObj, err := crypto.NewKey(pubKey)
	if err != nil {
		return err
	}

	signingKeyRing, err := crypto.NewKeyRing(publicKeyObj)
	if err != nil {
		return err
	}

	return signingKeyRing.VerifyDetached(crypto.NewPlainMessage(data), pgpSignature, crypto.GetUnixTime())
}

func isSignatureMatchedBase64(pubKey string, data []byte, signature string) error {
	pubKeyBin, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return err
	}

	signatureBin, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}

	return isSignatureMatched(pubKeyBin, data, signatureBin)
}

func verifySignature(signature *Signature, data []byte) error {
	pubKey := rootPubKey
	for _, pk := range signature.PublicKeys {
		pkData, err := base64.StdEncoding.DecodeString(pk.PublicKey)
		if err != nil {
			return err
		}
		err = isSignatureMatchedBase64(pubKey, pkData, pk.Signature)
		if err != nil {
			return err
		}
		pubKey = pk.PublicKey
	}

	return isSignatureMatchedBase64(pubKey, data, signature.Signature)
}

func isSignatureOK() bool {
	path := os.Args[1] + ".sig"
	stat, _ := os.Stat(path)
	if stat == nil {
		return false
	}

	signature, err := loadSignature(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "load signature failed: %s\n", err)
		return false
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		return false
	}

	err = verifySignature(signature, data)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "verify signature failed: %s\n", err)
	}
	return err == nil
}

func init() {
	uid := os.Getuid()
	if uid == 0 {
		return
	}

	setuidErr := syscall.Setuid(0)
	if setuidErr != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", setuidErr)
		os.Exit(1)
	}

	if !isPermissionOK(uid) && !isSignatureOK() {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", os.ErrPermission)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing command")
		os.Exit(1)
	}

	subProcess := exec.Command(os.Args[1], os.Args[2:]...)
	subProcess.Stdin = os.Stdin
	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr
	_ = subProcess.Run()
	os.Exit(subProcess.ProcessState.ExitCode())
}
