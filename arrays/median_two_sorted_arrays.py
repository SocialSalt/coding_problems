import sys


def median(nums):
    if len(nums) % 2 == 0:
        med = int(len(nums) / 2)
        return (nums[med] + nums[med - 1]) / 2
    else:
        return nums[int(len(nums) / 2)]


def findMedianSortedArrays(nums1, nums2):
    if nums1 == []:
        return median(nums2)
    if nums2 == []:
        return median(nums1)

    if len(nums1) > len(nums2):
        nums1, nums2 = nums2, nums1

    m = len(nums1)
    n = len(nums2)

    lower = 0
    upper = m

    while lower <= upper:

        med1 = (upper + lower) // 2
        med2 = (m + n + 1) // 2 - med1

        if med1 == 0:
            maxmin1 = -sys.maxsize
        else:
            maxmin1 = nums1[med1 - 1]

        if med2 == 0:
            maxmin2 = -sys.maxsize
        else:
            maxmin2 = nums2[med2 - 1]

        if med1 == m:
            minmax1 = sys.maxsize
        else:
            minmax1 = nums1[med1]

        if med2 == n:
            minmax2 = sys.maxsize
        else:
            minmax2 = nums2[med2]

        if maxmin1 <= minmax2 and maxmin2 <= minmax1:
            if (m + n) % 2 == 0:
                return (max(maxmin1, maxmin2) + min(minmax1, minmax2)) / 2
            else:
                return max(maxmin1, maxmin2)
        elif maxmin1 > minmax2:
            upper = med1 - 1
        else:
            lower = med1 + 1


l1 = [1, 4, 5, 8, 10, 12]
l2 = [2, 2, 2, 3, 6, 9, 32]

print(findMedianSortedArrays(l1, l2))
print()

l1 = [1]
l2 = [2, 2]

print(findMedianSortedArrays(l1, l2))
print()

l1 = [1, 2]
l2 = [3]

print(findMedianSortedArrays(l1, l2))
print()

l1 = [3]
l2 = [-2, -1]

print(findMedianSortedArrays(l1, l2))
print()

l1 = [0, 0, 0, 0, 0]
l2 = [-1, 0, 0, 0, 0, 0, 1]

print(findMedianSortedArrays(l1, l2))
print()

l1 = [1, 2]
l2 = [-1, 3]

print(findMedianSortedArrays(l1, l2))
print()
