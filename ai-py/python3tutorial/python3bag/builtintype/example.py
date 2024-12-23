# -*- encoding: utf-8 -*-
"""
@File   : example.py
@Time   : 2023-12-14 14:05:01
@Author : pan
"""

# 创建一个空列表
empty_list = []

# 创建一个包含整数的列表
numbers = [1, 2, 3, 4, 5]

# 创建一个包含字符串的列表
fruits = ["apple", "banana", "orange"]

# 创建一个包含混合数据类型的列表
mixed = [1, "apple", 3.14, True]

print(empty_list)  # 输出：[]
print(numbers)  # 输出：[1, 2, 3, 4, 5]
print(fruits)  # 输出：['apple', 'banana', 'orange']
print(mixed)  # 输出：[1, 'apple', 3.14, True]

list1 = ["123", "345"]
list1[0] = "Chinese"  # 修改 list1 索引值为0 的元素
print(list1[0])  # 打印一下

list1 = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]  # 创建列表
print(list1[0:5])  # 获取前5个元素
print(list1[:4])  # 输出前4个元素
print(list1[4:])  # 输出第4个元素以后的所有元素
print(list1[:])  # 从开始到结尾，输出所有元素
print(list[0:9:2])  # 步长为2 每遍历2个元素取出1个
print(list1[::-1])  # 列表分片 从开始到结束 步长为-1 相当于列表反向输出

a = [1, 2, 3]  #
b = [4, 5, 6]  #
print(a + b)  # 拼接两个列表
print(a * 3)

list_ = list("helloWorld!")
print(list_)


fruits = ["orange", "apple", "pear", "apple"]
fruits.append("grape")
print(fruits)  # 末尾添加
print(id(fruits))  # 看看地址

#
fruits.insert(0, "potato")
print(fruits)
print(id(fruits))
fruits.insert(1, "banana")
print(fruits)
print(id(fruits))

#
nums = [1, 2, 3, 4]
fruits.extend(nums)
print(fruits)
print(id(fruits))


fruits = ["orange", "apple", "pear", "banana"]
fruits.pop()  # 弹出并返回尾部元素（默认索引为-1）
print(fruits)
# print(fruits.pop())

#
fruits.pop(2)  # 弹出并返回索引为2的元素
print(fruits)

#
fruits.remove("apple")
print(fruits)

#
fruits.clear()
print(fruits)
#
fruits = ["orange", "apple", "pear", "banana"]
del fruits[3]  # 利用内置函数del() 删除列表中索引值为3的元素


fruits = ["orange", "apple", "pear", "banana", "apple"]
print(fruits.count("apple"))  # 元素apple在列表中出现的次数
print(fruits.index("apple"))  # 元素apple在列表中首次出现的索引值
print("pear" in fruits)  # pear 在列表中吗?
print("tomato" not in fruits)  # tomato 不在列表中吗？


fruits = ["orange", "apple", "pear", "banana", "apple"]
fruits.sort()  # 字典顺序排序
print(fruits)

fruits = ["orange", "apple", "pear", "banana", "apple"]
fruits.reverse()  # 逆序排序
print(fruits)

seasons = ["Spring", "Summer", "Fall", "Winter"]  # 构建一个列表
temp = enumerate(seasons)  # 枚举 seasons 及其元素的索引
print(temp)
a = list(temp)  # 转换为列表 则可以直接输出(默认下标从0开始)
print(a)
b = list(enumerate(seasons, start=1))  # 枚举索引从 1 开始
print(b)

tup1 = ()  # 创建空元组
print(type(tup1))

tup2 = "a", "b", 2, 100  # 定义一个没有括号的元组
print(type(tup2))


tup2 = "a", "b", 2, 100  # 定义一个没有括号的元组
print(type(tup2))
a = tup2[:2]  # 获取元组中前两个元素

a = (1, 2, 3)
b = (4, 5, 6)
c = a + b
print(a, b, c)


alist = [11, 22, 33]  # 定义一个列表
atuple = tuple(alist)  # 将列表转换为元组，此处 tuple 为关键字
print(atuple)

newtuple = tuple("Hello World!")  # 将字符串转换为元组
print(newtuple)  # 验证输出

tup3 = ("语文", "Chemistry", 97, 2.0)  # 将一个元组赋值给变量tup1
print(tup3[1])

tup3[1] = "English"  # 尝试修改元组中索引为1的元素值，失败！

tup4 = ("语文", "Chemistry", 97, 2.0)  # 创建元组
print(id(tup4))  # 查看tup4地址

tup4 = tup4[:2] + ("English",) + tup4[2:]  # 链接元组
print(tup4)
print(id(tup4))


dict1 = {"a": 1, "2020": [1, 2, 3], 100: ("Hello", "World")}
print(dict1)


dict1 = {"a": 1, "2020": [1, 2, 3], 100: ("Hello", "World")}
print(dict1)


age = {"Bob": 29, "Carol": 23, "Alice": 26}  # 定义一个名为 age 字典
a = age.get("Bob")  # 获取键为'Bob'的值
print(a)
b = age.get("Zhang")
print(b)


age = {"Bob": 29, "Carol": 23, "Alice": 26}  # 定义一个名为 age 字典
age["Bob"] = 40  # 将键为 'Bob' 的值改为 40print(age)

age = {"Bob": 29, "Carol": 23, "Alice": 26}  # 定义一个名为 age 字典
age["Zhang"] = 25  # 为字典添加新的键'Zhang' 然后赋值
print(age)


age = {"Bob": 29, "Carol": 23, "Alice": 26}  # 定义一个名为 age 字典
age1 = {"Zhao": 25}  # 更新到旧字典
age.update(age1)
print(age)


age = {"Bob": 29, "Carol": 23, "Alice": 26}  # 定义一个名为 age 字典
print(age.pop("Bob"))  # 弹出键为 'Bob' 的元素，并返回该键对应的值
print(age)


person = {"Name": "Alice", "Age": 11, "Sex": "Female"}
pop_obj = person.popitem()
print(pop_obj)
print(person)


a = {3, 4, 5}  # 创建一个集合a
print(a)
print(type(a))


a = {3, 3, 3}  # 创建一个集合a
print(a)
print(type(a))


list = [1, 3, 5, 5, 7]  # 列表中有重复元素5
a_set = set(list)  # 使用set()函数将列表转换为集合(重复元素被过滤)
print(a_set)


a_set = set([1, 2, 3])  # 列表转换为集合
b_set = {1, 2, 3, 4, 5}  # 直接创建的集合
c = b_set - a_set  # 求差集 数据项在a_set, 但不在b_set
print(c)


a_set = set([1, 2, 3])  # 列表转换为集合
b_set = {1, 2, 3, 3, 4, 5, 6, 7, 98}  # 直接创建的集合
c = a_set ^ b_set  # 求对称差集 : 数据项在a/b中 不会同时出现
print(c)
