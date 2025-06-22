import heapq
from typing import List


class Solution:
    def minMeetingRooms(self, intervals: List[List[int]]) -> int:
        if not intervals:
            return 0

        # 1. 先按会议开始时间排序
        intervals.sort(key=lambda x: x[0])

        # 2. 初始化一个最小堆，放第一个会议的结束时间
        heap = [intervals[0][1]]  # heap中存的是end时间

        # 3. 遍历剩下的会议
        for start, end in intervals[1:]:
            # 如果最早结束的会议 <= 当前会议的开始时间，复用会议室（弹出堆顶）
            if heap[0] <= start:
                heapq.heappop(heap)
            # 无论复用还是新开，都加入当前会议的结束时间
            heapq.heappush(heap, end)

        # 4. 最终堆大小 = 所需会议室数量
        return len(heap)
