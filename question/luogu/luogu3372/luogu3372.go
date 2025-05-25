package luogu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAXN = 100005
const MAXT = 4 * MAXN

var (
	n       int
	a       = make([]int64, MAXN)
	d       = make([]int64, MAXT) // 线段树节点的和
	b       = make([]int64, MAXT) // 懒标记
	scanner = bufio.NewScanner(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func readInts() []int {
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	nums := make([]int, len(fields))
	for i, f := range fields {
		nums[i], _ = strconv.Atoi(f)
	}
	return nums
}

func readInt64s() []int64 {
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	nums := make([]int64, len(fields))
	for i, f := range fields {
		nums[i], _ = strconv.ParseInt(f, 10, 64)
	}
	return nums
}

func build(l, r, p int) {
	if l == r {
		d[p] = a[l]
		return
	}
	m := (l + r) >> 1
	build(l, m, p<<1)
	build(m+1, r, p<<1|1)
	d[p] = d[p<<1] + d[p<<1|1]
}

func pushDown(p, s, t int) {
	if b[p] != 0 {
		m := (s + t) >> 1
		d[p<<1] += b[p] * int64(m-s+1)
		d[p<<1|1] += b[p] * int64(t-m)
		b[p<<1] += b[p]
		b[p<<1|1] += b[p]

		b[p] = 0
	}
}

func update(l, r int, c int64, s, t, p int) {
	if l <= s && t <= r {
		d[p] += c * int64(t-s+1)
		b[p] += c
		return
	}
	pushDown(p, s, t)
	m := (s + t) >> 1
	if l <= m {
		update(l, r, c, s, m, p<<1)
	}
	if r > m {
		update(l, r, c, m+1, t, p<<1|1)
	}
	d[p] = d[p<<1] + d[p<<1|1]
}

func getSum(l, r, s, t, p int) int64 {
	if l <= s && t <= r {
		return d[p]
	}
	pushDown(p, s, t)
	m := (s + t) >> 1
	var sum int64
	if l <= m {
		sum += getSum(l, r, s, m, p<<1)
	}
	if r > m {
		sum += getSum(l, r, m+1, t, p<<1|1)
	}
	return sum
}

func main() {
	scanner.Split(bufio.ScanLines)

	// 读取 n 和 q
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	n, _ = strconv.Atoi(parts[0])
	q, _ := strconv.Atoi(parts[1])

	// 读取数组
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	for i := 1; i <= n; i++ {
		a[i], _ = strconv.ParseInt(fields[i-1], 10, 64)
	}

	build(1, n, 1)

	for i := 0; i < q; i++ {
		scanner.Scan()
		args := strings.Fields(scanner.Text())
		if args[0] == "2" {
			l, _ := strconv.Atoi(args[1])
			r, _ := strconv.Atoi(args[2])
			fmt.Fprintln(writer, getSum(l, r, 1, n, 1))
		} else {
			l, _ := strconv.Atoi(args[1])
			r, _ := strconv.Atoi(args[2])
			c, _ := strconv.ParseInt(args[3], 10, 64)
			update(l, r, c, 1, n, 1)
		}
	}
	writer.Flush()
}
