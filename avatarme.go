package avatar

import (
	"crypto/sha512"
	"fmt"
	"io"
	"strings"
)

type Avatar struct {
	emailAddress string
	ipAddress    string
	publicKey    string
	hash         string
}

func (avatar *Avatar) Hash() string {
	if avatar.hash == "" {
		avatar.hash = avatar.createHash()
	}
	return avatar.hash
}

func (a *Avatar) hashMessage() string {
	str := []string{a.emailAddress, a.ipAddress, a.publicKey}
	return strings.Join(str, "")
}

func (a *Avatar) createHash() string {
	h := sha512.New()

	io.WriteString(h, a.hashMessage())
	return fmt.Sprintf("%x", h.Sum(nil))
}
