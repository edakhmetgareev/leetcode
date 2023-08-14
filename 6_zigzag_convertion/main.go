package main

import "strings"

func main() {
	result := convert("PAYPALISHIRING", 4)
	println(result)
}

func convert(s string, numRows int) string {
	if numRows > 1000 || len(s) > 1000 {
		return ""
	}

	if numRows < 2 {
		return s
	}

	sign := true
	cnt := -1
	rows := make([]string, numRows)
	for i := 0; i < len(s); i++ {
		if i > 1 && (cnt == 0 || cnt >= numRows-1) {
			sign = !sign
		}

		if sign {
			cnt++
		} else {
			cnt--
		}

		rows[cnt] += string(s[i])
	}

	var res strings.Builder
	for i := 0; i < numRows; i++ {
		res.WriteString(rows[i])
	}

	return res.String()
}
