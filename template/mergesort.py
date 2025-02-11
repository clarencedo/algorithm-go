from typing import List


def mergeSort(arr: List[int]) -> List[int]:
    if len(arr) <= 1:
        return arr
    mid = len(arr) >> 1
    left = mergeSort(arr[:mid])
    right = mergeSort(arr[mid:])

    return merge(left, right)


def merge(left: List[int], right: List[int]) -> List[int]:
    ans = []
    i, j = 0, 0

    while i < len(left) and j < len(right):
        if left[i] < right[j]:
            ans.append(left[i])
            i += 1
        else:
            ans.append(right[j])
            j += 1

    ans.extend(left[i:])
    ans.extend(right[j:])

    return ans
