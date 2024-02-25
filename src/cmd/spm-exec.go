package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

var (
	etcDir = "/var/packages/SimplePermissionManager/etc"

	permissionsPath = path.Join(etcDir, "permissions.json")
	configPath      = path.Join(etcDir, "config.json")
)

var rootPubKeys = []string{
	`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`,
	`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
MfTKoxnzYd+Lf/UpIfKpHCdlh+yzcNXMXCI=`,
}

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

type Config struct {
	TrustSignature bool `json:"trustSignature"`
}

func loadPermissionConfig() (*Permission, error) {
	fileContent, err := os.ReadFile(permissionsPath)
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

func loadConfig() (*Config, error) {
	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
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

func isUIDPermissionOK(uid int) bool {
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
	var err error
	for _, pk := range rootPubKeys {
		err = verifySignatureBySingleRootKey(pk, signature, data)
		if err == nil {
			return nil
		}
	}

	return err
}

func verifySignatureBySingleRootKey(rootPubKey string, signature *Signature, data []byte) error {
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

func isCommandSignatureOK(commandPath, sigPath string) bool {
	signature, err := loadSignature(sigPath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "load signature failed: %s\n", err)
		return false
	}

	data, err := os.ReadFile(commandPath)
	if err != nil {
		return false
	}

	err = verifySignature(signature, data)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "verify signature failed: %s\n", err)
	}
	return err == nil
}

func checkPermission(commandAndArgs []string) {
	if len(commandAndArgs) < 1 {
		_, _ = fmt.Fprintf(os.Stderr, "missing command")
		os.Exit(1)
	}

	uid := os.Getuid()
	if uid == 0 {
		return
	}

	setuidErr := syscall.Setuid(0)
	if setuidErr != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", setuidErr)
		os.Exit(1)
	}

	if isUIDPermissionOK(uid) {
		return
	}

	var (
		sigPath string
		config  *Config
		err     error
	)

	sigPath = commandAndArgs[0] + ".sig"
	stat, _ := os.Stat(sigPath)
	if stat == nil {
		_, _ = fmt.Fprintf(os.Stderr, "user with uid %d is not in permitted list and there is no signature for %s\n", uid, commandAndArgs[0])
		goto ErrPermission
	}

	config, err = loadConfig()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "load config failed: %s\n", err)
		goto ErrPermission
	}

	if !config.TrustSignature {
		_, _ = fmt.Fprintf(os.Stderr, "TrustSignature feature is disabled\n")
		goto ErrPermission
	}

	if isCommandSignatureOK(commandAndArgs[0], sigPath) {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, "%s signature is not match\n", commandAndArgs[0])

ErrPermission:
	_, _ = fmt.Fprintf(os.Stderr, "%s\n", os.ErrPermission)
	os.Exit(1)
}

func parseArgs(args []string) (spmArgs []string, commandAndArgs []string) {
	for i, arg := range args {
		if arg == "--" {
			if len(args) > i+1 {
				commandAndArgs = args[i+1:]
			}
			return
		}

		if strings.HasPrefix(arg, "-") {
			spmArgs = append(spmArgs, arg)
			continue
		}

		if len(args) > i {
			commandAndArgs = args[i:]
		}
		return
	}
	return
}

func runCommand(commandAndArgs []string, pid *string) {
	subProcess := exec.Command(commandAndArgs[0], commandAndArgs[1:]...)
	subProcess.Stdin = os.Stdin
	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr

	err := subProcess.Start()
	if err != nil {
		goto exit
	}

	if pid != nil {
		os.WriteFile(*pid, []byte(fmt.Sprintf("%d", subProcess.ProcessState.Pid())), 0644)
	}

	_ = subProcess.Wait()

exit:
	os.Exit(subProcess.ProcessState.ExitCode())
}

func main() {
	var pid *string

	spmArgs, commandAndArgs := parseArgs(os.Args[1:])

	if len(spmArgs) > 0 {
		fs := flag.NewFlagSet("spm", flag.ExitOnError)
		pid = fs.String("pid", "", "path to store command pid")
		fs.Parse(spmArgs)
	}

	checkPermission(commandAndArgs)

	runCommand(commandAndArgs, pid)
}
