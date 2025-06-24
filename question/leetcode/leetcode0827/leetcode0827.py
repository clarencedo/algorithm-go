from typing import List

class Solution:
    def largestIsland(self, grid: List[List[int]]) -> int:
        n = len(grid)
        island_id = 2
        area_map = {}  # island_id -> area

        def dfs(x, y, id_):
            if x < 0 or x >= n or y < 0 or y >= n or grid[x][y] != 1:
                return 0
            grid[x][y] = id_
            area = 1
            for dx, dy in [(-1,0), (1,0), (0,-1), (0,1)]:
                area += dfs(x + dx, y + dy, id_)
            return area

        # Step 1: Mark islands and record areas
        for i in range(n):
            for j in range(n):
                if grid[i][j] == 1:
                    area = dfs(i, j, island_id)
                    area_map[island_id] = area
                    island_id += 1

        max_area = max(area_map.values(), default=0)  # If no 1 exists

        # Step 2: Try flipping every 0
        for i in range(n):
            for j in range(n):
                if grid[i][j] == 0:
                    seen = set()
                    for dx, dy in [(-1,0), (1,0), (0,-1), (0,1)]:
                        ni, nj = i + dx, j + dy
                        if 0 <= ni < n and 0 <= nj < n and grid[ni][nj] > 1:
                            seen.add(grid[ni][nj])
                    new_area = 1 + sum(area_map[island] for island in seen)
                    max_area = max(max_area, new_area)

        return max_area