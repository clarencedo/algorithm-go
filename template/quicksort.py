from typing import List


def quickSort(arr: List[int]) -> List[int]:
    if len(arr) <= 1:
        return arr

    pivot = arr[len(arr) >> 1]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]

    # for x in arr:
    #     if x < pivot:
    #         left.append(x)
    #     elif x == pivot:
    #         middle.append(x)
    #     else:
    #         right.append(x)

    return quickSort(left) + middle + quickSort(right)