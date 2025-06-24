import heapq
from collections import Counter

def leastInterval(tasks, n):
    freq = Counter(tasks)
    # Python 默认 heapq 是最小堆，用负数实现最大堆
    max_heap = [-cnt for cnt in freq.values()]
    heapq.heapify(max_heap)

    time = 0
    while max_heap:
        temp = []
        for i in range(n + 1):
            if max_heap:
                cnt = heapq.heappop(max_heap)
                if cnt + 1 < 0:  # 还没做完，再次加入
                    temp.append(cnt + 1)
            time += 1
            if not max_heap and not temp:
                break

        for item in temp:
            heapq.heappush(max_heap, item)

    return time