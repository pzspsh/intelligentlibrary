# -*- encoding: utf-8 -*-
"""
@File   : towsum.py
@Time   : 2023-05-17 10:03:57
@Author : pan
"""


def twoSum(nums, target):
    # two point
    nums_index = [(v, index) for index, v in enumerate(nums)]
    nums_index.sort()
    begin, end = 0, len(nums) - 1
    while begin < end:
        curr = nums_index[begin][0] + nums_index[end][0]
        if curr == target:
            return [nums_index[begin][1], nums_index[end][1]]
        elif curr < target:
            begin += 1
        else:
            end -= 1
    return []


if __name__ == "__main__":
    res = twoSum([3, 5, 6, 3, 4, 9], 7)
    print(res)
