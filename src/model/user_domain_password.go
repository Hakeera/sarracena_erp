package model

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptPassword hashes the user's password using MD5.
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
