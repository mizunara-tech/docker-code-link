package hexutil

import "fmt"

func HexEncode(str string) string {
	hexStr := ""
	for _, r := range str {
		hexStr += fmt.Sprintf("%x", r)
	}
	return hexStr
}
