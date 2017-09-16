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
	idx := IndexComment(line)
	if idx == -1 {
		newLine = line
	} else {
		newLine = line[:idx]
	}

	return newLine
}

func IndexComment(line []byte) int {
	idx := 0
	for idx != -1 {
		i := bytes.IndexAny(line[idx:], `#`)
		if i == -1 {
			//fmt.Printf("[IndexComment]ret -1\n")
			return -1
		}
		//fmt.Printf("[IndexComment]i:%d\n", i)
		idx = idx + i
		if idx >= len(line) {
			//fmt.Printf("[IndexComment]over len\n")
			return -1
		}

		//fmt.Printf("[IndexComment]idx:%d\n", idx)
		dqCount := bytes.Count(line[:idx], []byte(`"`))
		str_dqCount := bytes.Count(line[:idx], []byte(`\"`))
		// 1.remove the count of " which as part of string
		dqCount = dqCount - str_dqCount
		//fmt.Printf("[IndexComment]dq count:%d\n", dqCount)
		if dqCount%2 == 0 {
			// 2.before this #, if there are two real ", then this # match
			//fmt.Printf("[IndexComment]ret idx:%d\n", idx)
			//fmt.Printf("[IndexComment]%s\n", line[:idx])
			return idx
		} else {

			// 3.find next #
			idx++
		}
	}
	//fmt.Printf("[IndexComment]ret -1\n")
	return -1
}
