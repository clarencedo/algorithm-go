from typing import List


class Solution:
    def dfs(self, grid, row, col):
        if (
            row < 0
            or col < 0
            or row >= len(grid)
            or col >= len(grid[0])
            or grid[row][col] != "1"
        ):
            return

        grid[row][col] = "0"  # 标记已访问

        self.dfs(grid, row - 1, col)
        self.dfs(grid, row + 1, col)
        self.dfs(grid, row, col - 1)
        self.dfs(grid, row, col + 1)

    def numIslands(self, grid: List[List[str]]) -> int:
        if len(grid) == 0:
            return 0

        rows = len(grid)
        cols = len(grid[0])
        count = 0

        for i in range(rows):
            for j in range(cols):
                if grid[i][j] == "1":
                    self.dfs(grid, i, j)
                    count += 1

        return count