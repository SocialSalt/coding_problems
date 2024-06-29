import sys
import collections

MAX = sys.maxsize


def shortestSubarray(nums, k):
    n = len(nums)
    P = [0]
    for num in nums:
        P.append(P[-1] + num)


def shortestSubarray(A, K):
    N = len(A)
    P = [0]
    for x in A:
        P.append(P[-1] + x)

    # Want smallest y-x with Py - Px >= K
    ans = N + 1  # N+1 is impossible
    monoq = collections.deque()  # opt(y) candidates, represented as indices of P
    print(P)
    for y, Py in enumerate(P):
        print(monoq)
        # Want opt(y) = largest x with Px <= Py - K
        while monoq and Py <= P[monoq[-1]]:
            monoq.pop()

        while monoq and Py - P[monoq[0]] >= K:
            ans = min(ans, y - monoq.popleft())
            print(ans)

        monoq.append(y)

    return ans if ans < N + 1 else -1


nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14]
k = 27
print(shortestSubarray(nums, k), 2)
print()


nums = [17, 85, 93, -45, -21]
k = 150
print(shortestSubarray(nums, k), 2)
print()


nums = [84, -37, 32, 40, 95]
k = 167
print(shortestSubarray(nums, k))
print()


nums = [1, 2]
k = 4
print(shortestSubarray(nums, k), -1)
print()


nums = [2, -1, 2]
k = 3
print(shortestSubarray(nums, k), 3)
print()

nums = [-28, 81, -20, 28, -29]
k = 89
print(shortestSubarray(nums, k), 3)
print()


nums = [56, -21, 56, 35, -9]
k = 61
print(shortestSubarray(nums, k), 2)
print()


nums = [41, 82, -2, -44, -19, -33, -14, 75, 39, -7]
k = 58
print(shortestSubarray(nums, k))
print()


nums = [
    -34,
    37,
    51,
    3,
    -12,
    -50,
    51,
    100,
    -47,
    99,
    34,
    14,
    -13,
    89,
    31,
    -14,
    -44,
    23,
    -38,
    6,
]
k = 151
print(shortestSubarray(nums, k), 2)
print()

nums = [27, 20, 79, 87, -36, 78, 76, 72, 50, -26]
k = 453
print(shortestSubarray(nums, k), 9)
