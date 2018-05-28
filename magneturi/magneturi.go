package magneturi

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

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

//MagnetURI this is the structure that will containg the parsed MagnetURI.
type MagnetURI struct {
	params []param
}
type param struct {
	prefix string
	value  string
}

//New - performs some light parsing (garbage is discarded) and extracts the valid data to a Magnet struct.
func New(uri string) (*MagnetURI, error) {
	params := make([]param, 0, 0)
	m := &MagnetURI{params}

	if !strings.HasPrefix(uri, magnetURIPrefix) {
		return m, fmt.Errorf("the following valid Magnet uri schema prefix is not present: %s", magnetURIPrefix)
	}

	u, err := url.Parse(uri)
	if err != nil {
		return m, err
	}

	for prefix := range pType() {
		prefixParams, err := extractParams(prefix, u)
		if err != nil {
			return m, err
		}
		if prefixParams != nil {
			params = append(params, prefixParams...)
		}
	}
	return &MagnetURI{params}, nil
}

func extractParams(prefix string, u *url.URL) ([]param, error) {
	ps, ok := u.Query()[prefix]
	if !ok {
		fmt.Printf("info: Magnet URI does not include parameter: %s\n", prefix)
		return nil, nil
	}
	params := make([]param, 0, len(ps))
	for _, p := range ps {
		params = append(params, param{prefix: prefix, value: p})
	}
	return params, nil
}

// Aseembles from struct - implements stringer
func (m *MagnetURI) String() string {
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

//Info pretty prints some info.
func (m *MagnetURI) Info() {
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "#\tPrefix\tDescription\tValue")
	fmt.Fprintln(w, "=\t======\t===========\t=====")
	for i, p := range m.params {
		fmt.Fprintln(w, strconv.Itoa(i)+"\t"+p.prefix+"\t"+pType()[p.prefix]+"\t"+p.value)
	}
	w.Flush()
}
