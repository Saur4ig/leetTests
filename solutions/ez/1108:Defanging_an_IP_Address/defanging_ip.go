package _108_Defanging_an_IP_Address

import (
	"strings"
)

func defangIPaddr(address string) string {
	var sb strings.Builder
	for _, char := range address {
		if string(char) == "." {
			sb.WriteString("[.]")
			continue
		}
		sb.WriteString(string(char))
	}
	return sb.String()
}
