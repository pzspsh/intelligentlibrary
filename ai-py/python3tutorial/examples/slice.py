# -*- encoding: utf-8 -*-
"""
@File   : slice.py
@Time   : 2024-10-08 15:21:05
@Author : pan
"""


def split_list(lst, max_size):
    return [lst[i : i + max_size] for i in range(0, len(lst), max_size)]


def splitrun():
    # 假设data是你的原始数组
    data = list(range(1001))  # 示例数据，生成一个包含1001个元素的列表
    # 调用函数，将数组分成每个小于等于100的元素的小数组
    split_data = split_list(data, 100)
    # 打印结果
    i = 0
    for small_list in split_data:
        i += 1
        print(f"小数组 {i}: {small_list}")


if __name__ == "__main__":
    splitrun()
    pass
