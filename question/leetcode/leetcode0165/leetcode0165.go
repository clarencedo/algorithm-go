package leetcode

func compareVersion(version1 string, version2 string) int {
	p1, p2 := 0, 0
	n, m := len(version1), len(version2)
	for p1 < n || p2 < m {
		num1, num2 := 0, 0
		for p1 < n && version1[p1] != '.' {
			num1 = num1*10 + int(version1[p1]-'0')
			p1++
		}
		for p2 < m && version2[p2] != '.' {
			num2 = num2*10 + int(version2[p2]-'0')
			p2++
		}
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
		p1++
		p2++
	}
	return 0
}
