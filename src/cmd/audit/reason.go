package audit

import "fmt"

func ReasonPackageValid(uid int, pkg string) string {
	return fmt.Sprintf("package %s is valid for user %d", pkg, uid)
}

func ReasonUserVaild(uid int) string {
	return fmt.Sprintf("user %d is valid", uid)
}

func ReasonSignatureVaild(signer string) string {
	return fmt.Sprintf("signature signer %s is valid", signer)
}

func ReasonInvaild(uid int) string {
	return fmt.Sprintf("user %d is not in approved user and package list and command did not have signature", uid)
}
