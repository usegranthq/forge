package utils

import (
	"bufio"
	"embed"
	"strings"
)

//go:embed disposable_email_blocklist.conf
var dfs embed.FS

var disposableList = make(map[string]struct{}, 3500)

func init() {
	f, _ := dfs.Open("disposable_email_blocklist.conf")
	for scanner := bufio.NewScanner(f); scanner.Scan(); {
		disposableList[scanner.Text()] = struct{}{}
	}
	f.Close()
}

func IsDisposableEmail(email string) (disposable bool) {
	segs := strings.Split(email, "@")
	if len(segs) < 2 {
		return false
	}

	_, disposable = disposableList[strings.ToLower(segs[len(segs)-1])]
	return
}
