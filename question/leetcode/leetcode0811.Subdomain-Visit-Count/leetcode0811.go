package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	countMap := make(map[string]int)

	for _, cpdomain := range cpdomains {
		parts := strings.Split(cpdomain, " ")
		count, _ := strconv.Atoi(parts[0])
		domain := parts[1]

		// 拆分子域名
		subdomains := strings.Split(domain, ".")
		for i := 0; i < len(subdomains); i++ {
			sub := strings.Join(subdomains[i:], ".")
			countMap[sub] += count
		}
	}

	// 构造结果
	res := []string{}
	for domain, count := range countMap {
		res = append(res, fmt.Sprintf("%d %s", count, domain))
	}

	return res
}