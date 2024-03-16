package audit

import (
	"gorm.io/gorm"
)

type Action int
type Type int

var (
	ActionApproved = Action(0)
	ActionDenied   = Action(1)

	TypeUser      = Type(0)
	TypePackage   = Type(1)
	TypeSignature = Type(2)
)

type Audit struct {
	gorm.Model

	Command   string
	UID       int
	Package   string
	Signature string
	Type      int
	Action    int
	Reason    string
}

func Add(cmd string, uid int, action Action, pkg string, signature string, typ Type, reason string) error {
	result := db.Create(&Audit{
		Command:   cmd,
		UID:       uid,
		Package:   pkg,
		Signature: signature,
		Type:      int(typ),
		Action:    int(action),
		Reason:    reason,
	})
	return result.Error
}

func AddPackageApproved(cmd string, uid int, pkg string) error {
	return Add(cmd, uid, ActionApproved, pkg, "", TypePackage, ReasonPackageValid(uid, pkg))
}

func AddUserApproved(cmd string, uid int) error {
	return Add(cmd, uid, ActionApproved, "", "", TypeUser, ReasonUserVaild(uid))
}

func AddSignatureApproved(cmd string, uid int, signature string) error {
	return Add(cmd, uid, ActionApproved, "", signature, TypeSignature, ReasonSignatureVaild(signature))
}

func AddDenied(cmd string, uid int) error {
	return Add(cmd, uid, ActionDenied, "", "", TypeUser, ReasonInvaild(uid))
}
