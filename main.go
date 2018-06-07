package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nmmh/magneturi/magneturi"
)

func main() {
	var (
		rawMagnetURI string
		softParse    bool
	)
	flag.StringVar(&rawMagnetURI, "rawmagnet", "magnet:?xt.1=urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1&xt.2=urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY&xt.3=urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q&xl=10826029&dn=mediawiki-1.15.1.tar.gz&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&as=http%3A%2F%2Fdownload.wikimedia.org%2Fmediawiki%2F1.15%2Fmediawiki-1.15.1.tar.gz&xs=http%3A%2F%2Fcache.example.org%2FXRX2PEFXOOEJFRVUCX6HMZMKS5TWG4K5&xs=dchub://example.org&x.Moz11=test",
		"The raw magnet uri to parse")
	flag.BoolVar(&softParse, "softparse", false, "specify softparse = true to progress discarding invalid parameters otherwise invalid parameters result in errors")
	flag.Parse()

	//sample code - see test for further features

	m, err := magneturi.Parse(rawMagnetURI, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
	//pretty print some info
	m.PrintVerbose()

	//Verify that certain parameter prefixes exist before filtering for them
	if m.HasPrefixes("xt", "dn", "tr") {
		//new magnet (m2) = filter m for prefixes
		m2, err := m.Filter("xt", "dn", "tr")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m2)
		m2.PrintVerbose()

	} else {
		fmt.Println("The magnet did not have the required paramater type(s) yadda")
	}

}
