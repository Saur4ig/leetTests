package stringsE

import (
	"strings"
)

// 1108. Defanging an IP Address
// Given a valid (IPv4) IP address, return a defanged version of that IP address.
// A defanged IP address replaces every period "." with "[.]".
// link - https://leetcode.com/problems/defanging-an-ip-address/
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
