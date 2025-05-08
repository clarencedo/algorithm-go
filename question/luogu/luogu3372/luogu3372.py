import sys
sys.setrecursionlimit(1 << 25)

n = 0
a = []
d = []
b = []

def build(l, r, p):
    if l == r:
        d[p] = a[l]
        return
    m = (l + r) // 2
    build(l, m, p << 1)
    build(m + 1, r, p << 1 | 1)
    d[p] = d[p << 1] + d[p << 1 | 1]

def pushdown(p, s, t):
    m = (s + t) // 2
    if b[p]:
        d[p << 1] += b[p] * (m - s + 1)
        d[p << 1 | 1] += b[p] * (t - m)
        b[p << 1] += b[p]
        b[p << 1 | 1] += b[p]
        b[p] = 0

def update(l, r, c, s, t, p):
    if l <= s and t <= r:
        d[p] += (t - s + 1) * c
        b[p] += c
        return
    pushdown(p, s, t)
    m = (s + t) // 2
    if l <= m:
        update(l, r, c, s, m, p << 1)
    if r > m:
        update(l, r, c, m + 1, t, p << 1 | 1)
    d[p] = d[p << 1] + d[p << 1 | 1]

def getsum(l, r, s, t, p):
    if l <= s and t <= r:
        return d[p]
    pushdown(p, s, t)
    m = (s + t) // 2
    res = 0
    if l <= m:
        res += getsum(l, r, s, m, p << 1)
    if r > m:
        res += getsum(l, r, m + 1, t, p << 1 | 1)
    return res

# 主函数
if __name__ == '__main__':
    import sys
    input = sys.stdin.readline

    n, q = map(int, input().split())
    a = [0] + list(map(int, input().split()))  # 1-indexed
    size = 4 * n + 5
    d = [0] * size
    b = [0] * size

    build(1, n, 1)

    for _ in range(q):
        args = list(map(int, input().split()))
        if args[0] == 2:
            print(getsum(args[1], args[2], 1, n, 1))
        else:
            update(args[1], args[2], args[3], 1, n, 1)