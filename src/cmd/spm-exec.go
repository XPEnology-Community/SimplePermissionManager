package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type Permission struct {
	UID  int    `json:"uid"`
	GID  int    `json:"gid"`
	User string `json:"user"`
}

func loadPermissionConfig() ([]Permission, error) {
	fileContent, err := os.ReadFile("/var/packages/SimplePermissionManager/etc/permissions.json")
	if err != nil {
		return nil, err
	}

	var permissions []Permission
	err = json.Unmarshal(fileContent, &permissions)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func isPermissionOK(uid int) bool {
	permissions, err := loadPermissionConfig()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "load config failed: %s\n", err)
		return false
	}
	for _, p := range permissions {
		if p.UID == uid {
			return true
		}
	}
	return false
}

func init() {
	uid := os.Getuid()
	if uid == 0 {
		return
	}

	if !isPermissionOK(uid) {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", os.ErrPermission)
		os.Exit(1)
	}

	setuidErr := syscall.Setuid(0)
	if setuidErr != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", setuidErr)
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
