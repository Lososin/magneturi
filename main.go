package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

/*
Credit these:
https://github.com/ngaut/goTorrent/blob/master/uri.go
https://github.com/elopio/magneturi/blob/master/magneturi.go
https://github.com/Donearm/m2t/blob/master/m2t.go

*/

const (
	magnetURIPrefix = "magnet:?"
)

var pType = func() map[string]string {
	return map[string]string{
		"xt": "exactTopic",
		"dn": "displayName",
		"kt": "keywordTopic",
		"mt": "manifestTopic",
		"tr": "tracker",
		"xs": "exactSource",
		"as": "acceptableSource",
		"xl": "exactLength",
	}
}

//Magnet this is the structure that will containg the parsed magnet.
type magnet struct {
	params []param
}
type param struct {
	prefix string
	value  string
}

func main() {
	s := "magnet:?urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1&xt=urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY&xt=urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q&xl=10826029&dn=mediawiki-1.15.1.tar.gz&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&as=http%3A%2F%2Fdownload.wikimedia.org%2Fmediawiki%2F1.15%2Fmediawiki-1.15.1.tar.gz&xs=http%3A%2F%2Fcache.example.org%2FXRX2PEFXOOEJFRVUCX6HMZMKS5TWG4K5&xs=dchub://example.org"
	m, err := parseMagnet(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
	m.Info()
}

func parseMagnet(uri string) (*magnet, error) {
	params := make([]param, 0, 0)
	m := &magnet{params}

	if !strings.HasPrefix(uri, magnetURIPrefix) {
		return m, fmt.Errorf("the magnet uri schema prefix is not present: %s", magnetURIPrefix)
	}

	u, err := url.Parse(uri)
	if err != nil {
		return m, err
	}

	for prefix := range pType() {
		//fmt.Printf("Should be a prefix: [%s]\n", prefix)
		prefixParams, err := parseParams(prefix, u)
		if err != nil {
			return m, err
		}
		if prefixParams != nil {
			params = append(params, prefixParams...)
		}
	}
	return &magnet{params}, nil
}

func parseParams(prefix string, u *url.URL) ([]param, error) {
	ps, ok := u.Query()[prefix]
	if !ok {
		fmt.Printf("info: magnet URI does not include parameter: %s\n", prefix)
		return nil, nil //errors.New("magnet URI missing parameter: " + prefix)
	}
	params := make([]param, 0, len(ps))
	for _, p := range ps {
		params = append(params, param{prefix: prefix, value: p})
	}
	return params, nil
}

func (m *magnet) String() string {
	if len(m.params) == 0 {
		return "the Magnet URI has no parameters"
	}
	var ret string
	for _, p := range m.params {
		ret += "&" + p.prefix + "=" + p.value
	}
	s := fmt.Sprintf("%s%s", magnetURIPrefix, strings.TrimLeft(ret, "&"))
	return s
}

func (m *magnet) Info() {
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "#\tPrefix\tMeaning\tValue")
	fmt.Fprintln(w, "=\t======\t=======\t=====")
	for i, p := range m.params {
		fmt.Fprintln(w, strconv.Itoa(i)+"\t"+p.prefix+"\t"+pType()[p.prefix]+"\t"+p.value)
	}
	w.Flush()
}
