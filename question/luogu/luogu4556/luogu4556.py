import sys
sys.setrecursionlimit(1 << 25)
from collections import defaultdict, Counter

n, m = map(int, input().split())

# 构建树结构
tree = [[] for _ in range(n + 1)]
for _ in range(n - 1):
    a, b = map(int, input().split())
    tree[a].append(b)
    tree[b].append(a)

# 预处理：深度、父节点
LOG = 17  # 因为 2^17 > 1e5
fa = [[0] * (LOG + 1) for _ in range(n + 1)]
depth = [0] * (n + 1)

def dfs(u, parent):
    fa[u][0] = parent
    for i in range(1, LOG + 1):
        fa[u][i] = fa[fa[u][i - 1]][i - 1]
    for v in tree[u]:
        if v != parent:
            depth[v] = depth[u] + 1
            dfs(v, u)

dfs(1, 0)

def lca(u, v):
    if depth[u] < depth[v]:
        u, v = v, u
    for i in range(LOG, -1, -1):
        if depth[fa[u][i]] >= depth[v]:
            u = fa[u][i]
    if u == v:
        return u
    for i in range(LOG, -1, -1):
        if fa[u][i] != fa[v][i]:
            u = fa[u][i]
            v = fa[v][i]
    return fa[u][0]

# 差分数组：对每个节点记录 z 粮食类型的增减情况
diff = [defaultdict(int) for _ in range(n + 1)]

for _ in range(m):
    x, y, z = map(int, input().split())
    l = lca(x, y)
    diff[x][z] += 1
    diff[y][z] += 1
    diff[l][z] -= 1
    if fa[l][0]:
        diff[fa[l][0]][z] -= 1

# 合并子树信息，得到每个节点最终的粮食统计
res = [0] * (n + 1)

def dfs2(u, parent):
    counter = Counter()
    for v in tree[u]:
        if v != parent:
            child_counter = dfs2(v, u)
            for k, v in child_counter.items():
                counter[k] += v
    for k, v in diff[u].items():
        counter[k] += v
    if counter:
        res[u] = min([k for k, v in counter.items() if v == max(counter.values())])
    return counter

dfs2(1, 0)

# 输出
for i in range(1, n + 1):
    print(res[i])