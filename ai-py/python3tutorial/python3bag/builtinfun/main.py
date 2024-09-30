# -*- encoding: utf-8 -*-
"""
@File   : main.py
@Time   : 2024-09-27 10:12:34
@Author : pan
"""


# builtin 函数功能示例
class Person:
    """
    __eq__方法
    """

    def __init__(self, name, age):
        self.name = name
        self.age = age

    def __eq__(self, other):
        # 检查 other 是否是 Person 的一个实例
        if isinstance(other, Person):
            # 比较 name 和 age 属性
            return self.name == other.name and self.age == other.age
        # 如果 other 不是 Person 的一个实例，返回 False
        return False


def PersonRun():
    # 创建两个 Person 对象
    person1 = Person("Alice", 30)
    person2 = Person("Alice", 30)
    person3 = Person("Bob", 25)
    # 使用 == 操作符比较对象
    print(person1 == person2)  # 输出: True，因为 name 和 age 都相同
    print(person1 == person3)  # 输出: False，因为 name 或 age 不同


class PersonProperty:
    """
    @property
    方法定义上面加一个@property 装饰器，在不改变原有调用方式的同时，来将一个属性改为一个方法
    """

    def __init__(self, first_name, last_name):
        self.first = first_name
        self.last = last_name

    @property  # 添加@property属性，可以在不改变原有调用规则的基础上，获得正确的fullname
    def fullname(self):
        return self.first + " " + self.last

    @fullname.setter
    def fullname(self, name):
        """
        setter 方法需要和@property 修饰的方法具有相同的名字
        它会将用户传给property的值，作为参数
        最后你需要在方法定义上添加@{methodname}.setter 装饰器
        """
        first_name, last_name = name.split()
        self.first = first_name
        self.last = last_name

    def email(self):
        return "{}.{}@email.com".format(self.first, self.last)


def PropertyRun():
    person = Person("zhang", "san")
    print(person.fullname)
    print(person.last)
    print(person.first)

    person.fullname = "li si"
    print(person.fullname)
    print(person.last)
    print(person.first)


def demo():
    string = "hello, world!"
    print(string.startswith("hello"))
    print(string.endswith("world!"))
    print(string.strip(" "))
    print(string.split(", "))
    print(string.capitalize())
    pass


if __name__ == "__main__":
    demo()
    # PersonRun()
    # PropertyRun()
    pass
