package jsonex

import (
	"bytes"
	"encoding/json"
)

func Unmarshal(data []byte, ret interface{}) error {

	lines := bytes.Split(data, []byte{'\n'})
	var trimed [][]byte
	for _, line := range lines {
		trimed = append(trimed, trimCommentsLine(line))
	}
	data = bytes.Join(trimed, []byte{'\n'})

	return json.Unmarshal(data, ret)
}

func trimCommentsLine(line []byte) []byte {

	var newLine []byte

	// 补全代码
	

	return newLine
}
