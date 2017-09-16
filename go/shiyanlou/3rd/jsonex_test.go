package jsonex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonEx(t *testing.T) {

	ast := assert.New(t)

	type exampleStruct struct {
		Str string `json:"str"`
		Int int    `json:"int"`
	}

	cases := []struct {
		isjson   bool
		str      string
		expected exampleStruct
	}{
		{
			isjson: true,
			str: `{
					"str": "aaaa", # comment
					"int": 1
				}`,
			expected: exampleStruct{
				Str: "aaaa",
				Int: 1,
			},
		},
		{
			isjson: true,
			str: `{
					"str": "aaaa#", # comment
					"int": 1
				}`,
			expected: exampleStruct{
				Str: "aaaa#",
				Int: 1,
			},
		},
		{
			isjson: true,
			str: `{
					"str": "aaaa" # comment,
					#"int": 1
				}`,
			expected: exampleStruct{
				Str: "aaaa",
				Int: 0,
			},
		},
		{
			isjson: false,
			str: `{
					"str": "aaaa", # comment
					#"int": 1
				}`,
		},
		{
			isjson: true,
			str: `{
					"str": "aaa\"a",
					"int": 1
				}`,
			expected: exampleStruct{
				Str: "aaa\"a",
				Int: 1,
			},
		},
	}

	for _, cs := range cases {
		var n exampleStruct
		err := Unmarshal([]byte(cs.str), &n)
		if cs.isjson {
			ast.Nil(err)
			ast.Equal(cs.expected, n, cs.str)
		} else {
			ast.NotNil(err)
		}
	}
}

func TestIndexComment(t *testing.T) {
	var tests = []struct {
		Line []byte
		Idx  int
	}{
		{[]byte(`#global setting`), 0},
		{[]byte(`"abc"#123`), 5},
		{[]byte(`"abc#"123#456`), 9},
		{[]byte(`"abc":123456`), -1},
	}

	for _, test := range tests {
		i := IndexComment(test.Line)
		if i != test.Idx {
			t.Errorf("[test] %s:%d, expect %d\n", test.Line, i, test.Idx)
		}

	}

}

/*
func TestTrimCommentsLine(t *testing.T) {
	var Tests = []struct {
		Line []byte
		Want []byte
	}{
		{[]byte(`#global setting`), []byte("")},
		{[]byte(`"abc"#123`), []byte(`"abc"`)},
		{[]byte(`"abc#"123#456`), []byte(`"abc#"123`)},
		{[]byte(`"abc":123456`), []byte(`"abc":123456`)},
		{[]byte(`"str": "aaa\"a"`), []byte(`"str": "aaa\"a"`)},
		{[]byte(`"abc#"123#"456"`), []byte(`"abc#"123`)},
		{[]byte(`"abc#"123"#456"`), []byte(`"abc#"123"#456"`)},
		{[]byte(`#"abc#"`), []byte(``)},
	}

	for _, test := range Tests {
		g := TrimCommentsLine(test.Line)
		if bytes.Equal(g, test.Want) == false {
			t.Errorf("[test] %s, %s, expect %s\n", test.Line, g, test.Want)
		}

	}
}*/
