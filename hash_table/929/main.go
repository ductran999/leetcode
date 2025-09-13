package main

import (
	"fmt"
	"strings"
)

func sanitize(email string) string {
	parts := strings.Split(email, "@")

	// remove all '.' in local part
	local := strings.ReplaceAll(parts[0], ".", "")

	// cut off everything after '+'
	if idx := strings.IndexByte(local, '+'); idx != -1 {
		local = local[:idx]
	}

	return local + "@" + parts[1]
}

func numUniqueEmails(emails []string) int {
	emailSanitized := make(map[string]int)
	for _, email := range emails {
		emailSanitized[sanitize(email)]++
	}

	return len(emailSanitized)
}

func main() {
	emails := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}
	fmt.Println(numUniqueEmails(emails))
}
