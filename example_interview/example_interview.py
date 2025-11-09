"""
. . . .
. . . .
. . . .
. . . .
S . . E
"""

"""
only options
right = 0
up right = 1
down right = 2
"""
import numpy as np


def count_ways(height: int, width: int, points: list[tuple[int, int]]) -> int:

    arr = np.zeros((height, width), dtype=np.int64)
    arr[height - 1, width - 1] = 1

    def is_valid(point: tuple[int, int]) -> bool:
        for p in points:
            if p[0] == point[0] and p[1] == point[1]:
                continue
            if point[1] < p[1]:
                continue
            vec = abs(p[0] - point[0]), abs(p[1] - point[1])
            if vec[0] == 0 or vec[1] == 0:
                return False
            if vec[0] / vec[1] < 1:
                return False
        return True

    for col in range(width - 2, -1, -1):
        for row in range(height - 1, -1, -1):
            if not is_valid((row, col)):
                continue
            val = 0
            if row + 1 < height:
                val += arr[row + 1, col + 1]
            if row - 1 > 0:
                val += arr[row - 1, col + 1]
            val += arr[row, col + 1]
            arr[row, col] = val
    print(arr)
    return arr[height - 1, 0]


print(count_ways(8, 8, [(6, 3)]))
