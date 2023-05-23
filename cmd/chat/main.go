package main

import (
	"chat"
	"os"
	"strings"
)

/*
build cmd: go build -o q cmd/chat/main.go && sudo mv q /usr/local/bin/
*/

func main() {
	chat.InitFromConf()
	switch len(os.Args) {
	case 1:
		chat.NewChat().Chat()
	default:
		qs := strings.Join(os.Args[1:], " ")
		chat.NewChatOnce().ChatOnce(qs)
	}
}
