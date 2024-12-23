# -*- encoding: utf-8 -*-
"""
@File   : evalfun.py
@Time   : 2023-12-07 10:00:05
@Author : pan
"""

""" 
eval() 函数用来执行一个字符串表达式，并返回表达式的值。

字符串表达式可以包含变量、函数调用、运算符和其他 Python 语法元素。


以下是 eval() 方法的语法:
eval(expression[, globals[, locals]])
参数
expression -- 表达式。
globals -- 变量作用域，全局命名空间，如果被提供，则必须是一个字典对象。
locals -- 变量作用域，局部命名空间，如果被提供，可以是任何映射对象。
eval() 函数将字符串 expression 解析为 Python 表达式，并在指定的命名空间中执行它。

返回值
eval() 函数将字符串转换为相应的对象，并返回表达式的结果。
"""

x = 8
print(eval("3 * x"))

print(eval("pow(2, 2)"))

n = 81
print(eval("n + 5"))

# 执行简单的数学表达式
result = eval("2 + 3 * 4")
print(result)  # 输出: 14

# 执行变量引用
x = 10
result = eval("x + 5")
print(result)  # 输出: 15

# 在指定命名空间中执行表达式
namespace = {"a": 2, "b": 3}
result = eval("a + b", namespace)
print(result)  # 输出: 5
