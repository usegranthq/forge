package utils

import (
	"bufio"
	"embed"
	"strings"

	"github.com/lindell/go-burner-email-providers/burner"
)

//go:embed disposable_email_blocklist.conf
var dfs embed.FS

var disposableList = make(map[string]struct{}, 3500)

type EmailsUtil struct{}

var Emails = &EmailsUtil{}

func (u *EmailsUtil) Init() {
	f, _ := dfs.Open("disposable_email_blocklist.conf")
	for scanner := bufio.NewScanner(f); scanner.Scan(); {
		disposableList[scanner.Text()] = struct{}{}
	}
	f.Close()
}

func (u *EmailsUtil) isInDisposableList(email string) bool {
	segs := strings.Split(email, "@")
	if len(segs) < 2 {
		return false
	}
	_, disposable := disposableList[strings.ToLower(segs[len(segs)-1])]
	return disposable
}

func (u *EmailsUtil) IsDisposableEmail(email string) bool {
	if u.isInDisposableList(email) {
		return true
	}

	return burner.IsBurnerEmail(email)
}
