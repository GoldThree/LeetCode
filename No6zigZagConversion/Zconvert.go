package main

import (
	"bytes"
)

func main() {
	convert("asdftgeqwr", 3)
}

func convert(s string, numRows int) string {

	if numRows == 1 || len(s) <= numRows {

		return s

	}

	res := bytes.Buffer{}

	p := numRows*2 - 2

	for i := 0; i < len(s); i += p {

		res.WriteByte(s[i])

	}

	for r := 1; r <= numRows-2; r++ {

		res.WriteByte(s[r])

		for k := p; k-r < len(s); k += p {

			res.WriteByte(s[k-r])

			if k+r < len(s) {

				res.WriteByte(s[k+r])

			}

		}

	}

	for i := numRows - 1; i < len(s); i += p {

		res.WriteByte(s[i])

	}

	return res.String()

}
