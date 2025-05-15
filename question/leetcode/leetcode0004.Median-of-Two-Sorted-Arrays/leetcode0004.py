from typing import List
import heapq

class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        # 方法1：合并后排序 - 时间复杂度 O((m+n)log(m+n))
        nums = nums1 + nums2
        nums.sort()
        length = len(nums)
        if length % 2 == 0:
            return (nums[length // 2 - 1] + nums[length // 2]) / 2
        else:
            return nums[length // 2]
            
    def findMedianSortedArrays_heap(self, nums1: List[int], nums2: List[int]) -> float:
        # 方法2：使用堆 - 需要先将所有元素弹出才能获取中间元素
        heap = []
        for num in nums1:
            heapq.heappush(heap, num)
        for num in nums2:
            heapq.heappush(heap, num)
            
        # 将堆中元素按顺序弹出到一个列表中
        sorted_nums = []
        while heap:
            sorted_nums.append(heapq.heappop(heap))
            
        length = len(sorted_nums)
        if length % 2 == 0:
            return (sorted_nums[length // 2 - 1] + sorted_nums[length // 2]) / 2
        else:
            return sorted_nums[length // 2]
    
    def findMedianSortedArrays_two_heaps(self, nums1: List[int], nums2: List[int]) -> float:
        # 方法3：使用大小堆 - 时间复杂度 O((m+n)log(m+n))
        # 最大堆存储较小的一半元素，最小堆存储较大的一半元素
        # Python默认是最小堆，所以对于最大堆，我们存储负值
        max_heap = []  # 存储较小的一半元素（取负值变成最大堆）
        min_heap = []  # 存储较大的一半元素
        
        # 合并两个数组
        nums = nums1 + nums2
        
        for num in nums:
            # 如果最大堆为空或者当前元素小于等于最大堆堆顶元素的负值，则加入最大堆
            if not max_heap or -max_heap[0] >= num:
                heapq.heappush(max_heap, -num)  # 存负值实现最大堆
            else:
                heapq.heappush(min_heap, num)
            
            # 保持平衡：两个堆的大小差不超过1
            if len(max_heap) > len(min_heap) + 1:
                heapq.heappush(min_heap, -heapq.heappop(max_heap))
            elif len(min_heap) > len(max_heap):
                heapq.heappush(max_heap, -heapq.heappop(min_heap))
        
        # 计算中位数
        if len(max_heap) > len(min_heap):
            # 奇数个元素，中位数是最大堆的堆顶
            return -max_heap[0]
        else:
            # 偶数个元素，中位数是两个堆顶的平均值
            return (-max_heap[0] + min_heap[0]) / 2
            
    def findMedianSortedArrays_optimal(self, nums1: List[int], nums2: List[int]) -> float:
        # 方法4：二分查找 - 时间复杂度 O(log(min(m,n)))
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1
        
        x, y = len(nums1), len(nums2)
        low, high = 0, x
        
        while low <= high:
            partitionX = (low + high) // 2
            partitionY = (x + y + 1) // 2 - partitionX
            
            maxX = float('-inf') if partitionX == 0 else nums1[partitionX - 1]
            minX = float('inf') if partitionX == x else nums1[partitionX]
            
            maxY = float('-inf') if partitionY == 0 else nums2[partitionY - 1]
            minY = float('inf') if partitionY == y else nums2[partitionY]
            
            if maxX <= minY and maxY <= minX:
                # 找到了正确的分割点
                if (x + y) % 2 == 0:
                    return (max(maxX, maxY) + min(minX, minY)) / 2
                else:
                    return max(maxX, maxY)
            elif maxX > minY:
                high = partitionX - 1
            else:
                low = partitionX + 1
