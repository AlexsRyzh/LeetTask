package string_task

import (
	"log"
	"strconv"
	"strings"
)

func SubdomainVisits(cpdomains []string) []string {
	domains := map[string]int{}

	for i := 0; i < len(cpdomains); i++ {
		parse_string := strings.Split(cpdomains[i], " ")
		number, err := strconv.Atoi(parse_string[0])
		if err != nil {
			log.Fatal(err)
		}
		domains_str := strings.Split(parse_string[1], ".")
		domain_str := ""
		for j := len(domains_str) - 1; j >= 0; j-- {
			domain_str = domains_str[j] + domain_str
			domains[domain_str] += number
			domain_str = "." + domain_str
		}
	}
	answer := make([]string, 0, len(domains))

	for k, v := range domains {
		answer = append(answer, strconv.Itoa(v)+" "+k)
	}
	return answer
}
