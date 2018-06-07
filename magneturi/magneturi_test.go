package magneturi

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		rawMagnetURI string
		softParse    bool
	}
	tests := []struct {
		name    string
		args    args
		want    *MagnetURI
		wantErr bool
	}{
		{
			name: "all example",
			args: args{
				rawMagnetURI: "magnet:?xt=urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q&xt.1=urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1&xt.2=urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY&xt.3=urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q&xl=10826029&dn=mediawiki-1.15.1.tar.gz&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&as=http%3A%2F%2Fdownload.wikimedia.org%2Fmediawiki%2F1.15%2Fmediawiki-1.15.1.tar.gz&xs=http%3A%2F%2Fcache.example.org%2FXRX2PEFXOOEJFRVUCX6HMZMKS5TWG4K5&xs=dchub://example.org&x.Moz11=test",
				softParse:    false,
			},
			want: &MagnetURI{
				params: []param{
					param{"xt", "", "urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q"},
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
					param{"xt", "3", "urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q"},
					param{"xl", "", "10826029"},
					param{"dn", "", "mediawiki-1.15.1.tar.gz"},
					param{"tr", "", "udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce"},
					param{"as", "", "http%3A%2F%2Fdownload.wikimedia.org%2Fmediawiki%2F1.15%2Fmediawiki-1.15.1.tar.gz"},
					param{"xs", "", "http%3A%2F%2Fcache.example.org%2FXRX2PEFXOOEJFRVUCX6HMZMKS5TWG4K5"},
					param{"xs", "", "dchub://example.org"},
					param{"x.", "Moz11", "test"},
				},
			},
			wantErr: false,
		},
		{
			name: "all softparse example",
			args: args{
				rawMagnetURI: "magnet:?xX=urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q&xt.1=urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1&xt.=urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY&xt.3=&xl=10826029&dn=mediawiki-1.15.1.tar.gz&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&as=http%3A%2F%2Fdownload.wikimedia.org%2Fmediawiki%2F1.15%2Fmediawiki-1.15.1.tar.gz&xs=http%3A%2F%2Fcache.example.org%2FXRX2PEFXOOEJFRVUCX6HMZMKS5TWG4K5&xs=dchub://example.org&x.Moz11=test",
				softParse:    true,
			},
			want: &MagnetURI{
				params: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xl", "", "10826029"},
					param{"dn", "", "mediawiki-1.15.1.tar.gz"},
					param{"tr", "", "udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce"},
					param{"as", "", "http%3A%2F%2Fdownload.wikimedia.org%2Fmediawiki%2F1.15%2Fmediawiki-1.15.1.tar.gz"},
					param{"xs", "", "http%3A%2F%2Fcache.example.org%2FXRX2PEFXOOEJFRVUCX6HMZMKS5TWG4K5"},
					param{"xs", "", "dchub://example.org"},
					param{"x.", "Moz11", "test"},
				},
			},
			wantErr: false,
		},
		{
			name: "kt example",
			args: args{
				rawMagnetURI: "magnet:?kt=martin+luther+king+mp3",
				softParse:    false,
			},
			want: &MagnetURI{
				params: []param{
					param{"kt", "", "martin+luther+king+mp3"},
				},
			},
			wantErr: false,
		},
		{
			name: "mt example",
			args: args{
				rawMagnetURI: "magnet:?mt=http://weblog.foo/all-my-favorites.rss",
				softParse:    false,
			},
			want: &MagnetURI{
				params: []param{
					param{"mt", "", "http://weblog.foo/all-my-favorites.rss"},
				},
			},
			wantErr: false,
		},
		{
			name: "xt example",
			args: args{
				rawMagnetURI: "magnet:?xt.1=urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1&xt.2=urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY&xt.3=urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q&xt=urn:ed2k:31D6CFE0D16AE931B73C59D7E0C089C0",
				softParse:    false,
			},
			want: &MagnetURI{
				params: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
					param{"xt", "3", "urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q"},
					param{"xt", "", "urn:ed2k:31D6CFE0D16AE931B73C59D7E0C089C0"},
				},
			},
			wantErr: false,
		},
		{
			name: "xl example",
			args: args{
				rawMagnetURI: "magnet:?xl=10826029",
				softParse:    false,
			},
			want: &MagnetURI{
				params: []param{
					param{"xl", "", "10826029"},
				},
			},
			wantErr: false,
		},
		{
			name: "tr example",
			args: args{
				rawMagnetURI: "magnet:?tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce",
				softParse:    false,
			},
			want: &MagnetURI{
				params: []param{
					param{"tr", "", "udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce"},
				},
			},
			wantErr: false,
		},
		{
			name: "as example",
			args: args{
				rawMagnetURI: "magnet:?as=http%3A%2F%2Fdownload.wikimedia.org%2Fmediawiki%2F1.15%2Fmediawiki-1.15.1.tar.gz",
				softParse:    false,
			},
			want: &MagnetURI{
				params: []param{
					param{"as", "", "http%3A%2F%2Fdownload.wikimedia.org%2Fmediawiki%2F1.15%2Fmediawiki-1.15.1.tar.gz"},
				},
			},
			wantErr: false,
		},
		{
			name: "xs example",
			args: args{
				rawMagnetURI: "magnet:?xs=http%3A%2F%2Fcache.example.org%2FXRX2PEFXOOEJFRVUCX6HMZMKS5TWG4K5&xs=dchub://example.org",
				softParse:    false,
			},
			want: &MagnetURI{
				params: []param{
					param{"xs", "", "http%3A%2F%2Fcache.example.org%2FXRX2PEFXOOEJFRVUCX6HMZMKS5TWG4K5"},
					param{"xs", "", "dchub://example.org"},
				},
			},
			wantErr: false,
		},
		{
			name: "x. example",
			args: args{
				rawMagnetURI: "magnet:?x.Moz11=test",
				softParse:    false,
			},
			want: &MagnetURI{
				params: []param{
					param{"x.", "Moz11", "test"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.rawMagnetURI, tt.args.softParse)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseWithErrors(t *testing.T) {
	tests := []struct {
		Name          string
		RawMagnetURI  string
		SoftParse     bool
		ExpectedError string
	}{
		{
			Name:          "URI without magnet schema prefix",
			RawMagnetURI:  "mUgnet",
			SoftParse:     false,
			ExpectedError: "uri doesn't start with the Magnet URI schema prefix \"magnet:?\"",
		},
		{
			Name:          "prefix without parameter",
			RawMagnetURI:  "magnet:?xt=",
			SoftParse:     false,
			ExpectedError: "parameter without prefix or prefix without parameter: \"xt=\"",
		},
		{
			Name:          "parameter without prefix",
			RawMagnetURI:  "magnet:?http://wtf",
			SoftParse:     false,
			ExpectedError: "parameter without prefix or prefix without parameter: \"http://wtf\"",
		},
		{
			Name:          "URI with invalid parameter prefix",
			RawMagnetURI:  "magnet:?invalid=value",
			SoftParse:     false,
			ExpectedError: "invalid parameter prefix: \"invalid\"",
		},
		{
			Name:          "URI with DOT parameter no index",
			RawMagnetURI:  "magnet:?xt.=blah",
			SoftParse:     false,
			ExpectedError: "dot index missing: \"xt.\"",
		},
		{
			Name:          "URI with experimental DOT parameter no index",
			RawMagnetURI:  "magnet:?x.=blah",
			SoftParse:     false,
			ExpectedError: "experimental info missing: \"x.\"",
		},
		{
			Name:          "URI with DOT parameter - parameter missing",
			RawMagnetURI:  "magnet:?xt.1=",
			SoftParse:     false,
			ExpectedError: "parameter without prefix or prefix without parameter: \"xt.1=\"",
		},
	}
	for _, test := range tests {
		magnetURI, error := Parse(test.RawMagnetURI, test.SoftParse)
		if !magnetURI.Equal(MagnetURI{}) {
			t.Errorf("Error on test %q: a non-empty Magnet URI was returned: %v.",
				test.Name, magnetURI)
		}
		if error == nil {
			t.Errorf("No error was returned on %s test.", test.Name)
		}
		if error.Error() != test.ExpectedError {
			t.Errorf(
				"Error on test %q: Expected error message: %q; got %q",
				test.Name, test.ExpectedError, error.Error())
		}
	}
}

func TestMagnetURI_addParam(t *testing.T) {
	type args struct {
		validParam param
	}
	tests := []struct {
		name    string
		m       MagnetURI
		p       param
		wantErr bool
	}{
		{
			name: "Add param valid param",
			m: MagnetURI{
				params: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
				},
			},
			p:       param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
			wantErr: false,
		},
		{
			name: "Add param invalid param",
			m: MagnetURI{
				params: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
				},
			},
			p:       param{"ll", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.addParam(tt.p); (err != nil) != tt.wantErr {
				t.Errorf("MagnetURI.addParam() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMagnetURI_String(t *testing.T) {
	tests := []struct {
		name string
		m    MagnetURI
		want string
	}{
		{
			name: "Test the stringer",
			m: MagnetURI{
				params: []param{
					param{"xt", "", "urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q"},
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
					param{"xt", "3", "urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q"},
					param{"xl", "", "10826029"},
					param{"dn", "", "mediawiki-1.15.1.tar.gz"},
					param{"tr", "", "udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce"},
					param{"as", "", "http%3A%2F%2Fdownload.wikimedia.org%2Fmediawiki%2F1.15%2Fmediawiki-1.15.1.tar.gz"},
					param{"xs", "", "http%3A%2F%2Fcache.example.org%2FXRX2PEFXOOEJFRVUCX6HMZMKS5TWG4K5"},
					param{"xs", "", "dchub://example.org"},
					param{"x.", "Moz11", "test"},
				},
			},
			want: "magnet:?xt=urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q&xt.1=urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1&xt.2=urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY&xt.3=urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q&xl=10826029&dn=mediawiki-1.15.1.tar.gz&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce&as=http%3A%2F%2Fdownload.wikimedia.org%2Fmediawiki%2F1.15%2Fmediawiki-1.15.1.tar.gz&xs=http%3A%2F%2Fcache.example.org%2FXRX2PEFXOOEJFRVUCX6HMZMKS5TWG4K5&xs=dchub://example.org&x.Moz11=test",
		},
		{
			name: "Test the stringer - no params",
			m: MagnetURI{
				params: []param{},
			},
			want: "the Magnet URI has no parameters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("MagnetURI.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagnetURI_Filter(t *testing.T) {
	type args struct {
		paramTypes []string
	}
	tests := []struct {
		name    string
		m       *MagnetURI
		args    []string
		want    *MagnetURI
		wantErr bool
	}{
		{
			name: "TestMagnetURI_Filter success",
			m: &MagnetURI{
				params: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
					param{"xt", "3", "urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q"},
					param{"xl", "", "10826029"},
					param{"dn", "", "mediawiki-1.15.1.tar.gz"},
					param{"tr", "", "udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce"},
					param{"as", "", "http%3A%2F%2Fdownload.wikimedia.org%2Fmediawiki%2F1.15%2Fmediawiki-1.15.1.tar.gz"},
					param{"xs", "", "http%3A%2F%2Fcache.example.org%2FXRX2PEFXOOEJFRVUCX6HMZMKS5TWG4K5"},
					param{"xs", "", "dchub://example.org"},
					param{"x.", "Moz11", "test"},
				},
			},
			args: []string{"xt", "dn", "tr"},
			want: &MagnetURI{
				params: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
					param{"xt", "3", "urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q"},
					param{"dn", "", "mediawiki-1.15.1.tar.gz"},
					param{"tr", "", "udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce"},
				},
			},
			wantErr: false,
		},
		{
			name: "TestMagnetURI_Filter success",
			m: &MagnetURI{
				params: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
					param{"xt", "3", "urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q"},
					param{"xl", "", "10826029"},
					param{"dn", "", "mediawiki-1.15.1.tar.gz"},
					param{"tr", "", "udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce"},
					param{"as", "", "http%3A%2F%2Fdownload.wikimedia.org%2Fmediawiki%2F1.15%2Fmediawiki-1.15.1.tar.gz"},
					param{"xs", "", "http%3A%2F%2Fcache.example.org%2FXRX2PEFXOOEJFRVUCX6HMZMKS5TWG4K5"},
					param{"xs", "", "dchub://example.org"},
					param{"x.", "Moz11", "test"},
				},
			},
			args: []string{"xt", "dn", "tr", "kt"},
			want: &MagnetURI{
				params: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
					param{"xt", "3", "urn:btih:QHQXPYWMACKDWKP47RRVIV7VOURXFE5Q"},
					param{"dn", "", "mediawiki-1.15.1.tar.gz"},
					param{"tr", "", "udp%3A%2F%2Ftracker.openbittorrent.com%3A80%2Fannounce"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Filter(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MagnetURI.Filter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MagnetURI.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagnetURI_HasPrefix(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		m      *MagnetURI
		prefix string
		want   bool
	}{
		{
			name: "TestMagnetURI_HasPrefix success",
			m: &MagnetURI{
				params: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
				},
			},
			prefix: "xt",
			want:   true,
		},
		{
			name: "TestMagnetURI_HasPrefix success",
			m: &MagnetURI{
				params: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
				},
			},
			prefix: "dn",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.HasPrefix(tt.prefix); got != tt.want {
				t.Errorf("MagnetURI.HasPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagnetURI_getParamsByPrefix(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name    string
		m       *MagnetURI
		prefix  string
		want    []param
		wantErr string
	}{
		{
			name: "MagnetURIGetParamsByPrefix return vals",
			m: &MagnetURI{
				params: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
					param{"dn", "", "mediawiki-1.15.1.tar.gz"},
				},
			},
			prefix: "xt",
			want: []param{
				param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
				param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
			},
			wantErr: "false",
		},
		{
			name: "MagnetURIGetParamsByPrefix return error",
			m: &MagnetURI{
				params: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
				},
			},
			prefix:  "dn",
			want:    []param{},
			wantErr: "no parameters by getParamsByPrefix using prefix \"dn\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.getParamsByPrefix(tt.prefix)

			if err != nil {
				if !reflect.DeepEqual(err.Error(), tt.wantErr) {
					t.Errorf("MagnetURI.getParamsByPrefix() error = %v, wantErr = %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MagnetURI.getParamsByPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidPrefix(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		prefix string
		want   bool
	}{
		{
			name:   "Test_isValidPrefix success",
			prefix: "xt",
			want:   true,
		},
		{
			name:   "Test_isValidPrefix success",
			prefix: "xX",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPrefix(tt.prefix); got != tt.want {
				t.Errorf("isValidPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsParam(t *testing.T) {
	type args struct {
		list  []param
		param param
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_containsParam",
			args: args{
				[]param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
				},
				param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
			},
			want: true,
		},
		{
			name: "Test_containsParam",
			args: args{
				[]param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
				},
				param{"xX", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsParam(tt.args.list, tt.args.param); got != tt.want {
				t.Errorf("containsParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareParams(t *testing.T) {
	type args struct {
		first  []param
		second []param
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_compareParams success",
			args: args{
				first: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
				},
				second: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
				},
			},
			want: true,
		},
		{
			name: "Test_compareParams fail",
			args: args{
				first: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
				},
				second: []param{
					param{"xt", "1", "urn:ed2k:354B15E68FB8F36D7CD88FF94116CDC1"},
					//param{"xt", "2", "urn:tree:tiger:7N5OAMRNGMSSEUE3ORHOKWN4WWIQ5X4EBOOTLJY"},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareParams(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("compareParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
