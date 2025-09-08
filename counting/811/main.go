package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	domainFreq := make(map[string]int)

	for _, d := range cpdomains {
		words := strings.Split(d, " ")
		times, _ := strconv.Atoi(words[0])
		domainStr := words[1]

		parts := strings.Split(domainStr, ".")
		for i := 0; i < len(parts); i++ {
			domain := strings.Join(parts[i:], ".")
			domainFreq[domain] += times
		}
	}

	res := make([]string, 0, len(domainFreq))
	for domain, count := range domainFreq {
		res = append(res, fmt.Sprintf("%d %s", count, domain))
	}
	return res
}

func main() {
	cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	fmt.Println(subdomainVisits(cpdomains))
}
