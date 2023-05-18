# -*- encoding: utf-8 -*-
"""
@File   : longest_substring_without_repeating_characters.py
@Time   : 2023-05-17 10:10:55
@Author : pan
"""


def lengthOfLongestSubstring(s):
    charMap = {}
    for i in range(256):
        charMap[i] = -1
    ls = len(s)
    i = max_len = 0
    for j in range(ls):
        # Note that when charMap[ord(s[j])] >= i, it means that there are
        # duplicate character in current i,j. So we need to update i.
        if charMap[ord(s[j])] >= i:
            i = charMap[ord(s[j])] + 1
        charMap[ord(s[j])] = j
        max_len = max(max_len, j - i + 1)
    return max_len


if __name__ == "__main__":
    res = lengthOfLongestSubstring("sdfsdssbdfdsfd")
    print(res)
