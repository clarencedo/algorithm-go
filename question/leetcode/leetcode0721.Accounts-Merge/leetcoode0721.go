package leetcode

import (
	"sort"
)

// 定义并查集结构体
type UnionFind struct {
	parent map[string]string
}

//	 例子：
// map[string]string{
//     "johnsmith@mail.com": "johnsmith@mail.com",
//     "john00@mail.com":    "johnsmith@mail.com",
//     "john_newyork@mail.com": "johnsmith@mail.com",
//     "johnnybravo@mail.com": "johnnybravo@mail.com",
//     "mary@mail.com": "mary@mail.com",
// }

// 初始化
func NewUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[string]string),
	}
}

// 找到某个节点的最终父节点（带路径压缩）
func (uf *UnionFind) Find(x string) string {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

// 合并两个节点
func (uf *UnionFind) Union(x, y string) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		uf.parent[rootY] = rootX // 随便选一个当父亲
	}
}

// 主函数
func accountsMerge(accounts [][]string) [][]string {
	uf := NewUnionFind()
	emailToName := make(map[string]string)

	// 第一步：构建邮箱之间的连接关系
	for _, account := range accounts {
		name := account[0]
		for i := 1; i < len(account); i++ {
			email := account[i]
			if _, exists := uf.parent[email]; !exists {
				uf.parent[email] = email // 初始父节点是自己
			}
			emailToName[email] = name
			uf.Union(account[1], email) // 把同一个账户的邮箱连起来
		}
	}

	// 第二步：按照父节点把邮箱归类
	rootToEmails := make(map[string][]string)
	for email := range emailToName {
		root := uf.Find(email) // 找到最终父节点
		rootToEmails[root] = append(rootToEmails[root], email)
	}

	// 第三步：组装结果
	var result [][]string
	for root, emails := range rootToEmails {
		sort.Strings(emails) // 邮箱排序
		name := emailToName[root]
		account := append([]string{name}, emails...)
		result = append(result, account)
	}

	return result
}
