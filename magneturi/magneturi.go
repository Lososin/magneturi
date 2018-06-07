package magneturi

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

const (
	magnetSchemaPrefix = "magnet:?"
)

var paramType = func() map[string]string {
	return map[string]string{
		"xt": "exactTopic",
		"dn": "displayName",
		"kt": "keywordTopic",
		"mt": "manifestTopic",
		"tr": "tracker",
		"xs": "exactSource",
		"as": "acceptableSource",
		"xl": "exactLength",
		"x.": "experimental",
	}
}

//MagnetURI this is the type that will containg the parsed MagnetURI.
type MagnetURI struct {
	params []param
}

//param is a paramter that makes up the MagnetURI
type param struct {
	prefix string
	index  string
	value  string
}

//Parse returns a magnet url or fails to parse.
//softparse == true will continue on error extracting all VALID parameters
func Parse(rawMagnetURI string, softParse bool) (*MagnetURI, error) {
	m := &MagnetURI{}
	if strings.HasPrefix(rawMagnetURI, magnetSchemaPrefix) {
		magnetNoSchemaPrefix := strings.TrimPrefix(rawMagnetURI, magnetSchemaPrefix)
		params := strings.Split(magnetNoSchemaPrefix, "&")
		for _, param := range params {
			validParam, err := parseParam(param)
			if err != nil {
				if softParse {
					//skip adding this invalid parameter
					err = nil
					continue
				}
				return m, err
			}
			//add valid parameter to the MagnetURI
			if err := m.addParam(validParam); err != nil {
				if softParse {
					err = nil
					continue
				}
				return m, err
			}
		}
		return m, nil
	}
	return m, fmt.Errorf("uri doesn't start with the Magnet URI schema prefix %q", magnetSchemaPrefix)
}

func parseParam(parameter string) (param, error) {
	paramSplit := strings.SplitN(parameter, "=", 2)
	if len(paramSplit) != 2 || (len(paramSplit) == 2 && paramSplit[1] == "") {
		return param{}, fmt.Errorf("parameter without prefix or prefix without parameter: %q", parameter)
	}
	prefix := paramSplit[0]
	prefix, index, err := splitDotPrefix(prefix)
	if err != nil {
		return param{}, err
	}
	if !isValidPrefix(prefix) {
		return param{}, fmt.Errorf("invalid parameter prefix: %q", prefix)
	}
	value := paramSplit[1]
	return param{prefix, index, value}, nil
}

func splitDotPrefix(prefix string) (string, string, error) {
	if strings.HasPrefix(prefix, "x.") {
		exp := strings.TrimLeft(prefix, "x.")
		if exp == "" {
			return "", "", fmt.Errorf("experimental info missing: %q", prefix)
		}
		return "x.", exp, nil
	} else if strings.Contains(prefix, ".") {
		prefixSplit := strings.SplitN(prefix, ".", 2)
		if len(prefixSplit) != 2 || (len(prefixSplit) == 2 && prefixSplit[1] == "") {
			return "", "", fmt.Errorf("dot index missing: %q", prefix)
		}
		index := prefixSplit[1]
		return prefixSplit[0], index, nil
	}
	return prefix, "", nil
}

//Filter extracts a new MagnetURI based on the supplied prefixes.
// Filters that are based on ParamTypes that are not present will
// return no results.
func (m *MagnetURI) Filter(paramTypes ...string) (*MagnetURI, error) {
	newM := &MagnetURI{}
	for _, pt := range paramTypes {
		filteredParams, _ := m.getParamsByPrefix(pt)
		for _, param := range filteredParams {
			if err := newM.addParam(param); err != nil {
				return newM, err
			}
		}
	}
	return newM, nil
}

func compareParams(first []param, second []param) bool {
	if len(first) == len(second) {
		for _, param := range first {
			if !containsParam(second, param) {
				return false
			}
		}
		return true
	}
	return false
}

func containsParam(list []param, param param) bool {
	for _, item := range list {
		if param.prefix == item.prefix &&
			param.index == item.index &&
			param.value == item.value {
			return true
		}
	}
	return false
}

func isValidPrefix(prefix string) bool {
	if _, ok := paramType()[prefix]; ok {
		return true
	}
	return false
}

func (m *MagnetURI) addParam(validParam param) error {
	if !isValidPrefix(validParam.prefix) {
		return fmt.Errorf("invalid parameter prefix: %q", validParam.prefix)
	}
	m.params = append(m.params, validParam)
	return nil
}

func (m *MagnetURI) getParamsByPrefix(prefix string) ([]param, error) {
	var p []param
	for _, param := range m.params {
		if prefix == param.prefix {
			p = append(p, param)
		}
	}
	if len(p) == 0 {
		return nil, fmt.Errorf("no parameters by getParamsByPrefix using prefix %q", prefix)
	}
	return p, nil
}

//HasPrefix returns false if the magneturi has not got a parameter
// with the passed prefix.
func (m *MagnetURI) HasPrefix(prefix string) bool {
	for _, item := range m.params {
		if prefix == item.prefix {
			return true
		}
	}
	return false
}

//HasPrefixes return false if it does not have a parameter with
// one of the supplied prefixes
func (m *MagnetURI) HasPrefixes(paramTypes ...string) bool {
	for _, pt := range paramTypes {
		if !m.HasPrefix(pt) {
			return false
		}
	}
	return true
}

//Equal returns true if a magneturl has the same parameters
func (m *MagnetURI) Equal(x MagnetURI) bool {
	return compareParams(m.params, x.params)
}

// Asembles from struct
func (m *MagnetURI) String() string {
	if len(m.params) == 0 {
		return "the Magnet URI has no parameters"
	}
	var ret string
	for _, p := range m.params {
		if p.index != "" {
			ret += "&" + strings.TrimRight(p.prefix, ".") + "." + p.index + "=" + p.value
		} else {
			ret += "&" + p.prefix + "=" + p.value
		}
	}
	s := fmt.Sprintf("%s%s", magnetSchemaPrefix, strings.TrimLeft(ret, "&"))
	return s
}

//PrintVerbose pretty prints some info.
func (m *MagnetURI) PrintVerbose() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(tw, "#\tPrefix\tIndex/Exp\tDescription\tValue")
	fmt.Fprintln(tw, "=\t======\t=========\t===========\t=====")
	for i, p := range m.params {
		fmt.Fprintln(tw, strconv.Itoa(i)+"\t"+p.prefix+"\t"+p.index+"\t"+paramType()[p.prefix]+"\t"+p.value)
	}
	tw.Flush()
}
