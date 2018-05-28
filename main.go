package main

import (
	"fmt"
	"log"

	"github.com/nmmh/magneturi/magneturi"
)

/*
Credit these:
https://github.com/ngaut/goTorrent/blob/master/uri.go
https://github.com/elopio/magneturi/blob/master/magneturi.go
https://github.com/Donearm/m2t/blob/master/m2t.go

*/
func main() {

	s := "magnet:?xt=urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1&xt=urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY&xt=urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q&xl=10826029&dn=mediawiki-1.15.1.tar.gz&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&as=http%3A%2F%2Fdownload.wikimedia.org%2Fmediawiki%2F1.15%2Fmediawiki-1.15.1.tar.gz&xs=http%3A%2F%2Fcache.example.org%2FXRX2PEFXOOEJFRVUCX6HMZMKS5TWG4K5&xs=dchub://example.org"
	err := magneturi.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	m, err := magneturi.New(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
	m.Info()
}
