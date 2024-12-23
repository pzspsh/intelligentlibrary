# -*- encoding: utf-8 -*-
"""
@File   : threesum.py
@Time   : 2023-05-17 10:17:56
@Author : pan
"""


def threeSum(nums):
    res = []
    nums.sort()
    ls = len(nums)
    for i in range(ls - 2):
        if i > 0 and nums[i] == nums[i - 1]:
            continue
        j = i + 1
        k = ls - 1
        while j < k:
            curr = nums[i] + nums[j] + nums[k]
            if curr == 0:
                res.append([nums[i], nums[j], nums[k]])
                while j < k and nums[j] == nums[j + 1]:
                    j += 1
                while j < k and nums[k] == nums[k - 1]:
                    k -= 1
                j += 1
                k -= 1
            elif curr < 0:
                j += 1
            else:
                k -= 1
    return res


if __name__ == "__main__":
    l1 = [-1, 0, 1, 2, -1, -4]
    res = threeSum(l1)
    print(res)
