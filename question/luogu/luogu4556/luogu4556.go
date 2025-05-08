package luogu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	N = 100005
	M = 5000005
)

var (
	n, m int
	fa   [N][22]int
	dep  [N]int
	rt   [N]int

	sum = [M]int{}
	res = [M]int{}
	ls  = [M]int{}
	rs  = [M]int{}
	cnt = 0

	ans = [N]int{}
	v   = make([][]int, N)
)

func update(x int) {
	if sum[ls[x]] < sum[rs[x]] {
		res[x] = res[rs[x]]
		sum[x] = sum[rs[x]]
	} else {
		res[x] = res[ls[x]]
		sum[x] = sum[ls[x]]
	}
}

func merge(a, b, l, r int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if l == r {
		sum[a] += sum[b]
		return a
	}
	mid := (l + r) >> 1
	ls[a] = merge(ls[a], ls[b], l, mid)
	rs[a] = merge(rs[a], rs[b], mid+1, r)
	update(a)
	return a
}

func add(id, l, r, co, val int) int {
	if id == 0 {
		cnt++
		id = cnt
	}
	if l == r {
		sum[id] += val
		res[id] = co
		return id
	}
	mid := (l + r) >> 1
	if co <= mid {
		ls[id] = add(ls[id], l, mid, co, val)
	} else {
		rs[id] = add(rs[id], mid+1, r, co, val)
	}
	update(id)
	return id
}

func initLCA(x int) {
	for i := 0; i <= 20; i++ {
		fa[x][i+1] = fa[fa[x][i]][i]
	}
	for _, y := range v[x] {
		if y == fa[x][0] {
			continue
		}
		dep[y] = dep[x] + 1
		fa[y][0] = x
		initLCA(y)
	}
}

func lca(x, y int) int {
	if dep[x] < dep[y] {
		x, y = y, x
	}
	d := dep[x] - dep[y]
	for i := 0; d > 0; i++ {
		if d&1 != 0 {
			x = fa[x][i]
		}
		d >>= 1
	}
	if x == y {
		return x
	}
	for i := 20; i >= 0; i-- {
		if fa[x][i] != fa[y][i] {
			x = fa[x][i]
			y = fa[y][i]
		}
	}
	return fa[x][0]
}

func cacl(x int) {
	for _, y := range v[x] {
		if y == fa[x][0] {
			continue
		}
		cacl(y)
		rt[x] = merge(rt[x], rt[y], 1, 100000)
	}
	ans[x] = res[rt[x]]
	if sum[rt[x]] == 0 {
		ans[x] = 0
	}
}

func readInts(scanner *bufio.Scanner, count int) []int {
	res := make([]int, count)
	for i := 0; i < count; i++ {
		scanner.Scan()
		res[i], _ = strconv.Atoi(scanner.Text())
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < n-1; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		v[a] = append(v[a], b)
		v[b] = append(v[b], a)
	}

	initLCA(1)

	for i := 0; i < m; i++ {
		a := readInts(scanner, 3)
		x, y, c := a[0], a[1], a[2]
		rt[x] = add(rt[x], 1, 100000, c, 1)
		rt[y] = add(rt[y], 1, 100000, c, 1)
		t := lca(x, y)
		rt[t] = add(rt[t], 1, 100000, c, -1)
		if fa[t][0] != 0 {
			rt[fa[t][0]] = add(rt[fa[t][0]], 1, 100000, c, -1)
		}
	}

	cacl(1)

	for i := 1; i <= n; i++ {
		fmt.Fprintln(writer, ans[i])
	}
}
