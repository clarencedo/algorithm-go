

  def numWays(self, n: int, k: int) -> int:
    if n == 0:
      return 0
    if n == 1:
      return k
    if n == 2:
      return k * k
    prev2 = k
    prev1 = k * k
    for i in range(3, n + 1):
      cur = (k - 1) * (prev1 + prev2)
      prev2 = prev1
      prev1 = cur
    return prev1
