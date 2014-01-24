package main

import (
	"bufio"
	"fmt"
	"koding/kite"
	"koding/kite/protocol"
	"os"
)

type Chat struct{}

func (Chat) Inbox(r *protocol.KiteRequest, result *string) error {
	fmt.Printf("[%s recv]: %s\n", r.Name, r.Args.(string))
	*result = "ok"
	return nil
}

func main() {
	k := kite.New("fatih/chat", new(Chat))
	go k.Start()

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		go k.Call("devrim/chat", "Inbox", s.Text(), func(err error, res string) {})
	}
}
